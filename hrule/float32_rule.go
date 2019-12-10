package hrule

import (
	"fmt"
	"strconv"
	"strings"
)

func NewFloat32Rule(expr string) (Rule, error) {
	fun, params := ruleExpr.Tokenizer(expr)

	generator, ok := Float32RuleGenerator[fun]
	if !ok {
		return nil, fmt.Errorf("no such generator: [%v]", fun)
	}
	return generator(params)
}

func RegisterFloat32RuleGenerator(fun string, generator RuleGenerator) {
	Float32RuleGenerator[fun] = generator
}

var Float32RuleGenerator = map[string]RuleGenerator{
	">=": func(params string) (Rule, error) {
		i, err := strconv.ParseFloat(params, 32)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(float32) >= float32(i)
		}, nil
	},
	"<=": func(params string) (Rule, error) {
		i, err := strconv.ParseFloat(params, 32)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(float32) <= float32(i)
		}, nil
	},
	">": func(params string) (Rule, error) {
		i, err := strconv.ParseFloat(params, 32)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(float32) > float32(i)
		}, nil
	},
	"<": func(params string) (Rule, error) {
		i, err := strconv.ParseFloat(params, 32)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(float32) < float32(i)
		}, nil
	},
	"in": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		iset := map[float32]bool{}
		for _, val := range vals {
			i, err := strconv.ParseFloat(val, 32)
			if err != nil {
				return nil, err
			}
			iset[float32(i)] = true
		}
		return func(v interface{}) bool {
			return iset[v.(float32)]
		}, nil
	},
	"range": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseFloat(vals[0], 32)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseFloat(vals[1], 32)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(float32) >= float32(x) && v.(float32) <= float32(y)
		}, nil
	},
}
