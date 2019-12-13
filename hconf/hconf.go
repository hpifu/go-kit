package hconf

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/yosuke-furukawa/json5/encoding/json5"
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

	return &HConf{data: data, separator: "."}, nil
}

type HConf struct {
	data      interface{}
	separator string
	envPrefix string
}

func (h *HConf) SetSeparator(separator string) {
	h.separator = separator
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
	return interfaceToStructV2(h.data, v)
}

func interfaceToStructV2(d interface{}, v interface{}) error {
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
				key = rt.Field(i).Name
			}
			value := dv[key]
			if rt.Field(i).Type.Kind() == reflect.Ptr {
				if field.IsNil() {
					nv := reflect.New(field.Type().Elem())
					field.Set(nv)
				}
				if err := interfaceToStructV2(value, field.Interface()); err != nil {
					return err
				}
			} else {
				if err := interfaceToStructV2(value, field.Addr().Interface()); err != nil {
					return err
				}
			}
		}
	case reflect.Slice:
		dv, ok := d.([]interface{})
		if !ok {
			return fmt.Errorf("convert data to []interface{} failed. which is %v", reflect.TypeOf(d))
		}
		nv := reflect.New(rt.Elem())
		for _, di := range dv {
			err := interfaceToStructV2(di, nv.Interface())
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
		default:
			return fmt.Errorf("unsupport type %v", rt)
		}
	}

	return nil
}

func interfaceToStructV1(d interface{}, v interface{}) error {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("invalid value type")
	}

	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()

	if reflect.ValueOf(v).Elem().Kind() == reflect.Slice {
		dv, ok := d.([]interface{})
		fmt.Println(rt.Elem())
		if !ok {
			return fmt.Errorf("convert data to []interface{} failed. which is %v", reflect.TypeOf(d))
		}

		nv := reflect.New(rt.Elem())
		for _, di := range dv {
			err := interfaceToStructV1(di, nv.Interface())
			if err != nil {
				return err
			}
			rv.Set(reflect.Append(rv, nv.Elem()))
		}

		return nil
	}

	dv, ok := d.(map[string]interface{})
	if !ok {
		return fmt.Errorf("convert data to map[string]interface{} failed. which is %v", reflect.TypeOf(d))
	}

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		value := dv[rt.Field(i).Tag.Get("hconf")]
		switch rt.Field(i).Type.Kind() {
		case reflect.Int:
			i, err := cast.ToIntE(value)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(i))
		case reflect.Int64:
			if field.Type() == reflect.TypeOf(time.Duration(0)) {
				i, err := cast.ToStringE(value)
				if err != nil {
					return err
				}
				t, err := time.ParseDuration(i)
				if err != nil {
					return err
				}
				field.Set(reflect.ValueOf(t))
			} else {
				i, err := cast.ToInt64E(value)
				if err != nil {
					return err
				}
				field.Set(reflect.ValueOf(i))
			}
		case reflect.Uint64:
			i, err := cast.ToUint64E(value)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(i))
		case reflect.Float64:
			i, err := cast.ToFloat64E(value)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(i))
		case reflect.String:
			i, err := cast.ToStringE(value)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(i))
		case reflect.Struct:
			if err := interfaceToStructV1(value, field.Addr().Interface()); err != nil {
				return err
			}
		case reflect.Ptr:
			nv := reflect.New(field.Type().Elem())
			if err := interfaceToStructV1(value, nv.Interface()); err != nil {
				return err
			}
			field.Set(nv)
		}
	}

	return nil
}

func (h HConf) GetDefaultInt(keys string, defaultValue ...int) int {
	v, err := h.GetInt(keys)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}

	return v
}

func (h HConf) GetDefaultFloat(keys string, defaultValue ...float64) float64 {
	v, err := h.GetFloat(keys)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0.0
		}
		return defaultValue[0]
	}

	return v
}

func (h HConf) GetDefaultString(keys string, defaultValue ...string) string {
	v, err := h.GetString(keys)
	if err != nil {
		if len(defaultValue) == 0 {
			return ""
		}
		return defaultValue[0]
	}

	return v
}

func (h HConf) GetDefaultDuration(keys string, defaultValue ...time.Duration) time.Duration {
	v, err := h.GetDuration(keys)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}

	return v
}

func (h HConf) GetInt(keys string) (int, error) {
	v, err := h.Get(keys)
	if err != nil {
		return 0, err
	}

	switch v.(type) {
	case int:
		return v.(int), nil
	case float64:
		return int(v.(float64)), nil
	case float32:
		return int(v.(float32)), nil
	case string:
		i, err := strconv.Atoi(v.(string))
		if err != nil {
			return 0, err
		}
		return i, nil
	}

	return 0, fmt.Errorf("convert to int failed")
}

func (h HConf) GetFloat(keys string) (float64, error) {
	v, err := h.Get(keys)
	if err != nil {
		return 0.0, err
	}

	switch v.(type) {
	case int:
		return float64(v.(int)), nil
	case float64:
		return v.(float64), nil
	case float32:
		return float64(v.(float32)), nil
	case string:
		f, err := strconv.ParseFloat(v.(string), 64)
		if err != nil {
			return 0.0, err
		}
		return f, nil
	}

	return 0.0, fmt.Errorf("convert to float64 failed")
}

func (h HConf) GetString(keys string) (string, error) {
	v, err := h.Get(keys)
	if err != nil {
		return "", err
	}

	switch v.(type) {
	case int:
		return strconv.Itoa(v.(int)), nil
	case float64:
		return fmt.Sprintf("%f", v.(float64)), nil
	case float32:
		return fmt.Sprintf("%f", v.(float32)), nil
	case string:
		return v.(string), nil
	}

	return "", fmt.Errorf("convert to string failed")
}

func (h HConf) GetDuration(keys string) (time.Duration, error) {
	v, err := h.Get(keys)
	if err != nil {
		return 0, err
	}

	switch v.(type) {
	case string:
		return time.ParseDuration(v.(string))
	case int:
		return time.Duration(v.(int)) * time.Second, nil
	}

	return 0, fmt.Errorf("convert to duration failed")
}

func (h HConf) Sub(keys string) (*HConf, error) {
	i, err := h.Get(keys)
	if err != nil {
		return nil, err
	}

	return &HConf{
		data:      i,
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
