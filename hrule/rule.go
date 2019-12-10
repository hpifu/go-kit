package hrule

import (
	"fmt"
	"reflect"
	"time"
)

var ruleExpr = RuleExpr{}

func NewRule(expr string, t reflect.Type) (Rule, error) {
	switch t {
	case reflect.TypeOf(int(0)):
		return NewIntRule(expr)
	case reflect.TypeOf(int64(0)):
		return NewInt64Rule(expr)
	case reflect.TypeOf(int32(0)):
		return NewInt32Rule(expr)
	case reflect.TypeOf(int16(0)):
		return NewInt16Rule(expr)
	case reflect.TypeOf(int8(0)):
		return NewInt8Rule(expr)
	case reflect.TypeOf(uint64(0)):
		return NewUint64Rule(expr)
	case reflect.TypeOf(uint32(0)):
		return NewUint32Rule(expr)
	case reflect.TypeOf(uint16(0)):
		return NewUint16Rule(expr)
	case reflect.TypeOf(uint8(0)):
		return NewUint8Rule(expr)
	case reflect.TypeOf(time.Duration(0)):
		return NewDurationRule(expr)
	case reflect.TypeOf(time.Time{}):
		return NewTimeRule(expr)
	case reflect.TypeOf(float64(0)):
		return NewFloat64Rule(expr)
	case reflect.TypeOf(float32(0)):
		return NewFloat32Rule(expr)
	case reflect.TypeOf(""):
		return NewStringRule(expr)
	}

	return nil, fmt.Errorf("unsupport type [%v]", t)
}

type Rule func(interface{}) bool
type RuleGenerator func(string) (Rule, error)
