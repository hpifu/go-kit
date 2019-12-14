#!/usr/bin/env python3

value_set_tpl = """
func (v *{vtype}Value) Set(str string) error {{
	val, err := hstring.To{name}(str)
	if err != nil {{
		return err
	}}
	*v = {vtype}Value(val)
	return nil
}}
"""

value_string_tpl = """
func (v {vtype}Value) String() string {{
	return hstring.{name}To({type}(v))
}}
"""


new_value_type_tpl = """
func NewValueType(typeStr string) Value {{
	switch typeStr {{{body}
	default:
		return nil
	}}
}}
"""

interface_to_type_tpl = """
func interfaceToType(v reflect.Value) (string, Value, error) {{
	switch v.Interface().(type) {{{body}
	default:
		return "", nil, fmt.Errorf("unsupport type [%v]", v.Type())
	}}
}}
"""


flagset_get_tpl = """
func (f *FlagSet) Get{name}(name string) (v {type}) {{
	flag := f.Lookup(name)
	if flag == nil {{
		return
	}}
	if flag.Type == "string" {{
		val, err := hstring.To{name}(flag.Value.String())
		if err != nil {{
			return
		}}
		return val
	}}
	if flag.Type != "{vtype}" {{
		return
	}}
	return {type}(*flag.Value.(*{vtype}Value))
}}
"""

flagset_get_slice_tpl = """
func (f *FlagSet) Get{name}Slice(name string) (v []{type}) {{
	flag := f.Lookup(name)
	if flag == nil {{
		return
	}}
	if flag.Type == "string" {{
		val, err := hstring.To{name}Slice(flag.Value.String())
		if err != nil {{
			return
		}}
		return val
	}}
	if flag.Type != "[]{vtype}" {{
		return
	}}
	return []{type}(*flag.Value.(*{vtype}SliceValue))
}}
"""

flagset_type_tpl = """
func (f *FlagSet) {name}(name string, defaultValue {type}, usage string) *{type} {{
	if err := f.addFlagAutoShorthand(name, usage, "{vtype}", hstring.{name}To(defaultValue)); err != nil {{
		panic(err)
	}}
	return (*{type})(f.nameToFlag[name].Value.(*{vtype}Value))
}}
"""

flagset_slice_type_tpl = """
func (f *FlagSet) {name}Slice(name string, defaultValue []{type}, usage string) *[]{type} {{
	if err := f.addFlagAutoShorthand(name, usage, "[]{vtype}", hstring.{name}SliceTo(defaultValue)); err != nil {{
		panic(err)
	}}
	return (*[]{type})(f.nameToFlag[name].Value.(*{vtype}SliceValue))
}}
"""

flagset_type_var_tpl = """
func (f *FlagSet) {name}Var(v *{type}, name string, defaultValue {type}, usage string) {{
	*v = defaultValue
	if err := f.addFlagAutoShorthand(name, usage, "{vtype}", hstring.{name}To(defaultValue)); err != nil {{
		panic(err)
	}}
	f.nameToFlag[name].Value = (*{vtype}Value)(v)
}}
"""

flagset_slice_type_var_tpl = """
func (f *FlagSet) {name}SliceVar(v *[]{type}, name string, defaultValue []{type}, usage string) {{
	*v = defaultValue
	if err := f.addFlagAutoShorthand(name, usage, "[]{vtype}", hstring.{name}SliceTo(defaultValue)); err != nil {{
		panic(err)
	}}
	f.nameToFlag[name].Value = (*{vtype}SliceValue)(v)
}}
"""

commandline_type_tpl = """
func {name}(name string, defaultValue {type}, usage string) *{type} {{
	return CommandLine.{name}(name, defaultValue, usage)
}}
"""

commandline_type_var_tpl = """
func {name}Var(v *{type}, name string, defaultValue {type}, usage string) {{
	CommandLine.{name}Var(v, name, defaultValue, usage)
}}
"""

commandline_get_type_tpl = """
func Get{name}(name string) {type} {{
	return CommandLine.Get{name}(name)
}}
"""

unmarshal_tpl = """		case {type}:
			if fl.Type == "string" {{
				v, err := hstring.To{name}(string(*fl.Value.(*stringValue)))
				if err != nil {{
					return err
				}}
				rv.Set(reflect.ValueOf(v))
			}} else if fl.Type == "{vtype}" {{
				rv.Set(reflect.ValueOf({type}(*fl.Value.(*{vtype}Value))))
			}} else {{
				return fmt.Errorf("expect a {vtype}, got [%v]", fl.Type)
			}}"""

unmarshal_slice_tpl = """		case []{type}:
			if fl.Type == "string" {{
				v, err := hstring.To{name}Slice(string(*fl.Value.(*stringValue)))
				if err != nil {{
					return err
				}}
				rv.Set(reflect.ValueOf(v))
			}} else if fl.Type == "[]{vtype}" {{
				rv.Set(reflect.ValueOf([]{type}(*fl.Value.(*{vtype}SliceValue))))
			}} else {{
				return fmt.Errorf("expect a []{vtype}, got [%v]", fl.Type)
			}}"""


def vtype(type):
    return type.split(".")[-1].lower()


def name(type):
    temp = type.split(".")[-1]
    return temp[0].upper() + temp[1:]


def gen_define_value(type):
    return "type {vtype}Value {type}".format(vtype=vtype(type), type=type)


def gen_define_slice_value(type):
    return "type {vtype}SliceValue []{type}".format(vtype=vtype(type), type=type)


