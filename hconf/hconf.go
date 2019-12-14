package hconf

import (
	"fmt"
	"github.com/hpifu/go-kit/hstring"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/yosuke-furukawa/json5/encoding/json5"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func NewHConfWithFile(filename string) (*HConf, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	var data interface{}
	if err = json5.NewDecoder(fp).Decode(&data); err != nil {
		return nil, err
	}

	return &HConf{
		filename:  filename,
		data:      data,
		separator: ".",
		log:       logrus.New(),
	}, nil
}

type HConf struct {
	filename  string
	data      interface{}
	separator string
	envPrefix string
	handlers  []OnChangeHandler
	log       *logrus.Logger
}

func (h *HConf) SetSeparator(separator string) {
	h.separator = separator
}

func (h *HConf) SetLogger(log *logrus.Logger) {
	h.log = log
}

const MapMod = 1
const ArrMod = 2

type KeyInfo struct {
	key string
	idx int
	mod int
}

func (h HConf) parseKey(keys string) ([]*KeyInfo, error) {
	var infos []*KeyInfo
	for _, key := range strings.Split(keys, h.separator) {
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

func (h HConf) Get(keys string) (interface{}, error) {
	data := h.data
	infos, err := h.parseKey(keys)
	if err != nil {
		return nil, err
	}

	for _, info := range infos {
		if info.mod == MapMod {
			val, ok := data.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("data is not a map. data: [%v]", data)
			}
			if data, ok = val[info.key]; !ok {
				return nil, fmt.Errorf("no such key")
			}
		} else {
			val, ok := data.([]interface{})
			if !ok {
				return nil, fmt.Errorf("data is not a array. data: [%v]", data)
			}
			if len(val) <= info.idx {
				return nil, fmt.Errorf("index out of bounds. index: [%v], data: [%v]", info.idx, data)
			}
			data = val[info.idx]
		}
	}

	return data, nil
}

func (h *HConf) Set(key string, val interface{}) error {
	data := h.data
	infos, err := h.parseKey(key)
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

func (h HConf) Unmarshal(v interface{}) error {
	return interfaceToStruct(h.data, v)
}

func interfaceToStruct(d interface{}, v interface{}) error {
	if d == nil {
		return nil
	}
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("invalid value type")
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

func (h HConf) Sub(key string) (*HConf, error) {
	v, err := h.Get(key)
	if err != nil {
		return nil, err
	}

	return &HConf{
		data:      v,
		separator: h.separator,
	}, nil
}

func (h *HConf) SetEnvPrefix(prefix string) {
	h.envPrefix = prefix
}

func (h *HConf) BindEnv(key string, env ...string) error {
	envkey := key
	if len(env) != 0 {
		envkey = env[0]
	} else {
		envkey = strings.ReplaceAll(envkey, h.separator, "_")
		envkey = strings.ReplaceAll(envkey, "[", "_")
		envkey = strings.ReplaceAll(envkey, "]", "")
		envkey = strings.ToUpper(envkey)
		if h.envPrefix != "" {
			envkey = h.envPrefix + "_" + envkey
		}
	}
	val := os.Getenv(envkey)

	if val == "" {
		return nil
	}

	v, err := h.Get(key)
	if err != nil {
		return h.Set(key, val)
	}

	switch v.(type) {
	case string:
		return h.Set(key, val)
	case float64:
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		return h.Set(key, f)
	case int:
		i, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		return h.Set(key, i)
	case float32:
		f, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return err
		}
		return h.Set(key, float32(f))
	}

	return fmt.Errorf("val type can not bind env, val: [%v]", v)
}
