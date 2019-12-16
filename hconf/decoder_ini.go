package hconf

import (
	"fmt"
	"strings"
)

type IniDecoder struct{}

func (d *IniDecoder) Decode(buf []byte) (Storage, error) {
	kvs := map[string]string{}

	lines := strings.FieldsFunc(string(buf), func(r rune) bool {
		return r == '\n' || r == '\r'
	})
	prefix := ""
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		if line[0] == ';' {
			continue
		}
		if line[0] == '[' {
			if line[len(line)-1] != ']' {
				return nil, fmt.Errorf("parse section [%v] failed", line[0])
			}
			prefix = line[1 : len(line)-1]
			continue
		}
		idx := strings.IndexAny(line, "=")
		key := strings.Trim(line[:idx], " ")
		val := strings.Trim(line[idx+1:], " ")
		if val[0] == '"' && val[len(val)-1] == '"' {
			val = val[1 : len(val)-1]
		}
		if prefix != "" {
			key = prefix + "." + key
		}
		kvs[key] = val
	}

	return NewMapStorage(kvs), nil
}