def gen_value_set(type):
    if type == "string":
        return """
func (v *stringValue) Set(str string) error {
	*v = stringValue(str)
	return nil
}
"""
    return value_set_tpl.format(vtype=vtype(type), name=name(type))


def gen_slice_value_set(type):
    return value_set_tpl.format(vtype=vtype(type) + "Slice", name=name(type) + "Slice")


def gen_value_string(type):
    if type == "string":
        return """
func (v stringValue) String() string {
	return string(v)
}
        """
    return value_string_tpl.format(vtype=vtype(type), type=type, name=name(type))


def gen_slice_value_string(type):
    return value_string_tpl.format(vtype=vtype(type) + "Slice", type="[]" + type, name=name(type) + "Slice")


def gen_new_value_type(types):
    tpl = """
	case "{vtype}":
		return new({vtype}Value)"""
    body = ""
    for type in types:
        body += tpl.format(vtype=vtype(type))
    tpl = """
	case "[]{vtype}":
		return new({vtype}SliceValue)"""
    for type in types:
        body += tpl.format(vtype=vtype(type))

    return new_value_type_tpl.format(body=body)


def gen_interface_to_type(types):
    body = ""
    tpl = """
	case {type}:
		return "{vtype}", (*{vtype}Value)(unsafe.Pointer(v.Addr().Pointer())), nil"""
    for type in types:
        body += tpl.format(type=type, vtype=vtype(type))
    tpl = """
	case []{type}:
		return "[]{vtype}", (*{vtype}SliceValue)(unsafe.Pointer(v.Addr().Pointer())), nil"""
    for type in types:
        body += tpl.format(type=type, vtype=vtype(type))
    return interface_to_type_tpl.format(body=body)


def gen_flagset_get_tpl(type):
    if type == "string":
        return """
func (f *FlagSet) GetString(name string) string {
	flag := f.Lookup(name)
	if flag == nil {
		return ""
	}
	return flag.Value.String()
}
"""
    return flagset_get_tpl.format(name=name(type), type=type, vtype=vtype(type))


def gen_flagset_get_slice_tpl(type):
    return flagset_get_slice_tpl.format(name=name(type), type=type, vtype=vtype(type))


def gen_flagset_type_tpl(type):
    return flagset_type_tpl.format(name=name(type), type=type, vtype=vtype(type))


def gen_flagset_slice_type_tpl(type):
    return flagset_slice_type_tpl.format(name=name(type), type=type, vtype=vtype(type))


def gen_flagset_type_var_tpl(type):
    return flagset_type_var_tpl.format(name=name(type), type=type, vtype=vtype(type))


def gen_flagset_slice_type_var_tpl(type):
    return flagset_slice_type_var_tpl.format(name=name(type), type=type, vtype=vtype(type))


def gen_commandline_type(type):
    return commandline_type_tpl.format(type=type, name=name(type))


def gen_commandline_slice_type(type):
    return commandline_type_tpl.format(type="[]" + type, name=name(type) + "Slice")


def gen_commandline_type_var(type):
    return commandline_type_var_tpl.format(type=type, name=name(type))


def gen_commandline_slice_type_var(type):
    return commandline_type_var_tpl.format(type="[]" + type, name=name(type) + "Slice")


def gen_commandline_get_type(type):
    return commandline_get_type_tpl.format(type=type, name=name(type))


def gen_commandline_get_slice_type(type):
    return commandline_get_type_tpl.format(type="[]" + type, name=name(type) + "Slice")


def gen_unmarshal(type):
    return unmarshal_tpl.format(type=type, name=name(type), vtype=vtype(type))


def gen_unmarshal_slice(type):
    return unmarshal_slice_tpl.format(type=type, name=name(type), vtype=vtype(type))


def main():
    types = [
        "bool", "int", "uint", "int64", "int32", "int16", "int8",
        "uint64", "uint32", "uint16", "uint8", "float64", "float32",
        "time.Duration", "time.Time", "net.IP", "string"
    ]

    # value.go
    # for type in types:
    #     print(gen_define_value(type))
    # for type in types:
    #     print(gen_define_slice_value(type))
    # for type in types:
    #     print(gen_value_set(type))
    # for type in types:
    #     print(gen_slice_value_set(type))
    # for type in types:
    #     print(gen_value_string(type))
    # for type in types:
    #     print(gen_slice_value_string(type))
    # print(gen_new_value_type(types))

    # hflag.go
    # print(gen_interface_to_type(types))

    # get.go
    # for type in types:
    #     print(gen_flagset_get_tpl(type))
    # for type in types:
    #     print(gen_flagset_get_slice_tpl(type))

    # flag.go
    # for type in types:
    #     print(gen_flagset_type_tpl(type))
    # for type in types:
    #     print(gen_flagset_slice_type_tpl(type))
    # for type in types:
    #     print(gen_flagset_type_var_tpl(type))
    # for type in types:
    #     print(gen_flagset_slice_type_var_tpl(type))

    # commandline.go
    # for type in types:
    #     print(gen_commandline_type(type))
    # for type in types:
    #     print(gen_commandline_slice_type(type))
    # for type in types:
    #     print(gen_commandline_type_var(type))
    # for type in types:
    #     print(gen_commandline_slice_type_var(type))
    # for type in types:
    #     print(gen_commandline_get_type(type))
    # for type in types:
    #     print(gen_commandline_get_slice_type(type))

    for type in types:
        print(gen_unmarshal(type))
    for type in types:
        print(gen_unmarshal_slice(type))


if __name__ == "__main__":
    main()
