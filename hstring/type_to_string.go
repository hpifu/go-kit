package hstring

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func ToString(v interface{}) string {
	switch v.(type) {
	case bool:
		return BoolTo(v.(bool))
	default:
		return fmt.Sprintf("%v", v)
	}
}

func BoolTo(v bool) string {
	return fmt.Sprintf("%v", v)
}

func IntTo(v int) string {
	return strconv.Itoa(v)
}

func Int64To(v int64) string {
	return strconv.FormatInt(v, 10)
}

func Int32To(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

func Int16To(v int16) string {
	return strconv.FormatInt(int64(v), 10)
}

func Int8To(v int8) string {
	return strconv.FormatInt(int64(v), 10)
}

func UintTo(v uint) string {
	return strconv.FormatUint(uint64(v), 10)
}

func Uint64To(v uint64) string {
	return strconv.FormatUint(v, 10)
}

func Uint32To(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}

func Uint16To(v uint16) string {
	return strconv.FormatUint(uint64(v), 10)
}

func Uint8To(v uint8) string {
	return strconv.FormatUint(uint64(v), 10)
}

func Float64To(v float64) string {
	return fmt.Sprintf("%f", v)
}

func Float32To(v float32) string {
	return fmt.Sprintf("%f", v)
}

func DurationTo(v time.Duration) string {
	return v.String()
}

func TimeTo(v time.Time) string {
	return v.Format(time.RFC3339)
}

func IPTo(v net.IP) string {
	return v.String()
}
