#!/usr/bin/env python3

slice_tpl = """func To{name}Slice(str string) ([]{type}, error) {{
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


def gen_slice_code(type):
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
    return slice_tpl.format(type=type, name=name)


def main():
    for type in [
        "string", "bool", "int", "uint", "int64", "int32", "int16", "int8",
        "uint64", "uint32", "uint16", "uint8", "float64", "float32",
        "time.Duration", "time.Time", "net.IP"
    ]:
        print(gen_slice_code(type))


if __name__ == "__main__":
    main()
