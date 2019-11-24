package hflag

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Any struct {
	typeStr  string
	assigned bool
	i        int
	s        string
	f        float64
	d        time.Duration
	b        bool
}

func (a *Any) Set(val string, typeStr string) error {
	switch typeStr {
	case "int":
		i, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		a.i = i
	case "string":
		a.s = val
	case "float":
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		a.f = f
	case "duration":
		d, err := time.ParseDuration(val)
		if err != nil {
			return err
		}
		a.d = d
	case "bool":
		if val == "false" {
			a.b = false
		} else {
			a.b = true
		}
	default:
		return fmt.Errorf("unsupport type [%v]", typeStr)
	}

	a.typeStr = typeStr
	a.assigned = true

	return nil
}

type Flag struct {
	name         string
	shorthand    string
	help         string
	required     bool
	typeStr      string
	defaultValue string
	value        *Any
}

type FlagSet struct {
	optionMap       map[string]*Flag
	shorthandToName map[string]string
	positionOption  []string
}

func NewFlagSet() *FlagSet {
	return &FlagSet{
		optionMap:       map[string]*Flag{},
		shorthandToName: map[string]string{},
	}
}

func (f *FlagSet) Usage() string {
	var buffer bytes.Buffer

	buffer.WriteString("usage: ")
	positionOptionSet := map[string]bool{}
	for _, name := range f.positionOption {
		positionOptionSet[name] = true
		p := f.optionMap[name]
		buffer.WriteString(fmt.Sprintf(" [%v]", p.name))
	}

	for k, p := range f.optionMap {
		if positionOptionSet[k] {
			continue
		}
		if p.required {
			buffer.WriteString(fmt.Sprintf(" <--%v=%v>", p.name, p.typeStr))
		} else {
			buffer.WriteString(fmt.Sprintf(" [--%v=%v]", p.name, p.typeStr))
		}
	}
	buffer.WriteString("\n\n")

	for _, v := range f.optionMap {
		//if positionOptionSet[k] {
		//	continue
		//}
		shorthand := ""
		if v.shorthand != "" {
			shorthand = "-" + v.shorthand
		}
		name := "--" + v.name
		defaultValue := v.typeStr
		if v.defaultValue != "" {
			defaultValue = v.typeStr + "=" + v.defaultValue
		}
		defaultValue = "[" + defaultValue + "]"
		buffer.WriteString(fmt.Sprintf("%11v, %-15v %-15v %v\n", shorthand, name, defaultValue, v.help))
	}
	fmt.Println(buffer.String())

	return buffer.String()
}

func (f *FlagSet) Bool(name string, defaultValue bool, help string) *bool {
	err := f.addOption(name, "", help, "bool", false, fmt.Sprintf("%v", defaultValue))
	if err != nil {
		panic(err)
	}

	return &f.optionMap[name].value.b
}

func (f *FlagSet) Int(name string, defaultValue int, help string) *int {
	if err := f.addOption(name, "", help, "int", false, strconv.Itoa(defaultValue)); err != nil {
		panic(err)
	}

	return &f.optionMap[name].value.i
}

func (f *FlagSet) String(name string, defaultValue string, help string) *string {
	if err := f.addOption(name, "", help, "string", false, defaultValue); err != nil {
		panic(err)
	}

	return &f.optionMap[name].value.s
}

func (f *FlagSet) Duration(name string, defaultValue time.Duration, help string) *time.Duration {
	if err := f.addOption(name, "", help, "duration", false, defaultValue.String()); err != nil {
		panic(err)
	}

	return &f.optionMap[name].value.d
}

