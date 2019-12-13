package hconf

import (
	"fmt"
	"github.com/hpifu/go-kit/hstring"
	"github.com/spf13/cast"
	"net"
	"reflect"
	"time"
)

func (h HConf) GetDefaultBool(key string, defaultValue ...bool) bool {
	v, err := h.GetBool(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return false
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultInt(key string, defaultValue ...int) int {
	v, err := h.GetInt(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultUint(key string, defaultValue ...uint) uint {
	v, err := h.GetUint(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultInt64(key string, defaultValue ...int64) int64 {
	v, err := h.GetInt64(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultInt32(key string, defaultValue ...int32) int32 {
	v, err := h.GetInt32(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultInt16(key string, defaultValue ...int16) int16 {
	v, err := h.GetInt16(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultInt8(key string, defaultValue ...int8) int8 {
	v, err := h.GetInt8(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultUint64(key string, defaultValue ...uint64) uint64 {
	v, err := h.GetUint64(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultUint32(key string, defaultValue ...uint32) uint32 {
	v, err := h.GetUint32(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultUint16(key string, defaultValue ...uint16) uint16 {
	v, err := h.GetUint16(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultUint8(key string, defaultValue ...uint8) uint8 {
	v, err := h.GetUint8(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultFloat64(key string, defaultValue ...float64) float64 {
	v, err := h.GetFloat64(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0.0
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultFloat32(key string, defaultValue ...float32) float32 {
	v, err := h.GetFloat32(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0.0
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultString(key string, defaultValue ...string) string {
	v, err := h.GetString(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return ""
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultDuration(key string, defaultValue ...time.Duration) time.Duration {
	v, err := h.GetDuration(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return 0
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultTime(key string, defaultValue ...time.Time) time.Time {
	v, err := h.GetTime(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return time.Unix(0, 0)
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetDefaultIP(key string, defaultValue ...net.IP) net.IP {
	v, err := h.GetIP(key)
	if err != nil {
		if len(defaultValue) == 0 {
			return nil
		}
		return defaultValue[0]
	}
	return v
}

func (h HConf) GetBool(key string) (bool, error) {
	v, err := h.Get(key)
	if err != nil {
		return false, err
	}
	return cast.ToBoolE(v)
}

func (h HConf) GetInt(key string) (int, error) {
	v, err := h.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToIntE(v)
}

func (h HConf) GetUint(key string) (uint, error) {
	v, err := h.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToUintE(v)
}

func (h HConf) GetInt64(key string) (int64, error) {
	v, err := h.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToInt64E(v)
}

func (h HConf) GetInt32(key string) (int32, error) {
	v, err := h.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToInt32E(v)
}

func (h HConf) GetInt16(key string) (int16, error) {
	v, err := h.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToInt16E(v)
}

func (h HConf) GetInt8(key string) (int8, error) {
	v, err := h.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToInt8E(v)
}

func (h HConf) GetUint64(key string) (uint64, error) {
	v, err := h.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToUint64E(v)
}

func (h HConf) GetUint32(key string) (uint32, error) {
	v, err := h.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToUint32E(v)
}

func (h HConf) GetUint16(key string) (uint16, error) {
	v, err := h.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToUint16E(v)
}

func (h HConf) GetUint8(key string) (uint8, error) {
	v, err := h.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToUint8E(v)
}

func (h HConf) GetFloat64(key string) (float64, error) {
	v, err := h.Get(key)
	if err != nil {
		return 0.0, err
	}
	return cast.ToFloat64E(v)
}

func (h HConf) GetFloat32(key string) (float32, error) {
	v, err := h.Get(key)
	if err != nil {
		return 0.0, err
	}
	return cast.ToFloat32E(v)
}

func (h HConf) GetString(key string) (string, error) {
	v, err := h.Get(key)
	if err != nil {
		return "", err
	}
	return cast.ToStringE(v)
}

func (h HConf) GetDuration(key string) (time.Duration, error) {
	v, err := h.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToDurationE(v)
}

func (h HConf) GetTime(key string) (time.Time, error) {
	v, err := h.Get(key)
	if err != nil {
		return time.Time{}, err
	}
	return cast.ToTimeE(v)
}

func (h HConf) GetIP(key string) (net.IP, error) {
	v, err := h.Get(key)
	if err != nil {
		return nil, err
	}
	switch v.(type) {
	case string:
		return hstring.ToIP(v.(string))
	case net.IP:
		return v.(net.IP), nil
	default:
		return nil, fmt.Errorf("convert type [%v] to ip failed", reflect.TypeOf(v))
	}
}

func (h HConf) GetBoolSlice(key string) ([]bool, error) {
	v, err := h.Get(key)
	if err != nil {
		return nil, err
	}
	switch v.(type) {
	case string:
		return hstring.ToBoolSlice(v.(string))
	case []bool:
		return v.([]bool), nil
	case []interface{}:
		var res []bool
		for _, val := range v.([]interface{}) {
			d, err := cast.ToBoolE(val)
			if err != nil {
				return nil, err
			}
			res = append(res, d)
		}
		return res, nil
	default:
		return nil, fmt.Errorf("convert to bool slice failed")
	}
}

func (h HConf) GetIntSlice(key string) ([]int, error) {
	v, err := h.Get(key)
	if err != nil {
		return nil, err
	}
	switch v.(type) {
	case string:
		return hstring.ToIntSlice(v.(string))
	case []int:
		return v.([]int), nil
	case []interface{}:
		var res []int
		for _, val := range v.([]interface{}) {
			d, err := cast.ToIntE(val)
			if err != nil {
				return nil, err
			}
			res = append(res, d)
		}
		return res, nil
	default:
		return nil, fmt.Errorf("convert to bool slice failed")
	}
}
