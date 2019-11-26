package hflag

import (
	"fmt"
	"strconv"
	"time"
)

type Value interface {
	Set(string) error
	String() string
}

type intValue int
type floatValue float64
type durationValue time.Duration
type stringValue string
type boolValue bool

func NewValueType(typeStr string) Value {
	switch typeStr {
	case "bool":
		return new(boolValue)
	case "int":
		return new(intValue)
	case "float":
		return new(floatValue)
	case "string":
		return new(stringValue)
	case "duration":
		return new(durationValue)
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

func (v intValue) String() string {
	return strconv.Itoa(int(v))
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
