package hrule

import (
	"fmt"
	"github.com/hpifu/go-kit/hstr"
	"regexp"
	"strconv"
	"strings"
)

func NewStringRule(expr string) (Rule, error) {
	fun, params := ruleExpr.Tokenizer(expr)
	generator, ok := StringRuleGenerator[fun]
	if !ok {
		return nil, fmt.Errorf("no such generator: [%v]", fun)
	}
	return generator(params)
}

func RegisterStringRuleGenerator(fun string, generator RuleGenerator) {
	StringRuleGenerator[fun] = generator
}

var StringRuleGenerator = map[string]RuleGenerator{
	"hasPrefix": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			return strings.HasPrefix(v.(string), params)
		}, nil
	},
	"hasSuffix": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			return strings.HasSuffix(v.(string), params)
		}, nil
	},
	"in": func(params string) (Rule, error) {
		set := map[string]bool{}
		for _, val := range strings.Split(params, ",") {
			set[val] = true
		}
		return func(v interface{}) bool {
			return set[v.(string)]
		}, nil
	},
	"contains": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			return strings.Contains(v.(string), params)
		}, nil
	},
	"regex": func(params string) (Rule, error) {
		re, err := regexp.Compile(params)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return re.MatchString(v.(string))
		}, nil
	},
	"atMost": func(params string) (Rule, error) {
		i, err := strconv.Atoi(params)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return len(v.(string)) <= i
		}, nil
	},
	"atLeast": func(params string) (Rule, error) {
		i, err := strconv.Atoi(params)
		if err != nil {
			return nil, err
		}
		return func(v interface{}) bool {
			return len(v.(string)) >= i
		}, nil
	},
	"isdigit": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			return hstr.All(v.(string), hstr.IsDigit)
		}, nil
	},
	"isxdigit": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			str := v.(string)
			if len(str) <= 2 || str[0] != '0' || hstr.ToLower(str[1]) != 'x' {
				return false
			}
			return hstr.All(str[2:], hstr.IsXdigit)
		}, nil
	},
	"isalnum": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			return hstr.All(v.(string), hstr.IsAlnum)
		}, nil
	},
	"isalpha": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			return hstr.All(v.(string), hstr.IsAlpha)
		}, nil
	},
	"islower": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			return hstr.All(v.(string), hstr.IsLower)
		}, nil
	},
	"isupper": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			return hstr.All(v.(string), hstr.IsUpper)
		}, nil
	},
	"isEmail": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			return hstr.IsEmail(v.(string))
		}, nil
	},
	"isPhone": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			return hstr.IsPhone(v.(string))
		}, nil
	},
	"isCode": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			return len(v.(string)) == 6 && hstr.All(v.(string), hstr.IsDigit)
		}, nil
	},
}
