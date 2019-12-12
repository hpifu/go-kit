#!/usr/bin/env python3

slice_tpl = """
func To{Type}Slice(str string) ([]{type}, error) {{
	vals := strings.Split(str, ",")
	res := make([]{type}, 0, len(vals))
	for _, val := range vals {{
		v, err := To{Type}(val)
		if err != nil {{
			return nil, err
		}}
		res = append(res, v)
	}}
	return res, nil
}}
"""


def type_map(type):
    map = {
        "time.Duration": "Duration",
        "time.Time": "Time",
        "net.IP": "IP",
    }

    if type in map:
        return map[type]
    return type.capitalize()


def gen_slice_code(type):
    return slice_tpl.format(type=type, Type=type_map(type))


def main():
    for type in [
        "bool", "int", "uint", "int64", "int32", "int16", "int8",
        "uint64", "uint32", "uint16", "uint8", "float64", "float32",
        "time.Duration", "time.Time", "net.IP"
    ]:
        print(gen_slice_code(type))


if __name__ == "__main__":
    main()
