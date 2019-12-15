package hconf

import "gopkg.in/yaml.v2"

type PropDecoder struct{}

func (d *PropDecoder) Decode(buf []byte) (Storage, error) {
	var data interface{}
	if err := yaml.Unmarshal(buf, &data); err != nil {
		return nil, err
	}
	return NewInterfaceStorage(data), nil
}
