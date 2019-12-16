package hconf

import (
	"fmt"
	"strings"
)

type PropDecoder struct{}

func (d *PropDecoder) Decode(buf []byte) (Storage, error) {
	kvs := map[string]string{}

	lines := strings.FieldsFunc(string(buf), func(r rune) bool {
		return r == '\n' || r == '\r'
	})
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		if line[0] == '!' || line[0] == '#' {
			continue
		}
		for line[len(line)-1] == '\\' && i+1 < len(lines) {
			i++
			nextLine := strings.Trim(lines[i], " \t")
			line = line[:len(line)-1] + nextLine
		}
		idx := strings.IndexAny(line, ":=")
		key := strings.Trim(line[:idx], " ")
		val := strings.Trim(line[idx+1:], " ")
		kvs[key] = val
	}
	fmt.Println(kvs)

	return NewMapStorage(kvs), nil
}
