package henv

import (
	"errors"
	"github.com/hpifu/go-kit/hstr"
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
	return hstr.ToBool(v)
}

func (h HEnv) GetInt(key string) (int, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstr.ToInt(v)
}

func (h HEnv) GetUint(key string) (uint, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstr.ToUint(v)
}

func (h HEnv) GetInt64(key string) (int64, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstr.ToInt64(v)
}

func (h HEnv) GetInt32(key string) (int32, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstr.ToInt32(v)
}

func (h HEnv) GetInt16(key string) (int16, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstr.ToInt16(v)
}

func (h HEnv) GetInt8(key string) (int8, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstr.ToInt8(v)
}

func (h HEnv) GetUint64(key string) (uint64, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstr.ToUint64(v)
}

func (h HEnv) GetUint32(key string) (uint32, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstr.ToUint32(v)
}

func (h HEnv) GetUint16(key string) (uint16, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstr.ToUint16(v)
}

func (h HEnv) GetUint8(key string) (uint8, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstr.ToUint8(v)
}

func (h HEnv) GetFloat64(key string) (float64, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstr.ToFloat64(v)
}

func (h HEnv) GetFloat32(key string) (float32, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstr.ToFloat32(v)
}

func (h HEnv) GetDuration(key string) (time.Duration, error) {
	v, ok := h.kvs[key]
	if !ok {
		return 0, NoSuchKeyErr
	}
	return hstr.ToDuration(v)
}

func (h HEnv) GetTime(key string) (time.Time, error) {
	v, ok := h.kvs[key]
	if !ok {
		return time.Time{}, NoSuchKeyErr
	}
	return hstr.ToTime(v)
}

func (h HEnv) GetIP(key string) (net.IP, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToIP(v)
}

func (h HEnv) GetStringSlice(key string) ([]string, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToStringSlice(v)
}

func (h HEnv) GetBoolSlice(key string) ([]bool, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToBoolSlice(v)
}

func (h HEnv) GetIntSlice(key string) ([]int, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToIntSlice(v)
}

func (h HEnv) GetUintSlice(key string) ([]uint, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToUintSlice(v)
}

func (h HEnv) GetInt64Slice(key string) ([]int64, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToInt64Slice(v)
}

func (h HEnv) GetInt32Slice(key string) ([]int32, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToInt32Slice(v)
}

func (h HEnv) GetInt16Slice(key string) ([]int16, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToInt16Slice(v)
}

func (h HEnv) GetInt8Slice(key string) ([]int8, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToInt8Slice(v)
}

func (h HEnv) GetUint64Slice(key string) ([]uint64, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToUint64Slice(v)
}

func (h HEnv) GetUint32Slice(key string) ([]uint32, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToUint32Slice(v)
}

func (h HEnv) GetUint16Slice(key string) ([]uint16, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToUint16Slice(v)
}

func (h HEnv) GetUint8Slice(key string) ([]uint8, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToUint8Slice(v)
}

func (h HEnv) GetFloat64Slice(key string) ([]float64, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToFloat64Slice(v)
}

func (h HEnv) GetFloat32Slice(key string) ([]float32, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToFloat32Slice(v)
}

func (h HEnv) GetDurationSlice(key string) ([]time.Duration, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToDurationSlice(v)
}

func (h HEnv) GetTimeSlice(key string) ([]time.Time, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToTimeSlice(v)
}

func (h HEnv) GetIPSlice(key string) ([]net.IP, error) {
	v, ok := h.kvs[key]
	if !ok {
		return nil, NoSuchKeyErr
	}
	return hstr.ToIPSlice(v)
}
