package hconf

import (
	"fmt"
	"github.com/hpifu/go-kit/hstring"
	"github.com/spf13/cast"
	"net"
	"reflect"
	"time"
)

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

func (h HConf) GetDefaultFloat64(keys string, defaultValue ...float64) float64 {
	v, err := h.GetFloat64(keys)
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
