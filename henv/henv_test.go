package henv

import (
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"os"
	"testing"
)

func TestHEnv(t *testing.T) {
	Convey("test env getter", t, func() {
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

	Convey("test env unmarshal", t, func() {
		type MySubStruct struct {
			Vi        []int
			IPAddress net.IP
		}

		type MyStruct struct {
			I   *int `henv:"INT"`
			S   string
			Sub MySubStruct
		}

		_ = os.Setenv("TEST_INT", "123")
		_ = os.Setenv("TEST_S", "hatlonely")
		_ = os.Setenv("TEST_SUB_VI", "1,2,3")
		_ = os.Setenv("TEST_SUB_IP_ADDRESS", "114.243.208.244")

		he := NewHEnv("TEST")
		ms := &MyStruct{}
		So(he.Unmarshal(ms), ShouldBeNil)
		So(*ms.I, ShouldEqual, 123)
		So(ms.S, ShouldEqual, "hatlonely")
		So(ms.Sub.Vi, ShouldResemble, []int{1, 2, 3})
		So(ms.Sub.IPAddress.String(), ShouldEqual, "114.243.208.244")
	})
}
