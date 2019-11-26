package hflag

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Flag struct {
	name         string
	shorthand    string
	help         string
	typeStr      string
	defaultValue string
	required     bool
	assigned     bool
	value        Value
}

type FlagSet struct {
	nameToFlag      map[string]*Flag
	shorthandToName map[string]string
	posFlagNames    []string
	flagNames       []string
}

func NewFlagSet() *FlagSet {
	return &FlagSet{
		nameToFlag:      map[string]*Flag{},
		shorthandToName: map[string]string{},
	}
}

func (f *FlagSet) Usage() string {
	var buffer bytes.Buffer

	buffer.WriteString("usage: ")
	for _, name := range f.posFlagNames {
		p := f.nameToFlag[name]
		buffer.WriteString(fmt.Sprintf(" [%v]", p.name))
	}

	sort.Strings(f.flagNames)
	for _, name := range f.flagNames {
		flag := f.nameToFlag[name]
		if flag.required {
			buffer.WriteString(fmt.Sprintf(" <--%v=%v>", flag.name, flag.typeStr))
		} else {
			buffer.WriteString(fmt.Sprintf(" [--%v=%v]", flag.name, flag.typeStr))
		}
	}
	buffer.WriteString("\n")

	buffer.WriteString("\npositional options:\n")
	for _, name := range f.posFlagNames {
		flag := f.nameToFlag[name]
		shorthand := ""
		name := flag.name
		defaultValue := flag.typeStr
		if flag.defaultValue != "" {
			defaultValue = flag.typeStr + "=" + flag.defaultValue
		}
		defaultValue = "[" + defaultValue + "]"
		buffer.WriteString(fmt.Sprintf("%4v  %-15v %-15v %v\n", shorthand, name, defaultValue, flag.help))
	}

	buffer.WriteString("\noptions:\n")
	for _, name := range f.flagNames {
		flag := f.nameToFlag[name]
		shorthand := ""
		if flag.shorthand != "" {
			shorthand = "-" + flag.shorthand
		}
		name := "--" + flag.name
		defaultValue := flag.typeStr
		if flag.defaultValue != "" {
			defaultValue = flag.typeStr + "=" + flag.defaultValue
		}
		defaultValue = "[" + defaultValue + "]"
		buffer.WriteString(fmt.Sprintf("%4v, %-15v %-15v %v\n", shorthand, name, defaultValue, flag.help))
	}
	fmt.Println(buffer.String())

	return buffer.String()
}

func (f *FlagSet) BoolVar(b *bool, name string, defaultValue bool, help string) {
	*b = defaultValue
	err := f.addFlag(name, "", help, "bool", false, fmt.Sprintf("%v", defaultValue))
	if err != nil {
		panic(err)
	}
	f.nameToFlag[name].value = (*boolValue)(b)
}

func (f *FlagSet) Bool(name string, defaultValue bool, help string) *bool {
	err := f.addFlag(name, "", help, "bool", false, fmt.Sprintf("%v", defaultValue))
	if err != nil {
		panic(err)
	}
	return (*bool)(f.nameToFlag[name].value.(*boolValue))
}

func (f *FlagSet) Int(name string, defaultValue int, help string) *int {
	if err := f.addFlag(name, "", help, "int", false, strconv.Itoa(defaultValue)); err != nil {
		panic(err)
	}
	return (*int)(f.nameToFlag[name].value.(*intValue))
}

func (f *FlagSet) String(name string, defaultValue string, help string) *string {
	if err := f.addFlag(name, "", help, "string", false, defaultValue); err != nil {
		panic(err)
	}
	return (*string)(f.nameToFlag[name].value.(*stringValue))
}

func (f *FlagSet) Duration(name string, defaultValue time.Duration, help string) *time.Duration {
	if err := f.addFlag(name, "", help, "duration", false, defaultValue.String()); err != nil {
		panic(err)
	}
	return (*time.Duration)(f.nameToFlag[name].value.(*durationValue))
}

func (f *FlagSet) Float(name string, defaultValue float64, help string) *float64 {
	if err := f.addFlag(name, "", help, "float", false, fmt.Sprintf("%f", defaultValue)); err != nil {
		panic(err)
	}
	return (*float64)(f.nameToFlag[name].value.(*floatValue))
}

