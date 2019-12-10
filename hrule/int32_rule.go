package hrule

import (
	"fmt"
	"strconv"
	"strings"
)

func NewInt32Rule(expr string) (Rule, error) {
	fun, params := ruleExpr.Tokenizer(expr)

	generator, ok := Int32RuleGenerator[fun]
	if !ok {
		return nil, fmt.Errorf("no such generator: [%v]", fun)
	}
	return generator(params)
}

func RegisterInt32RuleGenerator(fun string, generator RuleGenerator) {
	Int32RuleGenerator[fun] = generator
}

var Int32RuleGenerator = map[string]RuleGenerator{
	">=": func(params string) (Rule, error) {
		i, err := strconv.ParseInt(params, 10, 32)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(int32) >= int32(i)
		}, nil
	},
	"<=": func(params string) (Rule, error) {
		i, err := strconv.ParseInt(params, 10, 32)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(int32) <= int32(i)
		}, nil
	},
	">": func(params string) (Rule, error) {
		i, err := strconv.ParseInt(params, 10, 32)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(int32) > int32(i)
		}, nil
	},
	"<": func(params string) (Rule, error) {
		i, err := strconv.ParseInt(params, 10, 32)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(int32) < int32(i)
		}, nil
	},
	"mod": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseInt(vals[0], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseInt(vals[1], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(int32)%int32(x) == int32(y)
		}, nil
	},
	"in": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		iset := map[int32]bool{}
		for _, val := range vals {
			i, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				return nil, err
			}
			iset[int32(i)] = true
		}
		return func(v interface{}) bool {
			return iset[v.(int32)]
		}, nil
	},
	"range": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseInt(vals[0], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseInt(vals[1], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(int32) >= int32(x) && v.(int32) <= int32(y)
		}, nil
	},
}
