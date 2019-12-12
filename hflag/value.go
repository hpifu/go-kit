package hflag

import (
	"github.com/hpifu/go-kit/hstring"
	"net"
	"time"
)

func NewValueType(typeStr string) Value {
	switch typeStr {
	case "bool":
		return new(boolValue)
	case "int":
		return new(intValue)
	case "uint":
		return new(uintValue)
	case "int64":
		return new(int64Value)
	case "int32":
		return new(int32Value)
	case "int16":
		return new(int16Value)
	case "int8":
		return new(int8Value)
	case "uint64":
		return new(uint64Value)
	case "uint32":
		return new(uint32Value)
	case "uint16":
		return new(uint16Value)
	case "uint8":
		return new(uint8Value)
	case "float64", "float":
		return new(float64Value)
	case "float32":
		return new(float32Value)
	case "duration":
		return new(durationValue)
	case "time":
		return new(timeValue)
	case "ip":
		return new(ipValue)
	case "string":
		return new(stringValue)
	case "[]bool":
		return new(boolSliceValue)
	case "[]int":
		return new(intSliceValue)
	case "[]uint":
		return new(uintSliceValue)
	case "[]int64":
		return new(int64SliceValue)
	case "[]int32":
		return new(int32SliceValue)
	case "[]int16":
		return new(int16SliceValue)
	case "[]int8":
		return new(int8SliceValue)
	case "[]uint64":
		return new(uint64SliceValue)
	case "[]uint32":
		return new(uint32SliceValue)
	case "[]uint16":
		return new(uint16SliceValue)
	case "[]uint8":
		return new(uint8SliceValue)
	case "[]float64":
		return new(float64SliceValue)
	case "[]float32":
		return new(float32SliceValue)
	case "[]duration":
		return new(durationSliceValue)
	case "[]time":
		return new(timeSliceValue)
	case "[]ip":
		return new(ipSliceValue)
	case "[]string":
		return new(stringSliceValue)
	default:
		return nil
	}
}

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

func (v boolValue) String() string {
	return hstring.BoolTo(bool(v))
}

func (v intValue) String() string {
	return hstring.IntTo(int(v))
}

func (v uintValue) String() string {
	return hstring.UintTo(uint(v))
}

func (v int64Value) String() string {
	return hstring.Int64To(int64(v))
}

func (v int32Value) String() string {
	return hstring.Int32To(int32(v))
}

func (v int16Value) String() string {
	return hstring.Int16To(int16(v))
}

func (v int8Value) String() string {
	return hstring.Int8To(int8(v))
}

func (v uint64Value) String() string {
	return hstring.Uint64To(uint64(v))
}

func (v uint32Value) String() string {
	return hstring.Uint32To(uint32(v))
}

func (v uint16Value) String() string {
	return hstring.Uint16To(uint16(v))
}

func (v uint8Value) String() string {
	return hstring.Uint8To(uint8(v))
}

func (v float64Value) String() string {
	return hstring.Float64To(float64(v))
}

func (v float32Value) String() string {
	return hstring.Float32To(float32(v))
}

func (v durationValue) String() string {
	return hstring.DurationTo(time.Duration(v))
}

func (v timeValue) String() string {
	return hstring.TimeTo(time.Time(v))
}

func (v ipValue) String() string {
	return hstring.IPTo(net.IP(v))
}

func (v stringValue) String() string {
	return string(v)
}

func (v boolSliceValue) String() string {
	return hstring.BoolSliceTo([]bool(v))
}

func (v intSliceValue) String() string {
	return hstring.IntSliceTo([]int(v))
}

func (v uintSliceValue) String() string {
	return hstring.UintSliceTo([]uint(v))
}

func (v int64SliceValue) String() string {
	return hstring.Int64SliceTo([]int64(v))
}

func (v int32SliceValue) String() string {
	return hstring.Int32SliceTo([]int32(v))
}

func (v int16SliceValue) String() string {
	return hstring.Int16SliceTo([]int16(v))
}

func (v int8SliceValue) String() string {
	return hstring.Int8SliceTo([]int8(v))
}

func (v uint64SliceValue) String() string {
	return hstring.Uint64SliceTo([]uint64(v))
}

func (v uint32SliceValue) String() string {
	return hstring.Uint32SliceTo([]uint32(v))
}

func (v uint16SliceValue) String() string {
	return hstring.Uint16SliceTo([]uint16(v))
}

func (v uint8SliceValue) String() string {
	return hstring.Uint8SliceTo([]uint8(v))
}

func (v float64SliceValue) String() string {
	return hstring.Float64SliceTo([]float64(v))
}

func (v float32SliceValue) String() string {
	return hstring.Float32SliceTo([]float32(v))
}

func (v durationSliceValue) String() string {
	return hstring.DurationSliceTo([]time.Duration(v))
}

func (v timeSliceValue) String() string {
	return hstring.TimeSliceTo([]time.Time(v))
}

func (v ipSliceValue) String() string {
	return hstring.IPSliceTo([]net.IP(v))
}

func (v stringSliceValue) String() string {
	return hstring.StringSliceTo([]string(v))
}
