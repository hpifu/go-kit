package hrule

import (
	"fmt"
	"strconv"
	"strings"
)

func NewIntRule(expr string) (Rule, error) {
	fun, params := ruleExpr.Tokenizer(expr)

	generator, ok := IntRuleGenerator[fun]
	if !ok {
		return nil, fmt.Errorf("no such generator: [%v]", fun)
	}
	return generator(params)
}

func RegisterIntRuleGenerator(fun string, generator RuleGenerator) {
	Int64RuleGenerator[fun] = generator
}

var IntRuleGenerator = map[string]RuleGenerator{
	">=": func(params string) (Rule, error) {
		i, err := strconv.Atoi(params)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(int) >= i
		}, nil
	},
	"<=": func(params string) (Rule, error) {
		i, err := strconv.Atoi(params)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(int) <= i
		}, nil
	},
	">": func(params string) (Rule, error) {
		i, err := strconv.Atoi(params)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(int) > i
		}, nil
	},
	"<": func(params string) (Rule, error) {
		i, err := strconv.Atoi(params)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(int) < i
		}, nil
	},
	"mod": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.Atoi(vals[0])
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.Atoi(vals[1])
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(int)%x == y
		}, nil
	},
	"in": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		iset := map[int]bool{}
		for _, val := range vals {
			i, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			iset[i] = true
		}
		return func(v interface{}) bool {
			return iset[v.(int)]
		}, nil
	},
	"range": func(params string) (Rule, error) {
		vals := strings.Split(params, ",")
		if len(vals) != 2 {
			return nil, fmt.Errorf("params [%v] should be two number", params)
		}
		x, err := strconv.Atoi(vals[0])
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		y, err := strconv.Atoi(vals[1])
		if err != nil {
			return nil, fmt.Errorf("to number failed. [%v]", vals[0])
		}
		return func(v interface{}) bool {
			return v.(int) >= x && v.(int) <= y
		}, nil
	},
}
