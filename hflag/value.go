package hflag

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type Value interface {
	Set(string) error
	String() string
}

type intValue int
type uintValue uint
type int64Value int64
type uint64Value uint64
type floatValue float64
type durationValue time.Duration
type stringValue string
type boolValue bool
type intSliceValue []int
type stringSliceValue []string
type timeValue time.Time
type ipValue net.IP

func NewValueType(typeStr string) Value {
	switch typeStr {
	case "bool":
		return new(boolValue)
	case "int":
		return new(intValue)
	case "float", "float64":
		return new(floatValue)
	case "string":
		return new(stringValue)
	case "duration":
		return new(durationValue)
	case "uint":
		return new(uintValue)
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

func (v *intValue) Set(str string) error {
	i, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	*v = intValue(i)
	return nil
}

func (v *uintValue) Set(str string) error {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}
	*v = uintValue(i)
	return nil
}

func (v *uint64Value) Set(str string) error {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}
	*v = uint64Value(i)
	return nil
}

func (v *int64Value) Set(str string) error {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	*v = int64Value(i)
	return nil
}

func (v *floatValue) Set(str string) error {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}

	*v = floatValue(f)
	return nil
}

func (v *durationValue) Set(str string) error {
	d, err := time.ParseDuration(str)
	if err != nil {
		return err
	}

	*v = durationValue(d)
	return nil
}

func (v *stringValue) Set(str string) error {
	*v = stringValue(str)
	return nil
}

func (v *boolValue) Set(str string) error {
	b, err := strconv.ParseBool(str)
	if err != nil {
		return err
	}

	*v = boolValue(b)
	return nil
}

func (v *intSliceValue) Set(strs string) error {
	var ia []int
	for _, str := range strings.Split(strs, ",") {
		i, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		ia = append(ia, i)
	}
	*v = intSliceValue(ia)
	return nil
}

func (v *stringSliceValue) Set(strs string) error {
	*v = stringSliceValue(strings.Split(strs, ","))
	return nil
}

func (v *timeValue) Set(str string) error {
	var t time.Time
	var err error
	if str == "now" {
		t = time.Now()
	} else if len(str) == 10 {
		t, err = time.Parse("2006-01-02", str)
	} else if len(str) == 19 {
		t, err = time.Parse("2006-01-02T15:04:05", str)
	} else {
		t, err = time.Parse(time.RFC3339, str)
	}
	if err != nil {
		return err
	}
	*v = timeValue(t)
	return nil
}

func (v *ipValue) Set(str string) error {
	i := net.ParseIP(str)
	if i == nil {
		return fmt.Errorf("parse [%v] to ip failed", str)
	}
	*v = ipValue(i)
	return nil
}

func (v intValue) String() string {
	return strconv.Itoa(int(v))
}

func (v int64Value) String() string {
	return strconv.FormatInt(int64(v), 10)
}

func (v uintValue) String() string {
	return strconv.FormatUint(uint64(uint(v)), 10)
}

func (v uint64Value) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

func (v floatValue) String() string {
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
