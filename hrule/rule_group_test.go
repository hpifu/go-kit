package hrule

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestRuleGroup(t *testing.T) {
	Convey("test rule group", t, func() {
		type MySubStruct struct {
			F string `hrule:"in a,b,c"`
		}
		type MyStruct struct {
			I    int           `hrule:">=5 & <9"`
			S    string        `hrule:"hasPrefix hello & atLeast 10"`
			D    time.Duration `hrule:">10s & <1h"`
			Sub1 *MySubStruct
			Sub2 *MySubStruct `hrule:"-"`
		}
		var MyStructRule = MustCompile(&MyStruct{})

		So(MyStructRule.Evaluate(&MyStruct{
			I: 6,
			S: "hello world",
			D: time.Duration(3000) * time.Second,
			Sub1: &MySubStruct{
				F: "a",
			},
			Sub2: &MySubStruct{
				F: "x",
			},
		}), ShouldBeNil)

		So(MyStructRule.Evaluate(&MyStruct{
			I: 6,
			S: "hello world",
			D: time.Duration(3600) * time.Second,
			Sub1: &MySubStruct{
				F: "a",
			},
		}), ShouldNotBeNil)

		So(MyStructRule.Evaluate(&MyStruct{
			I: 4,
			S: "hello world",
			D: time.Duration(3000) * time.Second,
			Sub1: &MySubStruct{
				F: "a",
			},
		}), ShouldNotBeNil)

		err := Evaluate(&MyStruct{
			I: 6,
			S: "hello world",
			D: time.Duration(3000) * time.Second,
			Sub1: &MySubStruct{
				F: "a",
			},
		})
		So(err, ShouldBeNil)
	})
}