func (f *FlagSet) Parse(args []string) error {
	var posFlagValues []string
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--") {
			if strings.Contains(arg, "=") { // --key1=val1
				idx := strings.Index(arg, "=")
				name := arg[2:idx]
				val := arg[idx+1:]
				if flag, ok := f.nameToFlag[name]; !ok {
					return fmt.Errorf("unknow option [%v]", name)
				} else {
					if err := flag.value.Set(val); err != nil {
						return fmt.Errorf("set any failed. name: [%v], val: [%v], type: [%v]", name, val, flag.typeStr)
					}
				}
			} else { // --key1 val1
				name := arg[2:]
				if flag, ok := f.nameToFlag[name]; !ok {
					return fmt.Errorf("unknow option [%v]", name)
				} else if flag.typeStr != "bool" { // 参数不是 bool，后面必有一个值
					if i+1 >= len(args) {
						return fmt.Errorf("miss any for nonboolean option [%v]", name)
					}
					val := args[i+1]
					if err := flag.value.Set(val); err != nil {
						return fmt.Errorf("set any failed. name: [%v], val: [%v], type: [%v]", name, val, flag.typeStr)
					}
					i++
				} else { // 参数为 bool 类型，如果后面的值为 true 或者 false 则设为后面值，否则设置为 true
					val := "true"
					if i+1 < len(args) && (args[i+1] == "true" || args[i+1] == "false") {
						val = args[i+1]
						i++
					}
					if err := flag.value.Set(val); err != nil {
						return fmt.Errorf("set any failed. name: [%v], val: [%v], type: [%v]", name, val, flag.typeStr)
					}
				}
			}
		} else if strings.HasPrefix(arg, "-") {
			if len(arg) == 2 { // -k val
				name, ok := f.shorthandToName[arg[1:]]
				if !ok {
					return fmt.Errorf("unknow shorthand option [%v]", arg[1:])
				}
				flag := f.nameToFlag[name]
				if flag.typeStr != "bool" { // 参数不是 bool 类型，后面必有一个值
					if i+1 >= len(args) {
						return fmt.Errorf("miss any for nonboolean option [%v]", name)
					}
					val := args[i+1]
					if err := flag.value.Set(val); err != nil {
						return fmt.Errorf("set any failed. name: [%v], val: [%v], type: [%v]", name, val, flag.typeStr)
					}
					i++
				} else { // 参数为 bool 类型，如果后面的值为 true 或者 false 则设为后面值，否则设置为 true
					val := "true"
					if i+1 < len(args) && (args[i+1] == "true" || args[i+1] == "false") {
						val = args[i+1]
						i++
					}
					if err := flag.value.Set(val); err != nil {
						return fmt.Errorf("set any failed. name: [%v], val: [%v], type: [%v]", name, val, flag.typeStr)
					}
				}
			} else { // -kval
				allBool := true
				for i := 1; i < len(arg); i++ {
					name, ok := f.shorthandToName[arg[i:i+1]]
					if !ok {
						allBool = false
						break
					}
					if f.nameToFlag[name].typeStr != "bool" {
						allBool = false
					}
				}
				if allBool { // 全是 bool 选项，-kval 和 -k -v -f -l 等效
					for i := 1; i < len(arg); i++ {
						name := f.shorthandToName[arg[i:i+1]]
						flag := f.nameToFlag[name]
						if err := flag.value.Set("true"); err != nil {
							return fmt.Errorf("set any failed. name: [%v], val: [%v], type: [%v]", name, "true", flag.typeStr)
						}
					}
				} else {
					name, ok := f.shorthandToName[arg[1:2]]
					if !ok {
						return fmt.Errorf("unknow shorthand option [%v]", arg[1:2])
					}
					val := arg[2:]
					flag := f.nameToFlag[name]
					if err := flag.value.Set(val); err != nil {
						return fmt.Errorf("set any failed. name: [%v], val: [%v], type: [%v]", name, val, flag.typeStr)
					}
				}
			}
		} else {
			posFlagValues = append(posFlagValues, arg)
		}
	}

	for i, name := range f.posFlagNames {
		if i >= len(posFlagValues) {
			break
		}
		val := posFlagValues[i]
		p := f.nameToFlag[name]
		if err := p.value.Set(val); err != nil {
			return fmt.Errorf("set any failed. name: [%v], val: [%v], type: [%v]", name, val, p.typeStr)
		}
	}

	// required check
	//for name, flag := range f.nameToFlag {
	//	if flag.required {
	//		return fmt.Errorf("option [%v] is required, but not assigned", name)
	//	}
	//}

	//for name, val := range f.nameToFlag {
	//	fmt.Println(name, "=>", val.any)
	//}

	return nil
}

func (f *FlagSet) addFlag(name string, shorthand string, help string, typeStr string, required bool, defaultValue string) error {
	if _, ok := f.nameToFlag[name]; ok {
		return fmt.Errorf("conflict flag [%v]", name)
	}

	if shorthand != "" {
		if _, ok := f.shorthandToName[shorthand]; ok {
			return fmt.Errorf("conflict shorthand [%v]", shorthand)
		}
	}

	flag := &Flag{
		name:         name,
		shorthand:    shorthand,
		help:         help,
		typeStr:      typeStr,
		required:     required,
		defaultValue: defaultValue,
		value:        NewValueType(typeStr),
	}

	if len(defaultValue) != 0 {
		if err := flag.value.Set(defaultValue); err != nil {
			return fmt.Errorf("set default any failed. err: [%v]", err)
		}
	}

	f.nameToFlag[name] = flag
	f.shorthandToName[shorthand] = name
	f.flagNames = append(f.flagNames, name)

	return nil
}

func (f *FlagSet) addPosFlag(name string, help string, typeStr string, defaultValue string) error {
	if _, ok := f.nameToFlag[name]; ok {
		return fmt.Errorf("conflict flag [%v]", name)
	}

	flag := &Flag{
		name:         name,
		help:         help,
		typeStr:      typeStr,
		defaultValue: defaultValue,
		value:        NewValueType(typeStr),
	}

	if len(defaultValue) != 0 {
		if err := flag.value.Set(defaultValue); err != nil {
			return fmt.Errorf("set default any failed. err: [%v]", err)
		}
	}

	f.nameToFlag[name] = flag
	f.posFlagNames = append(f.posFlagNames, name)

	return nil
}
