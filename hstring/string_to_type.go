package hstring

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func ToBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

func ToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func ToUint(str string) (uint, error) {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}

func ToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func ToInt32(str string) (int32, error) {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

func ToInt16(str string) (int16, error) {
	i, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(i), nil
}

func ToInt8(str string) (int8, error) {
	i, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(i), nil
}

func ToUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

func ToUint32(str string) (uint32, error) {
	i, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(i), nil
}

func ToUint16(str string) (uint16, error) {
	i, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(i), nil
}

func ToUint8(str string) (uint8, error) {
	i, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(i), nil
}

func ToFloat64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func ToFloat32(str string) (float32, error) {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

func ToDuration(str string) (time.Duration, error) {
	return time.ParseDuration(str)
}

func ToTime(str string) (time.Time, error) {
	if str == "now" {
		return time.Now(), nil
	} else if len(str) == 10 {
		return time.Parse("2006-01-02", str)
	} else if len(str) == 19 {
		if str[10] == ' ' {
			return time.Parse("2006-01-02 15:04:05", str)
		}
		return time.Parse("2006-01-02T15:04:05", str)
	}

	return time.Parse(time.RFC3339, str)
}

func ToIP(str string) (net.IP, error) {
	ip := net.ParseIP(str)
	if ip == nil {
		return nil, fmt.Errorf("invalid ip [%v]", str)
	}
	return ip, nil
}
