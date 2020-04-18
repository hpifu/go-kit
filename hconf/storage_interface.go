package hconf

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hpifu/go-kit/href"
)

func NewInterfaceStorage(data interface{}) *InterfaceStorage {
	return &InterfaceStorage{
		data: data,
	}
}

type InterfaceStorage struct {
	data interface{}
}

func (s InterfaceStorage) Get(key string) (interface{}, error) {
	data := s.data
	infos, err := parseKey(key, ".")
	if err != nil {
		return nil, err
	}

	for _, info := range infos {
		if info.mod == MapMod {
			var ok bool
			switch data.(type) {
			case map[string]interface{}:
				if data, ok = data.(map[string]interface{})[info.key]; !ok {
					return nil, fmt.Errorf("no such key [%v]", key)
				}
			case map[interface{}]interface{}:
				if data, ok = data.(map[interface{}]interface{})[info.key]; !ok {
					return nil, fmt.Errorf("no such key [%v]", key)
				}
			default:
				return nil, fmt.Errorf("data is not a map. type: [%v], data: %v", reflect.TypeOf(data), data)
			}
		} else {
			switch data.(type) {
			case []interface{}:
				val := data.([]interface{})
				if len(val) <= info.idx {
					return nil, fmt.Errorf("index out of bounds. index: [%v], data: [%v]", info.idx, data)
				}
				data = val[info.idx]
			case []map[string]interface{}:
				val := data.([]map[string]interface{})
				if len(val) <= info.idx {
					return nil, fmt.Errorf("index out of bounds. index: [%v], data: [%v]", info.idx, data)
				}
				data = val[info.idx]
			//case []map[interface{}]interface{}:
			default:
				return nil, fmt.Errorf("data is not a array. type: [%v], data: %v", reflect.TypeOf(data), data)
			}
		}
	}

	return data, nil
}

func (s *InterfaceStorage) Set(key string, val interface{}) error {
	data := s.data
	infos, err := parseKey(key, ".")
	if err != nil {
		return err
	}

	for i, info := range infos {
		if info.mod == MapMod {
			v, ok := data.(map[string]interface{})
			if !ok {
				return fmt.Errorf("data is not a map. data: [%v]", data)
			}
			if i == len(infos)-1 {
				v[info.key] = val
			} else {
				if data, ok = v[info.key]; !ok {
					v[info.key] = map[string]interface{}{}
					data = v[info.key]
				}
			}
		} else {
			v, ok := data.([]interface{})
			if !ok {
				return fmt.Errorf("data is not a array. data: [%v]", data)
			}
			if len(v) <= info.idx {
				return fmt.Errorf("index out of bounds. index: [%v], data: [%v]", info.idx, data)
			}
			if i == len(infos)-1 {
				v[info.idx] = val
			} else {
				data = v[info.idx]
			}
		}
	}

	return nil
}

func (s InterfaceStorage) Unmarshal(v interface{}) error {
	return href.InterfaceToStruct(s.data, v)
}

func (s InterfaceStorage) Sub(key string) (Storage, error) {
	v, err := s.Get(key)
	if err != nil {
		return nil, err
	}

	return &InterfaceStorage{
		data: v,
	}, nil
}

func parseKey(keys string, separator string) ([]*KeyInfo, error) {
	var infos []*KeyInfo
	for _, key := range strings.Split(keys, separator) {
		fields := strings.Split(key, "[")
		if len(fields[0]) != 0 {
			infos = append(infos, &KeyInfo{
				key: fields[0],
				mod: MapMod,
			})
		}
		for i := 1; i < len(fields); i++ {
			if !strings.HasSuffix(fields[i], "]") {
				return nil, fmt.Errorf("invalid format. key: [%v]", key)
			}
			field := fields[i][0 : len(fields[i])-1]
			if len(field) == 0 {
				return nil, fmt.Errorf("index should not be empty. key: [%v]", key)
			}
			idx, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("index should be a number. key: [%v]", key)
			}
			infos = append(infos, &KeyInfo{
				idx: idx,
				mod: ArrMod,
			})
		}
	}

	return infos, nil
}
