package hrule

import (
	"fmt"
	"strconv"
	"strings"
)

func NewUint8Rule(expr string) (Rule, error) {
	fun, params := ruleExpr.Tokenizer(expr)

	generator, ok := Uint8RuleGenerator[fun]
	if !ok {
		return nil, fmt.Errorf("no such generator: [%v]", fun)
	}
	return generator(params)
}

func RegisterUint8RuleGenerator(fun string, generator RuleGenerator) {
	Uint8RuleGenerator[fun] = generator
}

var Uint8RuleGenerator = map[string]RuleGenerator{
	">=": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 8)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint8) >= uint8(i)
		}, nil
	},
	"<=": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 8)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint8) <= uint8(i)
		}, nil
	},
	">": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 8)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint8) > uint8(i)
		}, nil
	},
	"<": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 8)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint8) < uint8(i)
		}, nil
	},
	"mod": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseUint(vals[0], 10, 8)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseUint(vals[1], 10, 8)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(uint8)%uint8(x) == uint8(y)
		}, nil
	},
	"in": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		iset := map[uint8]bool{}
		for _, val := range vals {
			i, err := strconv.ParseUint(val, 10, 8)
			if err != nil {
				return nil, err
			}
			iset[uint8(i)] = true
		}
		return func(v interface{}) bool {
			return iset[v.(uint8)]
		}, nil
	},
	"range": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseUint(vals[0], 10, 8)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseUint(vals[1], 10, 8)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(uint8) >= uint8(x) && v.(uint8) <= uint8(y)
		}, nil
	},
}
