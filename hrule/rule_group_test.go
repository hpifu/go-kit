package hrule

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestRuleGroup(t *testing.T) {
	Convey("test rule group", t, func() {
		type MyStruct struct {
			I int           `hrule:">=5 & <9"`
			S string        `hrule:"hasPrefix hello & atLeast 10"`
			D time.Duration `hrule:">10s & <1h"`
		}
		var MyStructRule = MustCompile(&MyStruct{})

		So(MyStructRule.Evaluate(&MyStruct{
			I: 6,
			S: "hello world",
			D: time.Duration(3000) * time.Second,
		}), ShouldBeTrue)

		So(MyStructRule.Evaluate(&MyStruct{
			I: 6,
			S: "hello world",
			D: time.Duration(3600) * time.Second,
		}), ShouldBeFalse)

		So(MyStructRule.Evaluate(&MyStruct{
			I: 4,
			S: "hello world",
			D: time.Duration(3000) * time.Second,
		}), ShouldBeFalse)

		b, err := Evaluate(&MyStruct{
			I: 6,
			S: "hello world",
			D: time.Duration(3000) * time.Second,
		})
		So(err, ShouldBeNil)
		So(b, ShouldBeTrue)
	})
}
