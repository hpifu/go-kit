#!/usr/bin/env python3

value_set_tpl = """
func (v *{type}Value) Set(str string) error {{
	val, err := hstring.To{name}(str)
	if err != nil {{
		return err
	}}
	*v = {type}Value(val)
	return nil
}}
"""


def type_name_map(type):
    map = {
        "ip": "IP",
        "ipSlice": "IPSlice",
    }

    if type in map:
        return map[type]
    return type.capitalize()


def gen_type_def_code(type):
    if type == "string":
        return """
func (v *stringValue) Set(str string) error {
	*v = stringValue(str)
	return nil
}
"""
    return "type {type}Value {otype}".format(type=type.split(".")[-1].lower(), otype=type)


def gen_slice_type_def_code(type):
    return "type {type}SliceValue []{otype}".format(type=type.split(".")[-1].lower(), otype=type)


def gen_value_set_code(type):
    temp = type.split(".")[-1]
    name = temp[0].upper() + temp[1:]
    type = temp.lower()
    return value_set_tpl.format(type=type, name=name)


def gen_value_slice_set_code(type):
    temp = type.split(".")[-1]
    name = temp[0].upper() + temp[1:] + "Slice"
    type = temp.lower() + "Slice"
    return value_set_tpl.format(type=type, name=name)


def main():
    types = [
        "bool", "int", "uint", "int64", "int32", "int16", "int8",
        "uint64", "uint32", "uint16", "uint8", "float64", "float32",
        "time.Duration", "time.Time", "net.IP", "string"
        # "boolSlice", "intSlice", "int64Slice", "int32Slice", "int16Slice", "int8Slice",
        # "uint64Slice", "uint32Slice", "uint16Slice", "uint8Slice", "float64Slice", "float32Slice",
        # "durationSlice", "timeSlice", "ipSlice"
    ]
    for type in types:
        print(gen_type_def_code(type))
    for type in types:
        print(gen_slice_type_def_code(type))
    for type in types:
        print(gen_value_set_code(type))
    for type in types:
        print(gen_value_slice_set_code(type))


if __name__ == "__main__":
    main()
