package hconf

import (
	"fmt"
	"github.com/hpifu/go-kit/hstring"
	"reflect"
)

func NewMapStorage(kvs map[string]string) *MapStorage {
	return &MapStorage{
		kvs: kvs,
	}
}

type MapStorage struct {
	prefix string
	kvs    map[string]string
}

func (s MapStorage) Get(key string) (interface{}, error) {
	if s.prefix != "" {
		key = s.prefix + "." + key
	}
	val, ok := s.kvs[key]
	if !ok {
		return nil, fmt.Errorf("no such key [%v]", key)
	}
	return val, nil
}

func (s *MapStorage) Set(key string, val interface{}) error {
	if s.prefix != "" {
		key = s.prefix + "." + key
	}
	s.kvs[key] = hstring.ToString(val)
	return nil
}

func (s MapStorage) Unmarshal(v interface{}) error {
	return mapToStruct(s.kvs, v, s.prefix)
}

func (s MapStorage) Sub(key string) (Storage, error) {
	return &MapStorage{
		kvs:    s.kvs,
		prefix: key,
	}, nil
}

func mapToStruct(kvs map[string]string, v interface{}, prefix string) error {
	if kvs == nil {
		return nil
	}

	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("invalid value type, expect a pointer, got %v", reflect.TypeOf(v))
	}

	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()

	if rt.Kind() == reflect.Struct {
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			key := rt.Field(i).Tag.Get("hconf")
			if key == "-" {
				continue
			}
			if key == "" {
				key = hstring.CamelName(rt.Field(i).Name)
			}
			if prefix != "" {
				key = prefix + "." + key
			}
			if rt.Field(i).Type.Kind() == reflect.Ptr {
				if field.IsNil() {
					nv := reflect.New(field.Type().Elem())
					field.Set(nv)
				}
				if err := mapToStruct(kvs, field.Interface(), key); err != nil {
					return err
				}
			} else {
				if err := mapToStruct(kvs, field.Addr().Interface(), key); err != nil {
					return err
				}
			}
		}

		return nil
	}

	val, ok := kvs[prefix]
	if !ok {
		return nil
	}

	err := hstring.SetValue(rv, val)
	if err != nil {
		return err
	}

	return nil
}
