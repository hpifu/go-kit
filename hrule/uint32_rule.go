package hrule

import (
	"fmt"
	"strconv"
	"strings"
)

func NewUint32Rule(expr string) (Rule, error) {
	fun, params := ruleExpr.Tokenizer(expr)

	generator, ok := Uint32RuleGenerator[fun]
	if !ok {
		return nil, fmt.Errorf("no such generator: [%v]", fun)
	}
	return generator(params)
}

func RegisterUint32RuleGenerator(fun string, generator RuleGenerator) {
	Uint32RuleGenerator[fun] = generator
}

var Uint32RuleGenerator = map[string]RuleGenerator{
	">=": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 32)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint32) >= uint32(i)
		}, nil
	},
	"<=": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 32)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint32) <= uint32(i)
		}, nil
	},
	">": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 32)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint32) > uint32(i)
		}, nil
	},
	"<": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 32)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint32) < uint32(i)
		}, nil
	},
	"mod": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseUint(vals[0], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseUint(vals[1], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(uint32)%uint32(x) == uint32(y)
		}, nil
	},
	"in": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		iset := map[uint32]bool{}
		for _, val := range vals {
			i, err := strconv.ParseUint(val, 10, 32)
			if err != nil {
				return nil, err
			}
			iset[uint32(i)] = true
		}
		return func(v interface{}) bool {
			return iset[v.(uint32)]
		}, nil
	},
	"range": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseUint(vals[0], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseUint(vals[1], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(uint32) >= uint32(x) && v.(uint32) <= uint32(y)
		}, nil
	},
}
