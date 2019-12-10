package hrule

import (
	"fmt"
	"strconv"
	"strings"
)

func NewUint16Rule(expr string) (Rule, error) {
	fun, params := ruleExpr.Tokenizer(expr)

	generator, ok := Uint16RuleGenerator[fun]
	if !ok {
		return nil, fmt.Errorf("no such generator: [%v]", fun)
	}
	return generator(params)
}

func RegisterUint16RuleGenerator(fun string, generator RuleGenerator) {
	Uint16RuleGenerator[fun] = generator
}

var Uint16RuleGenerator = map[string]RuleGenerator{
	">=": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 16)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint16) >= uint16(i)
		}, nil
	},
	"<=": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 16)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint16) <= uint16(i)
		}, nil
	},
	">": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 16)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint16) > uint16(i)
		}, nil
	},
	"<": func(params string) (Rule, error) {
		i, err := strconv.ParseUint(params, 10, 16)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(uint16) < uint16(i)
		}, nil
	},
	"mod": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseUint(vals[0], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseUint(vals[1], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(uint16)%uint16(x) == uint16(y)
		}, nil
	},
	"in": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		iset := map[uint16]bool{}
		for _, val := range vals {
			i, err := strconv.ParseUint(val, 10, 16)
			if err != nil {
				return nil, err
			}
			iset[uint16(i)] = true
		}
		return func(v interface{}) bool {
			return iset[v.(uint16)]
		}, nil
	},
	"range": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.ParseUint(vals[0], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.ParseUint(vals[1], 10, 16)
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(uint16) >= uint16(x) && v.(uint16) <= uint16(y)
		}, nil
	},
}
