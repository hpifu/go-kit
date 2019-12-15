package hconf

import (
	"bytes"
	"github.com/yosuke-furukawa/json5/encoding/json5"
)

type Json5Decoder struct{}

func (d *Json5Decoder) Decode(buf []byte) (Storage, error) {
	var data interface{}
	if err := json5.NewDecoder(bytes.NewReader(buf)).Decode(&data); err != nil {
		return nil, err
	}
	return NewInterfaceStorage(data), nil
}
