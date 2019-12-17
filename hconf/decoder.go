package hconf

import "fmt"

func NewDecoder(name string) (Decoder, error) {
	switch name {
	case "json", "json5":
		return &Json5Decoder{}, nil
	case "yml", "yaml":
		return &YamlDecoder{}, nil
	case "toml":
		return &TomlDecoder{}, nil
	case "properties", "prop":
		return &PropDecoder{}, nil
	case "xml":
		return &XmlDecoder{}, nil
	case "ini", "cfg":
		return &IniDecoder{}, nil
	}

	return nil, fmt.Errorf("unsupport decoder. name: [%v]", name)
}

type Decoder interface {
	Decode(buf []byte) (Storage, error)
}
