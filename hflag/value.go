package hflag

import (
	"bytes"
	"fmt"
	"github.com/hpifu/go-kit/hstring"
	"net"
	"strconv"
	"strings"
	"time"
)

type Value interface {
	Set(string) error
	String() string
}

type boolValue bool
type intValue int
type uintValue uint
type int64Value int64
type int32Value int32
type int16Value int16
type int8Value int8
type uint64Value uint64
type uint32Value uint32
type uint16Value uint16
type uint8Value uint8
type float64Value float64
type float32Value float32
type durationValue time.Duration
type timeValue time.Time
type ipValue net.IP
type stringValue string
type boolSliceValue []bool
type intSliceValue []int
type uintSliceValue []uint
type int64SliceValue []int64
type int32SliceValue []int32
type int16SliceValue []int16
type int8SliceValue []int8
type uint64SliceValue []uint64
type uint32SliceValue []uint32
type uint16SliceValue []uint16
type uint8SliceValue []uint8
type float64SliceValue []float64
type float32SliceValue []float32
type durationSliceValue []time.Duration
type timeSliceValue []time.Time
type ipSliceValue []net.IP
type stringSliceValue []string

func (v *boolValue) Set(str string) error {
	val, err := hstring.ToBool(str)
	if err != nil {
		return err
	}
	*v = boolValue(val)
	return nil
}

func (v *intValue) Set(str string) error {
	val, err := hstring.ToInt(str)
	if err != nil {
		return err
	}
	*v = intValue(val)
	return nil
}

func (v *uintValue) Set(str string) error {
	val, err := hstring.ToUint(str)
	if err != nil {
		return err
	}
	*v = uintValue(val)
	return nil
}

func (v *int64Value) Set(str string) error {
	val, err := hstring.ToInt64(str)
	if err != nil {
		return err
	}
	*v = int64Value(val)
	return nil
}

func (v *int32Value) Set(str string) error {
	val, err := hstring.ToInt32(str)
	if err != nil {
		return err
	}
	*v = int32Value(val)
	return nil
}

func (v *int16Value) Set(str string) error {
	val, err := hstring.ToInt16(str)
	if err != nil {
		return err
	}
	*v = int16Value(val)
	return nil
}

func (v *int8Value) Set(str string) error {
	val, err := hstring.ToInt8(str)
	if err != nil {
		return err
	}
	*v = int8Value(val)
	return nil
}

func (v *uint64Value) Set(str string) error {
	val, err := hstring.ToUint64(str)
	if err != nil {
		return err
	}
	*v = uint64Value(val)
	return nil
}

func (v *uint32Value) Set(str string) error {
	val, err := hstring.ToUint32(str)
	if err != nil {
		return err
	}
	*v = uint32Value(val)
	return nil
}

func (v *uint16Value) Set(str string) error {
	val, err := hstring.ToUint16(str)
	if err != nil {
		return err
	}
	*v = uint16Value(val)
	return nil
}

func (v *uint8Value) Set(str string) error {
	val, err := hstring.ToUint8(str)
	if err != nil {
		return err
	}
	*v = uint8Value(val)
	return nil
}

func (v *float64Value) Set(str string) error {
	val, err := hstring.ToFloat64(str)
	if err != nil {
		return err
	}
	*v = float64Value(val)
	return nil
}

func (v *float32Value) Set(str string) error {
	val, err := hstring.ToFloat32(str)
	if err != nil {
		return err
	}
	*v = float32Value(val)
	return nil
}

func (v *durationValue) Set(str string) error {
	val, err := hstring.ToDuration(str)
	if err != nil {
		return err
	}
	*v = durationValue(val)
	return nil
}

func (v *timeValue) Set(str string) error {
	val, err := hstring.ToTime(str)
	if err != nil {
		return err
	}
	*v = timeValue(val)
	return nil
}

func (v *ipValue) Set(str string) error {
	val, err := hstring.ToIP(str)
	if err != nil {
		return err
	}
	*v = ipValue(val)
	return nil
}

func (v *stringValue) Set(str string) error {
	*v = stringValue(str)
	return nil
}

func (v *boolSliceValue) Set(str string) error {
	val, err := hstring.ToBoolSlice(str)
	if err != nil {
		return err
	}
	*v = boolSliceValue(val)
	return nil
}

func (v *intSliceValue) Set(str string) error {
	val, err := hstring.ToIntSlice(str)
	if err != nil {
		return err
	}
	*v = intSliceValue(val)
	return nil
}

func (v *uintSliceValue) Set(str string) error {
	val, err := hstring.ToUintSlice(str)
	if err != nil {
		return err
	}
	*v = uintSliceValue(val)
	return nil
}

