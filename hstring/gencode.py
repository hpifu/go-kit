#!/usr/bin/env python3

string_to_slice_tpl = """func To{name}Slice(str string) ([]{type}, error) {{
	vals := strings.Split(str, ",")
	res := make([]{type}, 0, len(vals))
	for _, val := range vals {{
		v, err := To{name}(val)
		if err != nil {{
			return nil, err
		}}
		res = append(res, v)
	}}
	return res, nil
}}
"""

slice_to_string_tpl = """func {name}SliceTo(vs []{type}) string {{
	var buffer bytes.Buffer
	for idx, v := range vs {{
		buffer.WriteString({name}To(v))
		if idx != len(vs)-1 {{
			buffer.WriteString(",")
		}}
	}}
	return buffer.String()
}}
"""

to_string_tpl = """func ToString(v interface{{}}) string {{
	switch v.(type) {{{body}
	default:
		return fmt.Sprintf("%v", v)
	}}
}}"""


def gen_to_string(types):
    tpl = """
	case {type}:
		return {name}To(v.({type}))"""
    body = ""
    for type in types:
        if type == "string":
            continue
        temp = type.split(".")[-1]
        name = temp[0].upper() + temp[1:]
        body += tpl.format(type=type, name=name)
    return to_string_tpl.format(body=body)


def gen_slice_to_string(type):
    if type == "string":
        return """
func StringSliceTo(vs []string) string {
	return strings.Join(vs, ",")
}
"""

    temp = type.split(".")[-1]
    name = temp[0].upper() + temp[1:]
    return slice_to_string_tpl.format(type=type, name=name)


def gen_string_to_slice(type):
    if type == "string":
        return """
func ToStringSlice(str string) ([]string, error) {
	if str == "" {
		return []string{}, nil
	}
	return strings.Split(str, ","), nil
}
"""
    temp = type.split(".")[-1]
    name = temp[0].upper() + temp[1:]
    return string_to_slice_tpl.format(type=type, name=name)


def main():
    types = [
        "string", "bool", "int", "uint", "int64", "int32", "int16", "int8",
        "uint64", "uint32", "uint16", "uint8", "float64", "float32",
        "time.Duration", "time.Time", "net.IP"
    ]
    # print(gen_to_string(types))
    # for type in types:
    #     print(gen_string_to_slice(type))
    for type in types:
        print(gen_slice_to_string(type))


if __name__ == "__main__":
    main()
