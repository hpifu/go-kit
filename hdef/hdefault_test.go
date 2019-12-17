package hdef

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSetDefault(t *testing.T) {
	Convey("test set", t, func() {
		type MySubStruct struct {
			Int   int     `hdef:"456"`
			Float float32 `hdef:"123.456"`
		}

		type MyStruct struct {
			Int      int    `hdef:"123"`
			Str      string `hdef:"hello"`
			IntSlice []int  `hdef:"1,2,3"`
			Ignore   int
			Sub1     MySubStruct
			Sub2     *MySubStruct
		}
		ms := &MyStruct{}
		So(SetDefault(ms), ShouldBeNil)
		So(ms.Int, ShouldEqual, 123)
		So(ms.Str, ShouldEqual, "hello")
		So(ms.Ignore, ShouldEqual, 0)
		So(ms.IntSlice, ShouldResemble, []int{1, 2, 3})
		So(ms.Sub1.Int, ShouldEqual, 456)
		So(ms.Sub1.Float, ShouldAlmostEqual, float32(123.456))
		So(ms.Sub2.Int, ShouldEqual, 456)
		So(ms.Sub2.Float, ShouldEqual, float32(123.456))
	})
}
