package henv

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestHEnv(t *testing.T) {
	Convey("test env", t, func() {
		_ = os.Setenv("TEST_INT", "123")
		_ = os.Setenv("TEST_INT_SLICE", "1,2,3")
		he := NewHEnv("TEST")

		i, err := he.GetUint("INT")
		So(err, ShouldBeNil)
		So(i, ShouldEqual, 123)

		s, err := he.GetString("INT")
		So(err, ShouldBeNil)
		So(s, ShouldEqual, "123")

		vi, err := he.GetInt32Slice("INT_SLICE")
		So(err, ShouldBeNil)
		So(vi, ShouldResemble, []int32{1, 2, 3})

		v, err := he.GetString("NOT_EXIST_KEY")
		So(err, ShouldEqual, NoSuchKeyErr)
		So(v, ShouldEqual, "")
	})
}
