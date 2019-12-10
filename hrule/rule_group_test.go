package hrule

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestRuleGroup(t *testing.T) {
	Convey("test rule group", t, func() {
		type info struct {
			I int           `hrule:">=5 & <9"`
			S string        `hrule:"hasPrefix hello & atLeast 10"`
			D time.Duration `hrule:">10s & <1h"`
		}

		rg, err := Compile(&info{})
		So(err, ShouldBeNil)
		So(rg.Evaluate(&info{
			I: 6,
			S: "hello world",
			D: time.Duration(3000) * time.Second,
		}), ShouldBeTrue)

		So(rg.Evaluate(&info{
			I: 6,
			S: "hello world",
			D: time.Duration(3600) * time.Second,
		}), ShouldBeFalse)

		So(rg.Evaluate(&info{
			I: 4,
			S: "hello world",
			D: time.Duration(3000) * time.Second,
		}), ShouldBeFalse)

		b, err := Evaluate(&info{
			I: 6,
			S: "hello world",
			D: time.Duration(3000) * time.Second,
		})
		So(err, ShouldBeNil)
		So(b, ShouldBeTrue)
	})
}
