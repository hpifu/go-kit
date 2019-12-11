package hrule

import (
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"reflect"
	"testing"
)

func TestRuleExprTokenizer(t *testing.T) {
	Convey("test new cond", t, func() {
		fun, params := ruleExpr.Tokenizer("mod 3,1")
		So(fun, ShouldEqual, "mod")
		So(params, ShouldEqual, "3,1")
		fun, params = ruleExpr.Tokenizer("isdigit")
		So(fun, ShouldEqual, "isdigit")
		So(params, ShouldEqual, "")
	})
}

func TestStringRule(t *testing.T) {
	Convey("test string rule", t, func() {
		RegisterStringRuleGenerator("isip", func(params string) (Rule, error) {
			return func(v interface{}) bool {
				ip := net.ParseIP(v.(string))
				if ip == nil {
					return false
				}
				return true
			}, nil
		})

		cond, err := NewCond("isip", reflect.TypeOf(""))
		So(err, ShouldBeNil)
		So(cond.Evaluate("114.243.208.244"), ShouldBeTrue)
		So(cond.Evaluate("113.456.7.8"), ShouldBeFalse)

		type MyStruct struct {
			S string `hrule:"isip"`
		}
		MyStructRule := MustCompile(&MyStruct{})
		So(MyStructRule.Evaluate(&MyStruct{S: "114.243.208.244"}), ShouldBeTrue)
	})
}
