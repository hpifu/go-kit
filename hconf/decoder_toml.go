package hconf

import (
	"github.com/BurntSushi/toml"
)

type TomlDecoder struct{}

func (d *TomlDecoder) Decode(buf []byte) (Storage, error) {
	var data interface{}
	if _, err := toml.Decode(string(buf), &data); err != nil {
		return nil, err
	}
	return NewInterfaceStorage(data)
}
