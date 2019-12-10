package hrule

import (
	"fmt"
	"strconv"
	"strings"
)

func NewInt16Rule(expr string) (Rule, error) {
	fun, params := ruleExpr.Tokenizer(expr)

	generator, ok := Int16RuleGenerator[fun]
	if !ok {
		return nil, fmt.Errorf("no such generator: [%v]", fun)
	}
	return generator(params)
}

func RegisterInt16RuleGenerator(fun string, generator RuleGenerator) {
	Int16RuleGenerator[fun] = generator
}

var Int16RuleGenerator = map[string]RuleGenerator{
	">=": func(params string) (Rule, error) {
		i, err := strconv.ParseInt(params, 10, 16)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(int16) >= int16(i)
		}, nil
	},
	"<=": func(params string) (Rule, error) {
		i, err := strconv.ParseInt(params, 10, 16)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(int16) <= int16(i)
		}, nil
	},
	">": func(params string) (Rule, error) {
		i, err := strconv.ParseInt(params, 10, 16)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(int16) > int16(i)
		}, nil
	},
	"<": func(params string) (Rule, error) {
		i, err := strconv.ParseInt(params, 10, 16)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(int16) < int16(i)
		}, nil
	},
	"mod": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseInt(vals[0], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseInt(vals[1], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(int16)%int16(x) == int16(y)
		}, nil
	},
	"in": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		iset := map[int16]bool{}
		for _, val := range vals {
			i, err := strconv.ParseInt(val, 10, 16)
			if err != nil {
				return nil, err
			}
			iset[int16(i)] = true
		}
		return func(v interface{}) bool {
			return iset[v.(int16)]
		}, nil
	},
	"range": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseInt(vals[0], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseInt(vals[1], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(int16) >= int16(x) && v.(int16) <= int16(y)
		}, nil
	},
}
