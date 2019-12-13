package henv

import (
	"errors"
	"github.com/hpifu/go-kit/hstring"
	"net"
	"time"
)

var NoSuchKeyErr = errors.New("no such key")

func (h HEnv) GetString(key string) (string, error) {
	v, ok := h.kvs[key]
	if !ok {
		return "", NoSuchKeyErr
	}
	return v, nil
}

func (h HEnv) GetBool(key string) (bool, error) {
	v, ok := h.kvs[key]
	if !ok {
		return false, NoSuchKeyErr
	}
	return hstring.ToBool(v)
}

func (h HEnv) GetInt(key string) (int, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstring.ToInt(v)
}

func (h HEnv) GetUint(key string) (uint, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstring.ToUint(v)
}

func (h HEnv) GetInt64(key string) (int64, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstring.ToInt64(v)
}

func (h HEnv) GetInt32(key string) (int32, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstring.ToInt32(v)
}

func (h HEnv) GetInt16(key string) (int16, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstring.ToInt16(v)
}

func (h HEnv) GetInt8(key string) (int8, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstring.ToInt8(v)
}

func (h HEnv) GetUint64(key string) (uint64, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstring.ToUint64(v)
}

func (h HEnv) GetUint32(key string) (uint32, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstring.ToUint32(v)
}

func (h HEnv) GetUint16(key string) (uint16, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstring.ToUint16(v)
}

func (h HEnv) GetUint8(key string) (uint8, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstring.ToUint8(v)
}

func (h HEnv) GetFloat64(key string) (float64, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstring.ToFloat64(v)
}

func (h HEnv) GetFloat32(key string) (float32, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstring.ToFloat32(v)
}

func (h HEnv) GetDuration(key string) (time.Duration, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstring.ToDuration(v)
}

func (h HEnv) GetTime(key string) (time.Time, error) {
	v, ok := h.kvs[key]
	if !ok {
		return time.Time{}, NoSuchKeyErr
	}
	return hstring.ToTime(v)
}

func (h HEnv) GetIP(key string) (net.IP, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToIP(v)
}

func (h HEnv) GetStringSlice(key string) ([]string, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToStringSlice(v)
}

func (h HEnv) GetBoolSlice(key string) ([]bool, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToBoolSlice(v)
}

func (h HEnv) GetIntSlice(key string) ([]int, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToIntSlice(v)
}

func (h HEnv) GetUintSlice(key string) ([]uint, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToUintSlice(v)
}

func (h HEnv) GetInt64Slice(key string) ([]int64, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToInt64Slice(v)
}

func (h HEnv) GetInt32Slice(key string) ([]int32, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToInt32Slice(v)
}

func (h HEnv) GetInt16Slice(key string) ([]int16, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToInt16Slice(v)
}

func (h HEnv) GetInt8Slice(key string) ([]int8, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToInt8Slice(v)
}

func (h HEnv) GetUint64Slice(key string) ([]uint64, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToUint64Slice(v)
}

func (h HEnv) GetUint32Slice(key string) ([]uint32, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToUint32Slice(v)
}

func (h HEnv) GetUint16Slice(key string) ([]uint16, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToUint16Slice(v)
}

func (h HEnv) GetUint8Slice(key string) ([]uint8, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToUint8Slice(v)
}

func (h HEnv) GetFloat64Slice(key string) ([]float64, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToFloat64Slice(v)
}

func (h HEnv) GetFloat32Slice(key string) ([]float32, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToFloat32Slice(v)
}

func (h HEnv) GetDurationSlice(key string) ([]time.Duration, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToDurationSlice(v)
}

func (h HEnv) GetTimeSlice(key string) ([]time.Time, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToTimeSlice(v)
}

func (h HEnv) GetIPSlice(key string) ([]net.IP, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstring.ToIPSlice(v)
}
