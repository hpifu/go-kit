package hrule

import (
	"fmt"
	"strconv"
	"strings"
)

func NewFloat64Rule(expr string) (Rule, error) {
	fun, params := ruleExpr.Tokenizer(expr)

	generator, ok := Float64RuleGenerator[fun]
	if !ok {
		return nil, fmt.Errorf("no such generator: [%v]", fun)
	}
	return generator(params)
}

func RegisterFloat64RuleGenerator(fun string, generator RuleGenerator) {
	Float64RuleGenerator[fun] = generator
}

var Float64RuleGenerator = map[string]RuleGenerator{
	">=": func(params string) (Rule, error) {
		i, err := strconv.ParseFloat(params, 64)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(float64) >= i
		}, nil
	},
	"<=": func(params string) (Rule, error) {
		i, err := strconv.ParseFloat(params, 64)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(float64) <= i
		}, nil
	},
	">": func(params string) (Rule, error) {
		i, err := strconv.ParseFloat(params, 64)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(float64) > i
		}, nil
	},
	"<": func(params string) (Rule, error) {
		i, err := strconv.ParseFloat(params, 64)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(float64) < i
		}, nil
	},
	"in": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		iset := map[float64]bool{}
		for _, val := range vals {
			i, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
			iset[i] = true
		}
		return func(v interface{}) bool {
			return iset[v.(float64)]
		}, nil
	},
	"range": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseFloat(vals[0], 64)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseFloat(vals[1], 64)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(float64) >= x && v.(float64) <= y
		}, nil
	},
}
