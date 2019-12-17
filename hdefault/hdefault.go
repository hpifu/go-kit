package hdefault

import (
	"fmt"
	"github.com/hpifu/go-kit/hstr"
	"reflect"
)

func SetDefault(v interface{}) error {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("invalid value type %v", reflect.TypeOf(v))
	}

	if v == nil {
		return fmt.Errorf("params should not be nil")
	}

	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		value := rt.Field(i).Tag.Get("hdef")
		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				nv := reflect.New(rt.Field(i).Type.Elem())
				field.Set(nv)
			}
			if err := SetDefault(field.Interface()); err != nil {
				return err
			}
		} else if field.Kind() == reflect.Struct {
			if err := SetDefault(field.Addr().Interface()); err != nil {
				return err
			}
		} else {
			if value == "" || value == "-" {
				continue
			}
			err := hstr.SetValue(field, value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
