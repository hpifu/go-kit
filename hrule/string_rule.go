package hrule

import (
	"fmt"
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
			for _, ch := range v.(string) {
				if ch < '0' || ch > '9' {
					return false
				}
			}
			return true
		}, nil
	},
	"isxdigit": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			str := strings.ToLower(v.(string))
			if !strings.HasPrefix(str, "0x") {
				return false
			}
			for i := 2; i < len(str); i++ {
				ch := str[i]
				if !((ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f')) {
					return false
				}
			}
			return true
		}, nil
	},
	"isalnum": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			for _, ch := range v.(string) {
				if !((ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) {
					return false
				}
			}
			return true
		}, nil
	},
	"isalpha": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			for _, ch := range v.(string) {
				if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) {
					return false
				}
			}
			return true
		}, nil
	},
	"islower": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			for _, ch := range v.(string) {
				if !(ch >= 'a' && ch <= 'z') {
					return false
				}
			}
			return true
		}, nil
	},
	"isupper": func(params string) (Rule, error) {
		return func(v interface{}) bool {
			for _, ch := range v.(string) {
				if !(ch >= 'A' && ch <= 'Z') {
					return false
				}
			}
			return true
		}, nil
	},
	"isEmail": func(params string) (Rule, error) {
		re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)
		return func(v interface{}) bool {
			return re.MatchString(v.(string))
		}, nil
	},
	"isPhone": func(params string) (Rule, error) {
		re := regexp.MustCompile(`^1[345789][0-9]{9}$`)
		return func(v interface{}) bool {
			return re.MatchString(v.(string))
		}, nil
	},
	"isCode": func(params string) (Rule, error) {
		re := regexp.MustCompile(`^[0-9]{6}$`)
		return func(v interface{}) bool {
			return re.MatchString(v.(string))
		}, nil
	},
}
