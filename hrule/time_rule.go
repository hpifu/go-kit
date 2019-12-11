package hrule

import (
	"fmt"
	"time"
)

func NewTimeRule(expr string) (Rule, error) {
	fun, params := ruleExpr.Tokenizer(expr)

	generator, ok := TimeRuleGenerator[fun]
	if !ok {
		return nil, fmt.Errorf("no such generator: [%v]", fun)
	}
	return generator(params)
}

func RegisterTimeRuleGenerator(fun string, generator RuleGenerator) {
	TimeRuleGenerator[fun] = generator
}

func parseTime(str string) (time.Time, error) {
	var t time.Time
	var err error
	if str == "now" {
		t = time.Now()
	} else if len(str) == 10 {
		t, err = time.Parse("2006-01-02", str)
	} else if len(str) == 19 {
		t, err = time.Parse("2006-01-02T15:04:05", str)
	} else {
		t, err = time.Parse(time.RFC3339, str)
	}
	if err != nil {
		return t, err
	}
	return t, nil
}

var TimeRuleGenerator = map[string]RuleGenerator{
	"before": func(params string) (Rule, error) {
		t, err := parseTime(params)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(time.Time).Before(t)
		}, nil
	},
	"after": func(params string) (Rule, error) {
		t, err := parseTime(params)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return v.(time.Time).After(t)
		}, nil
	},
}
