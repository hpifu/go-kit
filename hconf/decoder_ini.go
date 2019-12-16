package hconf

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type IniDecoder struct{}

func (d IniDecoder) escape(str string) (string, error) {
	var buf bytes.Buffer

	buf.WriteByte('"')
	i := 0
	for i < len(str) {
		if str[i] == '/' && i+1 < len(str) {
			switch str[i+1] {
			case ';':
				buf.WriteByte(';')
			case '#':
				buf.WriteByte('#')
			case '=':
				buf.WriteByte('=')
			case ':':
				buf.WriteByte(':')
			default:
				buf.WriteByte('\\')
				buf.WriteByte(str[i+1])
			}
			i += 2
		} else {
			buf.WriteByte(str[i])
			i++
		}
	}

	buf.WriteByte('"')
	return strconv.Unquote(buf.String())
}

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
			var err error
			if prefix, err = d.escape(line[1 : len(line)-1]); err != nil {
				return nil, err
			}
			continue
		}
		idx := strings.IndexAny(line, "=")
		key := strings.Trim(line[:idx], " ")
		val := strings.Trim(line[idx+1:], " ")
		if val[0] == '"' && val[len(val)-1] == '"' {
			val = val[1 : len(val)-1]
		}
		var err error
		if key, err = d.escape(key); err != nil {
			return nil, err
		}
		if val, err = d.escape(val); err != nil {
			return nil, err
		}
		if prefix != "" {
			key = prefix + "." + key
		}
		kvs[key] = val
	}

	return NewMapStorage(kvs), nil
}
