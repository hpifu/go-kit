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
    print(gen_interface_to_type(types))


if __name__ == "__main__":
    main()
