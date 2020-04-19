package href

import (
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/spf13/cast"

	"github.com/hpifu/go-kit/hstr"
)

func MapStringInterfaceToStruct(rt reflect.Type, rv reflect.Value, dv map[string]interface{}) error {
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		key := rt.Field(i).Tag.Get("hconf")
		if key == "-" {
			continue
		}
		if key == "" {
			key = hstr.CamelName(rt.Field(i).Name)
		}
		value, ok := dv[key]
		if !ok {
			continue
		}
		if rt.Field(i).Type.Kind() == reflect.Ptr {
			if field.IsNil() {
				nv := reflect.New(field.Type().Elem())
				field.Set(nv)
			}
			if err := InterfaceToStruct(value, field.Interface()); err != nil {
				return fmt.Errorf("key: [%v], err: [%v]", key, err)
			}
		} else if rt.Field(i).Type.Kind() == reflect.Interface {
			field.Set(reflect.ValueOf(value))
		} else {
			if err := InterfaceToStruct(value, field.Addr().Interface()); err != nil {
				return fmt.Errorf("key: [%v], err: [%v]", key, err)
			}
		}
	}

	return nil
}

func MapInterfaceInterfaceToStruct(rt reflect.Type, rv reflect.Value, dv map[interface{}]interface{}) error {
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		key := rt.Field(i).Tag.Get("hconf")
		if key == "-" {
			continue
		}
		if key == "" {
			key = hstr.CamelName(rt.Field(i).Name)
		}
		value, ok := dv[key]
		if !ok {
			continue
		}
		if rt.Field(i).Type.Kind() == reflect.Ptr {
			if field.IsNil() {
				nv := reflect.New(field.Type().Elem())
				field.Set(nv)
			}
			if err := InterfaceToStruct(value, field.Interface()); err != nil {
				return fmt.Errorf("key: [%v], err: [%v]", key, err)
			}
		} else if rt.Field(i).Type.Kind() == reflect.Interface {
			field.Set(reflect.ValueOf(value))
		} else {
			if err := InterfaceToStruct(value, field.Addr().Interface()); err != nil {
				return fmt.Errorf("key: [%v], err: [%v]", key, err)
			}
		}
	}

	return nil
}

func InterfaceToStruct(d interface{}, v interface{}) error {
	if d == nil {
		return nil
	}
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("invalid value type, expect a pointer, got %v", reflect.TypeOf(v))
	}

	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()
	switch rt.Kind() {
	case reflect.Map:
		switch rv.Interface().(type) {
		case map[string]string:
			if dv, ok := d.(map[string]string); ok {
				rv.Set(reflect.ValueOf(dv))
			} else if dv, ok := d.(map[string]interface{}); ok {
				m := map[string]string{}
				for key, val := range dv {
					v, err := cast.ToStringE(val)
					if err != nil {
						return fmt.Errorf("value type is not string, key: [%v], val: [%v]", key, val)
					}
					m[key] = v
				}
				rv.Set(reflect.ValueOf(m))
			} else if dv, ok := d.(map[interface{}]interface{}); ok {
				m := map[string]string{}
				for key, val := range dv {
					k, err := cast.ToStringE(key)
					if err != nil {
						return fmt.Errorf("value type is not string, key: [%v], val: [%v]", key, val)
					}
					v, err := cast.ToStringE(val)
					if err != nil {
						return fmt.Errorf("value type is not string, key: [%v], val: [%v]", key, val)
					}
					m[k] = v
				}
				rv.Set(reflect.ValueOf(m))
			} else {
				return fmt.Errorf("unsupport data type [%v]", reflect.ValueOf(d))
			}
		case map[string]interface{}:
			if dv, ok := d.(map[string]string); ok {
				m := map[string]interface{}{}
				for key, val := range dv {
					m[key] = val
				}
				rv.Set(reflect.ValueOf(m))
			} else if dv, ok := d.(map[string]interface{}); ok {
				rv.Set(reflect.ValueOf(dv))
			} else if dv, ok := d.(map[interface{}]interface{}); ok {
				m := map[string]interface{}{}
				for key, val := range dv {
					k, err := cast.ToStringE(key)
					if err != nil {
						return fmt.Errorf("value type is not string, key: [%v], val: [%v]", key, val)
					}
					m[k] = val
				}
				rv.Set(reflect.ValueOf(m))
			} else {
				return fmt.Errorf("unsupport data type [%v]", reflect.ValueOf(d))
			}
		default:
			return fmt.Errorf("unsupport type [%v]", rt)
		}
	case reflect.Struct:
		if dv, ok := d.(map[string]interface{}); ok {
			if err := MapStringInterfaceToStruct(rt, rv, dv); err != nil {
				return err
			}
		} else if dv, ok := d.(map[interface{}]interface{}); ok {
			if err := MapInterfaceInterfaceToStruct(rt, rv, dv); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unsupport data type: [%v]", reflect.TypeOf(d))
		}
	case reflect.Slice:
		dv, ok := d.([]interface{})
		if !ok {
			return fmt.Errorf("convert data to []interface{} failed. which is [%v], data: [%v]", reflect.TypeOf(d), d)
		}
		rv.Set(reflect.MakeSlice(rt, 0, rv.Cap()))
		for i, di := range dv {
			nv := reflect.New(rt.Elem())
			err := InterfaceToStruct(di, nv.Interface())
			if err != nil {
				return fmt.Errorf("idx: [%v], err: [%v]", i, err)
			}
			rv.Set(reflect.Append(rv, nv.Elem()))
		}
	case reflect.Ptr:
		nv := reflect.New(rt.Elem())
		err := InterfaceToStruct(d, nv.Interface())
		if err != nil {
			return err
		}
		rv.Set(nv)
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
				v, err := hstr.ToIP(d.(string))
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
			if rt.Kind() == reflect.Int {
				rv.SetInt(cast.ToInt64(v))
			} else {
				return fmt.Errorf("unsupport type [%v]", rt)
			}
		}
	}

	return nil
}
