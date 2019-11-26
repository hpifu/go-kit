package hflag

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Flag struct {
	Name      string
	Shorthand string
	Usage     string
	Type      string
	DefValue  string
	Required  bool
	Assigned  bool
	Value     Value
}

func (f *Flag) Set(val string) error {
	if err := f.Value.Set(val); err != nil {
		return err
	}

	f.Assigned = true
	return nil
}

type FlagSet struct {
	name            string
	nameToFlag      map[string]*Flag
	shorthandToName map[string]string
	posFlagNames    []string
	flagNames       []string
	args            []string
	parsed          bool
}

func NewFlagSet(name string) *FlagSet {
	return &FlagSet{
		name:            name,
		nameToFlag:      map[string]*Flag{},
		shorthandToName: map[string]string{},
		parsed:          false,
	}
}

func (f *FlagSet) allBoolFlag(name string) bool {
	for i := 0; i < len(name); i++ {
		flag := f.Lookup(name[i : i+1])
		if flag == nil || flag.Type != "bool" {
			return false
		}
	}

	return true
}

func isBoolValue(val string) bool {
	_, err := strconv.ParseBool(val)
	if err != nil {
		return false
	}
	return true
}

func (f *FlagSet) Parse(args []string) error {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if !strings.HasPrefix(arg, "-") {
			f.args = append(f.args, arg)
			continue
		}
		option := arg[1:]
		if strings.HasPrefix(arg, "--") {
			option = arg[2:]
		}
		if strings.Contains(option, "=") {
			idx := strings.Index(option, "=")
			name := option[0:idx]
			val := option[idx+1:]
			flag := f.Lookup(name)
			if flag == nil {
				return fmt.Errorf("unknow option [%v]", name)
			}
			if err := flag.Set(val); err != nil {
				return fmt.Errorf("set failed. name: [%v], val: [%v], type: [%v]", name, val, flag.Type)
			}
		} else if f.Lookup(option) != nil {
			name := option
			flag := f.Lookup(name)
			if flag == nil {
				return fmt.Errorf("unknow flag [%v]", name)
			}
			if flag.Type != "bool" { // 参数不是 bool，后面必有一个值
				if i+1 >= len(args) {
					return fmt.Errorf("miss argument for nonboolean option [%v]", name)
				}
				val := args[i+1]
				if err := flag.Set(val); err != nil {
					return fmt.Errorf("set failed. name: [%v], val: [%v], type: [%v]", name, val, flag.Type)
				}
				i++
			} else { // 参数为 bool 类型，如果后面的值为 true 或者 false 则设为后面值，否则设置为 true
				val := "true"
				if i+1 < len(args) && isBoolValue(args[i+1]) {
					val = args[i+1]
					i++
				}
				if err := flag.Set(val); err != nil {
					return fmt.Errorf("set failed. name: [%v], val: [%v], type: [%v]", name, val, flag.Type)
				}
			}
		} else if f.allBoolFlag(option) { // -kval 全是 bool 选项，-kval 和 -k -v -f -l 等效
			for i := 0; i < len(option); i++ {
				name := option[i : i+1]
				flag := f.Lookup(name)
				if err := flag.Set("true"); err != nil {
					return fmt.Errorf("set failed. name: [%v], val: [%v], type: [%v]", name, "true", flag.Type)
				}
			}
		} else {
			name := option[0:1]
			val := arg[1:]
			flag := f.Lookup(name)
			if flag == nil {
				return fmt.Errorf("unknow option [%v]", name)
			}
			if err := flag.Set(val); err != nil {
				return fmt.Errorf("set failed. name: [%v], val: [%v], type: [%v]", name, val, flag.Type)
			}
		}
	}

	for i, name := range f.posFlagNames {
		if i >= len(f.args) {
			break
		}
		val := f.args[i]
		flag := f.nameToFlag[name]
		if err := flag.Set(val); err != nil {
			return fmt.Errorf("set any failed. name: [%v], val: [%v], type: [%v]", name, val, flag.Type)
		}
	}

	// Required check
	for name, flag := range f.nameToFlag {
		if flag.Required && !flag.Assigned {
			return fmt.Errorf("option [%v] is required, but not assigned", name)
		}
	}

	f.parsed = true

	return nil
}

