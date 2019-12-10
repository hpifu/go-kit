package hrule

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestNewCond(t *testing.T) {
	Convey("test new cond", t, func() {
		cond, err := NewCond("(>=3 & <=4) | (>7 & <9)", reflect.TypeOf(int(0)))
		So(err, ShouldBeNil)
		So(cond.Evaluate(3), ShouldBeTrue)
		So(cond.Evaluate(4), ShouldBeTrue)
		So(cond.Evaluate(5), ShouldBeFalse)
		So(cond.Evaluate(6), ShouldBeFalse)
		So(cond.Evaluate(7), ShouldBeFalse)
		So(cond.Evaluate(8), ShouldBeTrue)
		So(cond.Evaluate(9), ShouldBeFalse)
	})
}
