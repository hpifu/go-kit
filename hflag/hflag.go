package hflag

import (
	"bytes"
	"fmt"
	"github.com/hpifu/go-kit/hstring"
	"net"
	"path"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
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
		if strings.Contains(option, "=") { // 选项中含有等号，按照等号分割成 name val
			idx := strings.Index(option, "=")
			name := option[0:idx]
			val := option[idx+1:]
			flag := f.Lookup(name)
			if flag == nil {
				return fmt.Errorf("unknow option [%v]", name)
			}
			if err := flag.Set(val); err != nil {
				return fmt.Errorf("set failed. name: [%v], val: [%v], type: [%v], err: [%v]", name, val, flag.Type, err)
			}
		} else if f.Lookup(option) != nil {
			name := option
			flag := f.Lookup(name)
			if flag == nil {
				return fmt.Errorf("unknow flag [%v]", name)
			}
			if flag.Type != "bool" { // 选项不是 bool，后面必有一个值
				if i+1 >= len(args) {
					return fmt.Errorf("miss argument for nonboolean option [%v]", name)
				}
				val := args[i+1]
				if err := flag.Set(val); err != nil {
					return fmt.Errorf("set failed. name: [%v], val: [%v], type: [%v], err: [%v]", name, val, flag.Type, err)
				}
				i++
			} else { // 选项为 bool 类型，如果后面的值为合法的 bool 值，否则设置为 true
				val := "true"
				if i+1 < len(args) && isBoolValue(args[i+1]) {
					val = args[i+1]
					i++
				}
				if err := flag.Set(val); err != nil {
					return fmt.Errorf("set failed. name: [%v], val: [%v], type: [%v], err: [%v]", name, val, flag.Type, err)
				}
			}
		} else if f.allBoolFlag(option) { // -aux 全是 bool 选项，-aux 和 -a -u -x 等效
			for i := 0; i < len(option); i++ {
				name := option[i : i+1]
				flag := f.Lookup(name)
				if err := flag.Set("true"); err != nil {
					return fmt.Errorf("set failed. name: [%v], val: [%v], type: [%v], err: [%v]", name, "true", flag.Type, err)
				}
			}
		} else { // -p123456 和 -p 123456 等效
			name := option[0:1]
			val := option[1:]
			flag := f.Lookup(name)
			if flag == nil {
				return fmt.Errorf("unknow option [%v]", name)
			}
			if err := flag.Set(val); err != nil {
				return fmt.Errorf("set failed. name: [%v], val: [%v], type: [%v], err: [%v]", name, val, flag.Type, err)
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

type FlagOptions struct {
	shorthand    string
	typeStr      string
	required     bool
	defaultValue string
}

type FlagOption func(*FlagOptions)

func NewFlagOptions() *FlagOptions {
	return &FlagOptions{
		shorthand:    "",
		typeStr:      "string",
		required:     false,
		defaultValue: "",
	}
}

func Required() FlagOption {
	return func(o *FlagOptions) {
		o.required = true
	}
}

func DefaultValue(val string) FlagOption {
	return func(o *FlagOptions) {
		o.defaultValue = val
	}
}

func Shorthand(shorthand string) FlagOption {
	return func(o *FlagOptions) {
		o.shorthand = shorthand
	}
}

func Type(typeStr string) FlagOption {
	return func(o *FlagOptions) {
		o.typeStr = typeStr
	}
}

func (f *FlagSet) AddFlag(name string, usage string, opts ...FlagOption) {
	o := NewFlagOptions()
	for _, opt := range opts {
		opt(o)
	}

	if err := f.addFlag(name, usage, o.shorthand, o.typeStr, o.required, o.defaultValue); err != nil {
		panic(err)
	}
}

func (f *FlagSet) AddPosFlag(name string, usage string, opts ...FlagOption) {
	o := NewFlagOptions()
	for _, opt := range opts {
		opt(o)
	}

	if err := f.addPosFlag(name, usage, o.typeStr, o.required, o.defaultValue); err != nil {
		panic(err)
	}
}

func (f *FlagSet) addFlag(name string, usage string, shorthand string, typeStr string, required bool, defaultValue string) error {
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

	if flag.Value == nil {
		return fmt.Errorf("type [%v] not support", typeStr)
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

func (f *FlagSet) addPosFlag(name string, usage string, typeStr string, required bool, defaultValue string) error {
	if _, ok := f.nameToFlag[name]; ok {
		return fmt.Errorf("conflict flag [%v]", name)
	}

	flag := &Flag{
		Name:     name,
		Usage:    usage,
		Type:     typeStr,
		Required: required,
		DefValue: defaultValue,
		Value:    NewValueType(typeStr),
	}

	if flag.Value == nil {
		return fmt.Errorf("type [%v] not support", typeStr)
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
	type info struct {
		shorthand   string
		name        string
		typeDefault string
		usage       string
	}

	var posFlagInfos []*info
	var flagInfos []*info

	for _, name := range f.posFlagNames {
		flag := f.nameToFlag[name]
		defaultValue := flag.Type
		if flag.DefValue != "" {
			defaultValue = flag.Type + "=" + flag.DefValue
		}
		posFlagInfos = append(posFlagInfos, &info{
			shorthand:   "",
			name:        flag.Name,
			typeDefault: "[" + defaultValue + "]",
			usage:       flag.Usage,
		})
	}

	sort.Strings(f.flagNames)
	for _, name := range f.flagNames {
		flag := f.nameToFlag[name]
		defaultValue := flag.Type
		if flag.DefValue != "" {
			defaultValue = flag.Type + "=" + flag.DefValue
		}
		shorthand := ""
		if flag.Shorthand != "" {
			shorthand = "-" + flag.Shorthand
		}
		flagInfos = append(flagInfos, &info{
			shorthand:   shorthand,
			name:        "--" + flag.Name,
			typeDefault: "[" + defaultValue + "]",
			usage:       flag.Usage,
		})
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var shorthandWidth, nameWidth, typeDefaultWidth int
	for _, i := range append(posFlagInfos, flagInfos...) {
		shorthandWidth = max(len(i.shorthand), shorthandWidth)
		nameWidth = max(len(i.name), nameWidth)
		typeDefaultWidth = max(len(i.typeDefault), typeDefaultWidth)
	}

	var buffer bytes.Buffer

	buffer.WriteString("usage: ")
	buffer.WriteString(path.Base(f.name))
	for _, name := range f.posFlagNames {
		p := f.nameToFlag[name]
		buffer.WriteString(fmt.Sprintf(" [%v]", p.Name))
	}

	for _, name := range f.flagNames {
		flag := f.nameToFlag[name]
		nameShorthand := flag.Name
		if flag.Shorthand != "" {
			nameShorthand = flag.Shorthand + "," + flag.Name
		}
		if flag.DefValue != "" {
			buffer.WriteString(fmt.Sprintf(" [-%v %v=%v]", nameShorthand, flag.Type, flag.DefValue))
		} else if flag.Required {
			buffer.WriteString(fmt.Sprintf(" <-%v %v>", nameShorthand, flag.Type))
		} else {
			buffer.WriteString(fmt.Sprintf(" [-%v %v]", nameShorthand, flag.Type))
		}
	}
	buffer.WriteString("\n")

	if len(posFlagInfos) != 0 {
		buffer.WriteString("\npositional options:\n")
		posFormat := fmt.Sprintf("  %%%dv  %%-%dv  %%-%dv  %%v\n", shorthandWidth, nameWidth, typeDefaultWidth)
		for _, i := range posFlagInfos {
			buffer.WriteString(fmt.Sprintf(posFormat, i.shorthand, i.name, i.typeDefault, i.usage))
		}
	}
	buffer.WriteString("\noptions:\n")
	format := fmt.Sprintf("  %%%dv, %%-%dv  %%-%dv  %%v\n", shorthandWidth, nameWidth, typeDefaultWidth)
	for _, i := range flagInfos {
		buffer.WriteString(fmt.Sprintf(format, i.shorthand, i.name, i.typeDefault, i.usage))
	}

	return buffer.String()
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

func parseTag(tag string) (name string, shorthand string, usage string, required bool, defaultValue string, position bool, err error) {
	if strings.Trim(tag, " ") == "" {
		position = false
		return
	}
	for _, field := range strings.Split(tag, ";") {
		field = strings.Trim(field, " ")
		if field == "required" { // required
			required = true
		} else if strings.HasPrefix(field, "--") { // --int-option, -i
			names := strings.Split(field, ",")
			name = strings.Trim(names[0], " ")[2:]
			if len(names) > 2 {
				err = fmt.Errorf("expected name field format is '--<name>[, -<shorthand>]', got [%v]", field)
				return
			} else if len(names) == 2 {
				shorthand = strings.Trim(names[1], " ")
				if !strings.HasPrefix(shorthand, "-") {
					err = fmt.Errorf("expected name field format is '--<name>[, -<shorthand>]', got [%v]", field)
					return
				}
				shorthand = shorthand[1:]
			}
		} else if strings.Contains(field, ":") { // default: 10; usage: int flag
			kvs := strings.Split(field, ":")
			if len(kvs) != 2 {
				err = fmt.Errorf("expected format '<key>:<value>', got [%v]", field)
				return
			}
			key := strings.Trim(kvs[0], " ")
			val := strings.Trim(kvs[1], " ")
			switch key {
			case "default":
				defaultValue = val
			case "usage":
				usage = val
			}
		} else { // pos
			name = strings.Trim(field, " ")
			position = true
		}
	}

	return
}

func interfaceToType(v reflect.Value) (string, Value, error) {
	switch v.Interface().(type) {
	case bool:
		return "bool", (*boolValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case int:
		return "int", (*intValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case uint:
		return "uint", (*uintValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case int64:
		return "int64", (*int64Value)(unsafe.Pointer(v.Addr().Pointer())), nil
	case int32:
		return "int32", (*int32Value)(unsafe.Pointer(v.Addr().Pointer())), nil
	case int16:
		return "int16", (*int16Value)(unsafe.Pointer(v.Addr().Pointer())), nil
	case int8:
		return "int8", (*int8Value)(unsafe.Pointer(v.Addr().Pointer())), nil
	case uint64:
		return "uint64", (*uint64Value)(unsafe.Pointer(v.Addr().Pointer())), nil
	case uint32:
		return "uint32", (*uint32Value)(unsafe.Pointer(v.Addr().Pointer())), nil
	case uint16:
		return "uint16", (*uint16Value)(unsafe.Pointer(v.Addr().Pointer())), nil
	case uint8:
		return "uint8", (*uint8Value)(unsafe.Pointer(v.Addr().Pointer())), nil
	case float64:
		return "float64", (*float64Value)(unsafe.Pointer(v.Addr().Pointer())), nil
	case float32:
		return "float32", (*float32Value)(unsafe.Pointer(v.Addr().Pointer())), nil
	case time.Duration:
		return "duration", (*durationValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case time.Time:
		return "time", (*timeValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case net.IP:
		return "ip", (*ipValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case string:
		return "string", (*stringValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []bool:
		return "[]bool", (*boolSliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []int:
		return "[]int", (*intSliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []uint:
		return "[]uint", (*uintSliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []int64:
		return "[]int64", (*int64SliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []int32:
		return "[]int32", (*int32SliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []int16:
		return "[]int16", (*int16SliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []int8:
		return "[]int8", (*int8SliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []uint64:
		return "[]uint64", (*uint64SliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []uint32:
		return "[]uint32", (*uint32SliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []uint16:
		return "[]uint16", (*uint16SliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []uint8:
		return "[]uint8", (*uint8SliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []float64:
		return "[]float64", (*float64SliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []float32:
		return "[]float32", (*float32SliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []time.Duration:
		return "[]duration", (*durationSliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []time.Time:
		return "[]time", (*timeSliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []net.IP:
		return "[]ip", (*ipSliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	case []string:
		return "[]string", (*stringSliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil
	default:
		return "", nil, fmt.Errorf("unsupport type [%v]", v.Type())
	}
}

func (f *FlagSet) Bind(v interface{}) error {
	return f.bind(v, "")
}

func (f *FlagSet) bind(v interface{}, prefix string) error {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("expected a pointer, got [%v]", reflect.TypeOf(v))
	}

	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()

	if rt.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct, got [%v]", rt)
	}

	for i := 0; i < rv.NumField(); i++ {
		tag := rt.Field(i).Tag.Get("hflag")
		t := rt.Field(i).Type

		if tag == "-" {
			continue
		}

		typeStr, value, err := interfaceToType(rv.Field(i))
		if err == nil {
			name, shorthand, usage, required, defaultValue, position, err := parseTag(tag)
			if err != nil {
				return err
			}
			if name == "" {
				name = hstring.KebabName(rt.Field(i).Name)
			}
			if prefix != "" {
				name = prefix + "-" + name
			}
			if position {
				if err := f.addPosFlag(name, usage, typeStr, required, defaultValue); err != nil {
					return err
				}
			} else {
				if err := f.addFlag(name, usage, shorthand, typeStr, required, defaultValue); err != nil {
					return err
				}
			}
			f.nameToFlag[name].Value = value
			if defaultValue != "" {
				if err := f.nameToFlag[name].Set(defaultValue); err != nil {
					return err
				}
			}
		} else if t.Kind() == reflect.Struct {
			p := prefix
			if prefix != "" {
				p = prefix + "-" + tag
			} else {
				p = tag
			}
			if err := f.bind(rv.Field(i).Addr().Interface(), p); err != nil {
				return err
			}
		} else if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct {
			rv.Field(i).Set(reflect.New(rv.Field(i).Type().Elem()))
			p := prefix
			if prefix != "" {
				p = prefix + "-" + tag
			} else {
				p = tag
			}
			if err := f.bind(rv.Field(i).Interface(), p); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
