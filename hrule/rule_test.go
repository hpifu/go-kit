package hrule

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewRule(t *testing.T) {
	Convey("test new cond", t, func() {
		fun, params := ruleExpr.Tokenizer("mod 3,1")
		So(fun, ShouldEqual, "mod")
		So(params, ShouldEqual, "3,1")
		fun, params = ruleExpr.Tokenizer("isdigit")
		So(fun, ShouldEqual, "isdigit")
		So(params, ShouldEqual, "")
	})
}