func (v *int64SliceValue) Set(str string) error {
	val, err := hstring.ToInt64Slice(str)
	if err != nil {
		return err
	}
	*v = int64SliceValue(val)
	return nil
}

func (v *int32SliceValue) Set(str string) error {
	val, err := hstring.ToInt32Slice(str)
	if err != nil {
		return err
	}
	*v = int32SliceValue(val)
	return nil
}

func (v *int16SliceValue) Set(str string) error {
	val, err := hstring.ToInt16Slice(str)
	if err != nil {
		return err
	}
	*v = int16SliceValue(val)
	return nil
}

func (v *int8SliceValue) Set(str string) error {
	val, err := hstring.ToInt8Slice(str)
	if err != nil {
		return err
	}
	*v = int8SliceValue(val)
	return nil
}

func (v *uint64SliceValue) Set(str string) error {
	val, err := hstring.ToUint64Slice(str)
	if err != nil {
		return err
	}
	*v = uint64SliceValue(val)
	return nil
}

func (v *uint32SliceValue) Set(str string) error {
	val, err := hstring.ToUint32Slice(str)
	if err != nil {
		return err
	}
	*v = uint32SliceValue(val)
	return nil
}

func (v *uint16SliceValue) Set(str string) error {
	val, err := hstring.ToUint16Slice(str)
	if err != nil {
		return err
	}
	*v = uint16SliceValue(val)
	return nil
}

func (v *uint8SliceValue) Set(str string) error {
	val, err := hstring.ToUint8Slice(str)
	if err != nil {
		return err
	}
	*v = uint8SliceValue(val)
	return nil
}

func (v *float64SliceValue) Set(str string) error {
	val, err := hstring.ToFloat64Slice(str)
	if err != nil {
		return err
	}
	*v = float64SliceValue(val)
	return nil
}

func (v *float32SliceValue) Set(str string) error {
	val, err := hstring.ToFloat32Slice(str)
	if err != nil {
		return err
	}
	*v = float32SliceValue(val)
	return nil
}

func (v *durationSliceValue) Set(str string) error {
	val, err := hstring.ToDurationSlice(str)
	if err != nil {
		return err
	}
	*v = durationSliceValue(val)
	return nil
}

func (v *timeSliceValue) Set(str string) error {
	val, err := hstring.ToTimeSlice(str)
	if err != nil {
		return err
	}
	*v = timeSliceValue(val)
	return nil
}

func (v *ipSliceValue) Set(str string) error {
	val, err := hstring.ToIPSlice(str)
	if err != nil {
		return err
	}
	*v = ipSliceValue(val)
	return nil
}

func (v *stringSliceValue) Set(str string) error {
	val, err := hstring.ToStringSlice(str)
	if err != nil {
		return err
	}
	*v = stringSliceValue(val)
	return nil
}

func NewValueType(typeStr string) Value {
	switch typeStr {
	case "bool":
		return new(boolValue)
	case "int":
		return new(intValue)
	case "float", "float64":
		return new(float64Value)
	case "string":
		return new(stringValue)
	case "duration":
		return new(durationValue)
	case "int64":
		return new(int64Value)
	case "uint64":
		return new(uint64Value)
	case "[]int":
		return new(intSliceValue)
	case "[]string":
		return new(stringSliceValue)
	case "time":
		return new(timeValue)
	case "ip":
		return new(ipValue)
	}

	return nil
}

func (v intValue) String() string {
	return strconv.Itoa(int(v))
}

func (v uintValue) String() string {
	return strconv.FormatUint(uint64(uint(v)), 10)
}

func (v int64Value) String() string {
	return strconv.FormatInt(int64(v), 10)
}

func (v uint64Value) String() string {
	return strconv.FormatUint(uint64(uint(v)), 10)
}

func (v float64Value) String() string {
	return fmt.Sprintf("%f", float64(v))
}

func (v stringValue) String() string {
	return string(v)
}

func (v durationValue) String() string {
	return time.Duration(v).String()
}

func (v boolValue) String() string {
	return fmt.Sprintf("%v", bool(v))
}

func (v intSliceValue) String() string {
	var buffer bytes.Buffer

	vi := []int(v)
	for idx, i := range vi {
		buffer.WriteString(strconv.Itoa(i))
		if idx != len(vi)-1 {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func (v stringSliceValue) String() string {
	return strings.Join([]string(v), ",")
}

func (v timeValue) String() string {
	return time.Time(v).Format(time.RFC3339)
}

func (v ipValue) String() string {
	return net.IP(v).String()
}