func (f *FlagSet) AddFlag(name string, shorthand string, usage string, typeStr string, required bool, defaultValue string) error {
	if _, ok := f.nameToFlag[name]; ok {
		return fmt.Errorf("conflict flag [%v]", name)
	}

	if shorthand != "" {
		if _, ok := f.shorthandToName[shorthand]; ok {
			return fmt.Errorf("conflict shorthand [%v]", shorthand)
		}
	}

	flag := &Flag{
		Name:      name,
		Shorthand: shorthand,
		Usage:     usage,
		Type:      typeStr,
		Required:  required,
		DefValue:  defaultValue,
		Value:     NewValueType(typeStr),
	}

	if len(defaultValue) != 0 {
		if err := flag.Set(defaultValue); err != nil {
			return fmt.Errorf("set default failed. err: [%v]", err)
		}
	}

	f.nameToFlag[name] = flag
	f.shorthandToName[shorthand] = name
	f.flagNames = append(f.flagNames, name)

	return nil
}

func (f *FlagSet) AddPosFlag(name string, usage string, typeStr string, defaultValue string) error {
	if _, ok := f.nameToFlag[name]; ok {
		return fmt.Errorf("conflict flag [%v]", name)
	}

	flag := &Flag{
		Name:     name,
		Usage:    usage,
		Type:     typeStr,
		DefValue: defaultValue,
		Value:    NewValueType(typeStr),
	}

	if len(defaultValue) != 0 {
		if err := flag.Set(defaultValue); err != nil {
			return fmt.Errorf("set default failed. err: [%v]", err)
		}
	}

	f.nameToFlag[name] = flag
	f.posFlagNames = append(f.posFlagNames, name)

	return nil
}

func (f *FlagSet) Usage() string {
	var buffer bytes.Buffer

	buffer.WriteString("usage: ")
	for _, name := range f.posFlagNames {
		p := f.nameToFlag[name]
		buffer.WriteString(fmt.Sprintf(" [%v]", p.Name))
	}

	sort.Strings(f.flagNames)
	for _, name := range f.flagNames {
		flag := f.nameToFlag[name]
		if flag.Required {
			buffer.WriteString(fmt.Sprintf(" <--%v=%v>", flag.Name, flag.Type))
		} else {
			buffer.WriteString(fmt.Sprintf(" [--%v=%v]", flag.Name, flag.Type))
		}
	}
	buffer.WriteString("\n")

	buffer.WriteString("\npositional options:\n")
	for _, name := range f.posFlagNames {
		flag := f.nameToFlag[name]
		shorthand := ""
		name := flag.Name
		defaultValue := flag.Type
		if flag.DefValue != "" {
			defaultValue = flag.Type + "=" + flag.DefValue
		}
		defaultValue = "[" + defaultValue + "]"
		buffer.WriteString(fmt.Sprintf("%4v  %-15v %-15v %v\n", shorthand, name, defaultValue, flag.Usage))
	}

	buffer.WriteString("\noptions:\n")
	for _, name := range f.flagNames {
		flag := f.nameToFlag[name]
		shorthand := ""
		if flag.Shorthand != "" {
			shorthand = "-" + flag.Shorthand
		}
		name := "--" + flag.Name
		defaultValue := flag.Type
		if flag.DefValue != "" {
			defaultValue = flag.Type + "=" + flag.DefValue
		}
		defaultValue = "[" + defaultValue + "]"
		buffer.WriteString(fmt.Sprintf("%4v, %-15v %-15v %v\n", shorthand, name, defaultValue, flag.Usage))
	}

	return buffer.String()
}
