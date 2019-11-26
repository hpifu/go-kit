package hflag

import (
	"fmt"
	"strconv"
	"time"
)

func (f *FlagSet) Lookup(name string) *Flag {
	flag, ok := f.nameToFlag[name]
	if ok {
		return flag
	}
	k, ok := f.shorthandToName[name]
	if ok {
		return f.nameToFlag[k]
	}
	return nil
}

func (f *FlagSet) Set(name string, val string) error {
	flag := f.Lookup(name)
	if flag == nil {
		return fmt.Errorf("no such flag, Name [%v]", name)
	}

	return flag.Set(val)
}

func (f *FlagSet) Visit(callback func(f *Flag)) {
	for _, flag := range f.nameToFlag {
		if flag.Assigned {
			callback(flag)
		}
	}
}

func (f *FlagSet) VisitAll(callback func(f *Flag)) {
	for _, flag := range f.nameToFlag {
		callback(flag)
	}
}

func (f *FlagSet) Parsed() bool {
	return f.parsed
}

func (f *FlagSet) NArg() int {
	return len(f.args)
}

func (f *FlagSet) Args() []string {
	return f.args
}

func (f *FlagSet) Arg(i int) string {
	if i >= len(f.args) || i < 0 {
		return ""
	}

	return f.args[i]
}

func (f *FlagSet) NFlag() int {
	n := 0
	for _, flag := range f.nameToFlag {
		if flag.Assigned {
			n++
		}
	}
	return n
}

func (f *FlagSet) PrintDefaults() {
	fmt.Println(f.Usage())
}

func (f *FlagSet) Bool(name string, defaultValue bool, usage string) *bool {
	if err := f.addFlagAutoShorthand(name, usage, "bool", fmt.Sprintf("%v", defaultValue)); err != nil {
		panic(err)
	}
	return (*bool)(f.nameToFlag[name].Value.(*boolValue))
}

func (f *FlagSet) Int(name string, defaultValue int, usage string) *int {
	if err := f.addFlagAutoShorthand(name, usage, "int", strconv.Itoa(defaultValue)); err != nil {
		panic(err)
	}
	return (*int)(f.nameToFlag[name].Value.(*intValue))
}

func (f *FlagSet) Int64(name string, defaultValue int64, usage string) *int64 {
	if err := f.addFlagAutoShorthand(name, usage, "int64", strconv.FormatInt(defaultValue, 10)); err != nil {
		panic(err)
	}
	return (*int64)(f.nameToFlag[name].Value.(*int64Value))
}

func (f *FlagSet) Uint(name string, defaultValue uint, usage string) *uint {
	if err := f.addFlagAutoShorthand(name, usage, "uint", strconv.FormatUint(uint64(defaultValue), 10)); err != nil {
		panic(err)
	}
	return (*uint)(f.nameToFlag[name].Value.(*uintValue))
}

func (f *FlagSet) Uint64(name string, defaultValue uint64, usage string) *uint64 {
	if err := f.addFlagAutoShorthand(name, usage, "uint64", strconv.FormatUint(defaultValue, 10)); err != nil {
		panic(err)
	}
	return (*uint64)(f.nameToFlag[name].Value.(*uint64Value))
}

func (f *FlagSet) String(name string, defaultValue string, usage string) *string {
	if err := f.addFlagAutoShorthand(name, usage, "string", defaultValue); err != nil {
		panic(err)
	}
	return (*string)(f.nameToFlag[name].Value.(*stringValue))
}

func (f *FlagSet) Duration(name string, defaultValue time.Duration, usage string) *time.Duration {
	if err := f.addFlagAutoShorthand(name, usage, "duration", defaultValue.String()); err != nil {
		panic(err)
	}
	return (*time.Duration)(f.nameToFlag[name].Value.(*durationValue))
}

func (f *FlagSet) Float64(name string, defaultValue float64, usage string) *float64 {
	if err := f.addFlagAutoShorthand(name, usage, "float", fmt.Sprintf("%f", defaultValue)); err != nil {
		panic(err)
	}
	return (*float64)(f.nameToFlag[name].Value.(*floatValue))
}

func (f *FlagSet) IntSlice(name string, defaultValue []int, usage string) *[]int {
	if err := f.addFlagAutoShorthand(name, usage, "[]int", intSliceValue(defaultValue).String()); err != nil {
		panic(err)
	}
	return (*[]int)(f.nameToFlag[name].Value.(*intSliceValue))
}

