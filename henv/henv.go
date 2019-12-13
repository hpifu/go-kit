package henv

import (
	"fmt"
	"github.com/hpifu/go-kit/hstring"
	"net"
	"os"
	"reflect"
	"strings"
	"time"
)

func NewHEnv(prefix string) *HEnv {
	kvs := map[string]string{}
	for _, kv := range os.Environ() {
		idx := strings.Index(kv, "=")
		if idx < 0 {
			continue
		}
		key := kv[:idx]
		val := kv[idx+1:]

		if prefix != "" {
			if strings.HasPrefix(key, prefix+"_") {
				kvs[key[len(prefix)+1:]] = val
			}
		} else {
			kvs[key] = val
		}
	}

	return &HEnv{
		kvs:    kvs,
		prefix: prefix,
	}
}

type HEnv struct {
	prefix string
	kvs    map[string]string
}

func (h HEnv) Unmarshal(v interface{}) error {
	return mapToStruct(h.kvs, v, "")
}

func mapToStruct(kvs map[string]string, v interface{}, prefix string) error {
	if kvs == nil {
		return nil
	}
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("invalid value type [%v]", reflect.TypeOf(v))
	}

	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()
	switch rt.Kind() {
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			key := rt.Field(i).Tag.Get("henv")
			if key == "-" {
				continue
			}
			if key == "" {
				key = hstring.SnakeNameAllCaps(rt.Field(i).Name)
			}
			if prefix != "" {
				key = prefix + "_" + key
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
	default:
		val, ok := kvs[prefix]
		if !ok || val == "" {
			return nil
		}
		switch rv.Interface().(type) {
		case string:
			rv.Set(reflect.ValueOf(val))
		case bool:
			v, err := hstring.ToBool(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case int:
			v, err := hstring.ToInt(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case uint:
			v, err := hstring.ToUint(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case int64:
			v, err := hstring.ToInt64(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case int32:
			v, err := hstring.ToInt32(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case int16:
			v, err := hstring.ToInt16(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case int8:
			v, err := hstring.ToInt8(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case uint64:
			v, err := hstring.ToUint64(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case uint32:
			v, err := hstring.ToUint32(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case uint16:
			v, err := hstring.ToUint16(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case uint8:
			v, err := hstring.ToUint8(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case float64:
			v, err := hstring.ToFloat64(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case float32:
			v, err := hstring.ToFloat32(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case time.Duration:
			v, err := hstring.ToDuration(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case time.Time:
			v, err := hstring.ToTime(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case net.IP:
			v, err := hstring.ToIP(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []string:
			v, err := hstring.ToStringSlice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []bool:
			v, err := hstring.ToBoolSlice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []int:
			v, err := hstring.ToIntSlice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []uint:
			v, err := hstring.ToUintSlice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []int64:
			v, err := hstring.ToInt64Slice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []int32:
			v, err := hstring.ToInt32Slice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []int16:
			v, err := hstring.ToInt16Slice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []int8:
			v, err := hstring.ToInt8Slice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []uint64:
			v, err := hstring.ToUint64Slice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []uint32:
			v, err := hstring.ToUint32Slice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []uint16:
			v, err := hstring.ToUint16Slice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []uint8:
			v, err := hstring.ToUint8Slice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []float64:
			v, err := hstring.ToFloat64Slice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []float32:
			v, err := hstring.ToFloat32Slice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []time.Duration:
			v, err := hstring.ToDurationSlice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []time.Time:
			v, err := hstring.ToTimeSlice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		case []net.IP:
			v, err := hstring.ToIPSlice(val)
			if err != nil {
				return err
			}
			rv.Set(reflect.ValueOf(v))
		}
	}

	return nil
}
