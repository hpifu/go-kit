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
		return NewValue(false)
	case "int":
		return NewValue(0)
	case "float":
		return NewValue(0.0)
	case "string":
		return NewValue("")
	case "duration":
		return NewValue(time.Duration(0))
	}

	return nil
}

func NewValue(defaultValue interface{}) Value {
	switch defaultValue.(type) {
	case int:
		v := intValue(defaultValue.(int))
		return &v
	case float64:
		v := floatValue(defaultValue.(float64))
		return &v
	case bool:
		v := boolValue(defaultValue.(bool))
		return &v
	case string:
		v := stringValue(defaultValue.(string))
		return &v
	case time.Duration:
		v := durationValue(defaultValue.(time.Duration))
		return &v
	}

	return nil
}

func NewValueVar(p interface{}) Value {
	switch p.(type) {
	case *int:
		return (*intValue)(p.(*int))
	case *bool:
		return (*boolValue)(p.(*bool))
	case *float64:
		return (*floatValue)(p.(*float64))
	case *string:
		return (*stringValue)(p.(*string))
	case *time.Duration:
		return (*durationValue)(p.(*time.Duration))
	}

	return nil
}

func NewIntValue(defaultValue int, v *int) *intValue {
	*v = defaultValue
	return (*intValue)(v)
}

func NewFloatValue(defaultValue float64, v *float64) *floatValue {
	*v = defaultValue
	return (*floatValue)(v)
}

func NewDurationValue(defaultValue time.Duration, v *time.Duration) *durationValue {
	*v = defaultValue
	return (*durationValue)(v)
}

func NewStringValue(defaultValue string, v *string) *stringValue {
	*v = defaultValue
	return (*stringValue)(v)
}

func NewBoolValue(defaultValue bool, v *bool) *boolValue {
	*v = defaultValue
	return (*boolValue)(v)
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
