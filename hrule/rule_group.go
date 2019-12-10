package hrule

import (
	"fmt"
	"reflect"
)

func Evaluate(v interface{}) (bool, error) {
	rg, err := Compile(v)
	if err != nil {
		return false, err
	}
	return rg.Evaluate(v), nil
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

	rg := &RuleGroup{
		t:       t,
		condMap: map[string]*Cond{},
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("hrule")
		if tag == "" {
			continue
		}
		cond, err := NewCond(field.Tag.Get("hrule"), field.Type)
		if err != nil {
			return nil, fmt.Errorf("new cond failed. field: [%v], tag: [%v], err: [%v]", field.Name, tag, err)
		}
		rg.condMap[field.Name] = cond
	}

	return rg, nil
}

func (g *RuleGroup) Evaluate(v interface{}) bool {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if g.t != val.Type() {
		return false
	}
	t := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		cond, ok := g.condMap[t.Field(i).Name]
		if !ok {
			continue
		}
		if !cond.Evaluate(field.Interface()) {
			return false
		}
	}

	return true
}
