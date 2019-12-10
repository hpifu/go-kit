package hrule

import (
	"fmt"
	"strings"
	"time"
)

func NewDurationRule(expr string) (Rule, error) {
	fun, params := ruleExpr.Tokenizer(expr)

	generator, ok := DurationRuleGenerator[fun]
	if !ok {
		return nil, fmt.Errorf("no such generator: [%v]", fun)
	}
	return generator(params)
}

func RegisterDurationRuleGenerator(fun string, generator RuleGenerator) {
	Int64RuleGenerator[fun] = generator
}

var DurationRuleGenerator = map[string]RuleGenerator{
	">": func(params string) (Rule, error) {
		d, err := time.ParseDuration(params)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(time.Duration) > d
		}, nil
	},
	"<": func(params string) (Rule, error) {
		d, err := time.ParseDuration(params)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(time.Duration) < d
		}, nil
	},
	">=": func(params string) (Rule, error) {
		d, err := time.ParseDuration(params)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(time.Duration) >= d
		}, nil
	},
	"<=": func(params string) (Rule, error) {
		d, err := time.ParseDuration(params)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(time.Duration) <= d
		}, nil
	},
	"in": func(params string) (Rule, error) {
		set := map[time.Duration]bool{}
		for _, val := range strings.Split(params, ",") {
			d, err := time.ParseDuration(val)
			if err != nil {
				return nil, err
			}
			set[d] = true
		}
		return func(v interface{}) bool {
			return set[v.(time.Duration)]
		}, nil
	},
}
