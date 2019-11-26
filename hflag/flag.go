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

func (f *FlagSet) NArg() int {
	return len(f.args)
}

func (f *FlagSet) Args() []string {
	return f.args
}

func (f *FlagSet) Arg(i int) string {
	if i >= len(f.args) {
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

func (f *FlagSet) BoolVar(b *bool, name string, defaultValue bool, help string) {
	*b = defaultValue
	err := f.AddFlag(name, "", help, "bool", false, fmt.Sprintf("%v", defaultValue))
	if err != nil {
		panic(err)
	}
	f.nameToFlag[name].Value = (*boolValue)(b)
}

func (f *FlagSet) Bool(name string, defaultValue bool, help string) *bool {
	err := f.AddFlag(name, "", help, "bool", false, fmt.Sprintf("%v", defaultValue))
	if err != nil {
		panic(err)
	}
	return (*bool)(f.nameToFlag[name].Value.(*boolValue))
}

func (f *FlagSet) Int(name string, defaultValue int, help string) *int {
	if err := f.AddFlag(name, "", help, "int", false, strconv.Itoa(defaultValue)); err != nil {
		panic(err)
	}
	return (*int)(f.nameToFlag[name].Value.(*intValue))
}

func (f *FlagSet) String(name string, defaultValue string, help string) *string {
	if err := f.AddFlag(name, "", help, "string", false, defaultValue); err != nil {
		panic(err)
	}
	return (*string)(f.nameToFlag[name].Value.(*stringValue))
}

func (f *FlagSet) Duration(name string, defaultValue time.Duration, help string) *time.Duration {
	if err := f.AddFlag(name, "", help, "duration", false, defaultValue.String()); err != nil {
		panic(err)
	}
	return (*time.Duration)(f.nameToFlag[name].Value.(*durationValue))
}

func (f *FlagSet) Float(name string, defaultValue float64, help string) *float64 {
	if err := f.AddFlag(name, "", help, "float", false, fmt.Sprintf("%f", defaultValue)); err != nil {
		panic(err)
	}
	return (*float64)(f.nameToFlag[name].Value.(*floatValue))
}