func (f *FlagSet) parse(args []string) error {
	var positionParam []string
	//optionParam := map[string]string{}
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--") {
			if strings.Contains(arg, "=") { // --key1=val1
				idx := strings.Index(arg, "=")
				key := arg[2:idx]
				val := arg[idx+1:]
				if p, ok := f.optionMap[key]; !ok {
					return fmt.Errorf("unknow option [%v]", key)
				} else {
					if err := p.value.Set(val, p.typeStr); err != nil {
						return fmt.Errorf("set value failed. key: [%v], val: [%v], type: [%v]", key, val, p.typeStr)
					}
				}
			} else { // --key1 val1
				key := arg[2:]
				if p, ok := f.optionMap[key]; !ok {
					return fmt.Errorf("unknow option [%v]", key)
				} else if p.typeStr != "bool" { // 参数不是 bool，后面必有一个值
					if i+1 >= len(args) {
						return fmt.Errorf("miss value for nonboolean option [%v]", key)
					}
					val := args[i+1]
					if err := p.value.Set(val, p.typeStr); err != nil {
						return fmt.Errorf("set value failed. key: [%v], val: [%v], type: [%v]", key, val, p.typeStr)
					}
					i++
				} else { // 参数为 bool 类型，如果后面的值为 true 或者 false 则设为后面值，否则设置为 true
					val := "true"
					if i+1 < len(args) && (args[i+1] == "true" || args[i+1] == "false") {
						val = args[i+1]
						i++
					}
					if err := p.value.Set(val, p.typeStr); err != nil {
						return fmt.Errorf("set value failed. key: [%v], val: [%v], type: [%v]", key, val, p.typeStr)
					}
				}
			}
		} else if strings.HasPrefix(arg, "-") {
			if len(arg) == 2 { // -k val
				key, ok := f.shorthandToName[arg[1:]]
				if !ok {
					return fmt.Errorf("unknow shorthand option [%v]", arg[1:])
				}
				p := f.optionMap[key]
				if p.typeStr != "bool" { // 参数不是 bool 类型，后面必有一个值
					if i+1 >= len(args) {
						return fmt.Errorf("miss value for nonboolean option [%v]", key)
					}
					val := args[i+1]
					if err := p.value.Set(val, p.typeStr); err != nil {
						return fmt.Errorf("set value failed. key: [%v], val: [%v], type: [%v]", key, val, p.typeStr)
					}
					i++
				} else { // 参数为 bool 类型，如果后面的值为 true 或者 false 则设为后面值，否则设置为 true
					val := "true"
					if i+1 < len(args) && (args[i+1] == "true" || args[i+1] == "false") {
						val = args[i+1]
						i++
					}
					if err := p.value.Set(val, p.typeStr); err != nil {
						return fmt.Errorf("set value failed. key: [%v], val: [%v], type: [%v]", key, val, p.typeStr)
					}
				}
			} else { // -kval
				allBool := true
				for i := 1; i < len(arg); i++ {
					key, ok := f.shorthandToName[arg[i:i+1]]
					if !ok {
						allBool = false
						break
					}
					if f.optionMap[key].typeStr != "bool" {
						allBool = false
					}
				}
				if allBool { // 全是 bool 选项，-kval 和 -k -v -f -l 等效
					for i := 1; i < len(arg); i++ {
						key := f.shorthandToName[arg[i:i+1]]
						p := f.optionMap[key]
						if err := p.value.Set("true", p.typeStr); err != nil {
							return fmt.Errorf("set value failed. key: [%v], val: [%v], type: [%v]", key, "true", p.typeStr)
						}
					}
				} else {
					key, ok := f.shorthandToName[arg[1:2]]
					if !ok {
						return fmt.Errorf("unknow shorthand option [%v]", arg[1:2])
					}
					//p := f.optionMap[key]
					val := arg[2:]
					p := f.optionMap[key]
					if err := p.value.Set(val, p.typeStr); err != nil {
						return fmt.Errorf("set value failed. key: [%v], val: [%v], type: [%v]", key, val, p.typeStr)
					}
				}
			}
		} else {
			positionParam = append(positionParam, arg)
		}
	}

	for i, key := range f.positionOption {
		if i >= len(positionParam) {
			break
		}
		val := positionParam[i]
		p := f.optionMap[key]
		if err := p.value.Set(val, p.typeStr); err != nil {
			return fmt.Errorf("set value failed. key: [%v], val: [%v], type: [%v]", key, val, p.typeStr)
		}
	}

	for key, p := range f.optionMap {
		if p.required && !p.value.assigned {
			return fmt.Errorf("option [%v] is required, but not assigned", key)
		}
	}

	//for key, val := range f.optionMap {
	//	fmt.Println(key, "=>", val.value)
	//}

	return nil
}

func (f *FlagSet) addOption(name string, shorthand string, help string, typeStr string, required bool, defaultValue string) error {
	if _, ok := f.optionMap[name]; ok {
		return fmt.Errorf("conflict option [%v]", name)
	}

	if shorthand != "" {
		if _, ok := f.shorthandToName[shorthand]; ok {
			return fmt.Errorf("conflict shorthand [%v]", shorthand)
		}
	}

	option := &Flag{
		name:         name,
		shorthand:    shorthand,
		help:         help,
		typeStr:      typeStr,
		required:     required,
		defaultValue: defaultValue,
		value:        &Any{},
	}

	if len(defaultValue) != 0 {
		if err := option.value.Set(defaultValue, typeStr); err != nil {
			return fmt.Errorf("set default value failed. err: [%v]", err)
		}
	}

	f.optionMap[name] = option
	f.shorthandToName[shorthand] = name

	return nil
}

func (f *FlagSet) addPositionOption(name string, help string, typeStr string, defaultValue string) error {
	if _, ok := f.optionMap[name]; ok {
		return fmt.Errorf("conflict option [%v]", name)
	}

	option := &Flag{
		name:         name,
		help:         help,
		typeStr:      typeStr,
		defaultValue: defaultValue,
		value:        &Any{},
	}

	if len(defaultValue) != 0 {
		if err := option.value.Set(defaultValue, typeStr); err != nil {
			return fmt.Errorf("set default value failed. err: [%v]", err)
		}
	}

	f.optionMap[name] = option
	f.positionOption = append(f.positionOption, name)

	return nil
}
