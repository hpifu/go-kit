package hrule

import (
	"fmt"
	"reflect"
)

func Evaluate(v interface{}) error {
	rg, err := Compile(v)
	if err != nil {
		return err
	}
	return rg.Evaluate(v)
}

type RuleGroup struct {
	t       reflect.Type
	condMap map[string]*Cond
}

func MustCompile(v interface{}) *RuleGroup {
	rg, err := Compile(v)
	if err != nil {
		panic(err)
	}
	return rg
}

func Compile(v interface{}) (*RuleGroup, error) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	condMap := map[string]*Cond{}
	if err := interfaceToRules(condMap, t, ""); err != nil {
		return nil, err
	}

	return &RuleGroup{
		t:       t,
		condMap: condMap,
	}, nil
}

func interfaceToRules(condMap map[string]*Cond, rt reflect.Type, prefix string) error {
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	for i := 0; i < rt.NumField(); i++ {
		t := rt.Field(i).Type
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		tag := rt.Field(i).Tag.Get("hrule")
		if tag == "-" {
			continue
		}

		key := rt.Field(i).Name
		if prefix != "" {
			key = prefix + "." + key
		}
		if t.Kind() == reflect.Struct {
			if err := interfaceToRules(condMap, t, key); err != nil {
				return err
			}
			continue
		}
		if tag == "" {
			continue
		}
		cond, err := NewCond(tag, t)
		if err != nil {
			return fmt.Errorf("new cond failed. field: [%v], tag: [%v], err: [%v]", key, tag, err)
		}
		condMap[key] = cond
	}
	return nil
}

func (g *RuleGroup) Evaluate(v interface{}) error {
	return g.evaluate(v, "")
}

func (g *RuleGroup) evaluate(v interface{}, prefix string) error {
	rt := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)
	if rt.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rt = rt.Elem()
		rv = rv.Elem()
	}

	if rt.Kind() == reflect.Struct {
		for i := 0; i < rt.NumField(); i++ {
			key := rt.Field(i).Name
			if prefix != "" {
				key = prefix + "." + key
			}
			if err := g.evaluate(rv.Field(i).Interface(), key); err != nil {
				return err
			}
		}
		return nil
	}

	cond, ok := g.condMap[prefix]
	if !ok {
		return nil
	}
	if !cond.Evaluate(rv.Interface()) {
		return fmt.Errorf("evaluate [%v] not pass", prefix)
	}
	return nil
}