func (f *FlagSet) StringSlice(name string, defaultValue []string, usage string) *[]string {
	if err := f.addFlagAutoShorthand(name, usage, "[]string", stringSliceValue(defaultValue).String()); err != nil {
		panic(err)
	}
	return (*[]string)(f.nameToFlag[name].Value.(*stringSliceValue))
}

func (f *FlagSet) BoolVar(v *bool, name string, defaultValue bool, usage string) {
	*v = defaultValue
	if err := f.addFlagAutoShorthand(name, usage, "bool", fmt.Sprintf("%v", defaultValue)); err != nil {
		panic(err)
	}
	f.nameToFlag[name].Value = (*boolValue)(v)
}

func (f *FlagSet) IntVar(v *int, name string, defaultValue int, usage string) {
	*v = defaultValue
	if err := f.addFlagAutoShorthand(name, usage, "int", strconv.Itoa(defaultValue)); err != nil {
		panic(err)
	}
	f.nameToFlag[name].Value = (*intValue)(v)
}

func (f *FlagSet) Int64Var(v *int64, name string, defaultValue int64, usage string) {
	*v = defaultValue
	if err := f.addFlagAutoShorthand(name, usage, "int64", strconv.FormatInt(defaultValue, 10)); err != nil {
		panic(err)
	}
	f.nameToFlag[name].Value = (*int64Value)(v)
}

func (f *FlagSet) UintVar(v *uint, name string, defaultValue uint, usage string) {
	*v = defaultValue
	if err := f.addFlagAutoShorthand(name, usage, "uint", strconv.FormatUint(uint64(defaultValue), 10)); err != nil {
		panic(err)
	}
	f.nameToFlag[name].Value = (*uintValue)(v)
}

func (f *FlagSet) Uint64Var(v *uint64, name string, defaultValue uint64, usage string) {
	*v = defaultValue
	if err := f.addFlagAutoShorthand(name, usage, "uint64", strconv.FormatUint(defaultValue, 10)); err != nil {
		panic(err)
	}
	f.nameToFlag[name].Value = (*uint64Value)(v)
}

func (f *FlagSet) StringVar(v *string, name string, defaultValue string, usage string) {
	*v = defaultValue
	if err := f.addFlagAutoShorthand(name, usage, "string", defaultValue); err != nil {
		panic(err)
	}
	f.nameToFlag[name].Value = (*stringValue)(v)
}

func (f *FlagSet) DurationVar(v *time.Duration, name string, defaultValue time.Duration, usage string) {
	*v = defaultValue
	if err := f.addFlagAutoShorthand(name, usage, "duration", defaultValue.String()); err != nil {
		panic(err)
	}
	f.nameToFlag[name].Value = (*durationValue)(v)
}

func (f *FlagSet) Float64Var(v *float64, name string, defaultValue float64, usage string) {
	*v = defaultValue
	if err := f.addFlagAutoShorthand(name, usage, "float", fmt.Sprintf("%f", defaultValue)); err != nil {
		panic(err)
	}
	f.nameToFlag[name].Value = (*floatValue)(v)
}

func (f *FlagSet) IntSliceVar(v *[]int, name string, defaultValue []int, usage string) {
	*v = defaultValue
	if err := f.addFlagAutoShorthand(name, usage, "[]int", intSliceValue(defaultValue).String()); err != nil {
		panic(err)
	}
	f.nameToFlag[name].Value = (*intSliceValue)(v)
}

func (f *FlagSet) StringSliceVar(v *[]string, name string, defaultValue []string, usage string) {
	*v = defaultValue
	if err := f.addFlagAutoShorthand(name, usage, "[]string", stringSliceValue(defaultValue).String()); err != nil {
		panic(err)
	}
	f.nameToFlag[name].Value = (*stringSliceValue)(v)
}

func (f *FlagSet) addFlagAutoShorthand(name string, usage string, typeStr string, defaultValue string) error {
	if len(name) == 1 {
		return f.AddFlag(name, usage, Type(typeStr), DefaultValue(defaultValue), Shorthand(name))
	}

	return f.AddFlag(name, usage, Type(typeStr), DefaultValue(defaultValue))
}
