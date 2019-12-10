package hrule

import (
	"fmt"
	"strconv"
	"strings"
)

func NewUint64Rule(expr string) (Rule, error) {
	fun, params := ruleExpr.Tokenizer(expr)

	generator, ok := Uint64RuleGenerator[fun]
	if !ok {
		return nil, fmt.Errorf("no such generator: [%v]", fun)
	}
	return generator(params)
}

func RegisterUint64RuleGenerator(fun string, generator RuleGenerator) {
	Uint64RuleGenerator[fun] = generator
}

var Uint64RuleGenerator = map[string]RuleGenerator{
	">=": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 64)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint64) >= i
		}, nil
	},
	"<=": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 64)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint64) <= i
		}, nil
	},
	">": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 64)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint64) > i
		}, nil
	},
	"<": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 64)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint64) < i
		}, nil
	},
	"mod": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseUint(vals[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseUint(vals[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(uint64)%x == y
		}, nil
	},
	"in": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		iset := map[uint64]bool{}
		for _, val := range vals {
			i, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				return nil, err
			}
			iset[i] = true
		}
		return func(v interface{}) bool {
			return iset[v.(uint64)]
		}, nil
	},
	"range": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseUint(vals[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseUint(vals[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(uint64) >= x && v.(uint64) <= y
		}, nil
	},
}
