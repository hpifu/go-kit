package hconf

import (
	"fmt"
	"github.com/hpifu/go-kit/hstring"
	"github.com/spf13/cast"
	"net"
	"reflect"
	"time"
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
	return interfaceToStruct(s.data, v)
}

func interfaceToStruct(d interface{}, v interface{}) error {
	if d == nil {
		return nil
	}
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("invalid value type, expect a pointer, got %v", reflect.TypeOf(v))
	}

	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()
	switch rt.Kind() {
	case reflect.Struct:
		dv, ok := d.(map[string]interface{})
		if !ok {
			return fmt.Errorf("convert data to map[string]interface{} failed. which is %v", reflect.TypeOf(d))
		}
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			key := rt.Field(i).Tag.Get("hconf")
			if key == "-" {
				continue
			}
			if key == "" {
				key = hstring.CamelName(rt.Field(i).Name)
			}
			value := dv[key]
			if rt.Field(i).Type.Kind() == reflect.Ptr {
				if field.IsNil() {
					nv := reflect.New(field.Type().Elem())
					field.Set(nv)
				}
				if err := interfaceToStruct(value, field.Interface()); err != nil {
					return err
				}
			} else {
				if err := interfaceToStruct(value, field.Addr().Interface()); err != nil {
					return err
				}
			}
		}
	case reflect.Slice:
		dv, ok := d.([]interface{})
		if !ok {
			return fmt.Errorf("convert data to []interface{} failed. which is %v", reflect.TypeOf(d))
		}
		rv.Set(reflect.MakeSlice(rt, 0, rv.Cap()))
		for _, di := range dv {
			nv := reflect.New(rt.Elem())
			err := interfaceToStruct(di, nv.Interface())
			if err != nil {
				return err
			}
			rv.Set(reflect.Append(rv, nv.Elem()))
		}
	default:
		switch rv.Interface().(type) {
		case string:
			v, err := cast.ToStringE(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case bool:
			v, err := cast.ToBoolE(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case int:
			v, err := cast.ToIntE(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case uint:
			v, err := cast.ToUintE(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case int64:
			v, err := cast.ToInt64E(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case int32:
			v, err := cast.ToInt32E(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case int16:
			v, err := cast.ToInt16E(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case int8:
			v, err := cast.ToInt8E(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case uint64:
			v, err := cast.ToUint64E(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case uint32:
			v, err := cast.ToUint32E(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case uint16:
			v, err := cast.ToUint16E(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case uint8:
			v, err := cast.ToUint8E(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case float64:
			v, err := cast.ToFloat64E(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case float32:
			v, err := cast.ToFloat32E(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case time.Duration:
			v, err := cast.ToDurationE(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case time.Time:
			v, err := cast.ToTimeE(d)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case net.IP:
			switch v.(type) {
			case string:
				v, err := hstring.ToIP(v.(string))
				if err != nil {
					return err
				}
				rv.Set(reflect.ValueOf(v))
			case net.IP:
				rv.Set(reflect.ValueOf(v.(net.IP)))
			default:
				return fmt.Errorf("convert type [%v] to ip failed", reflect.TypeOf(v))
			}
		default:
			return fmt.Errorf("unsupport type %v", rt)
		}
	}

	return nil
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
