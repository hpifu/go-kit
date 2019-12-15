package hconf

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMapStorage(t *testing.T) {
	Convey("test map storage", t, func() {
		s := NewMapStorage(map[string]string{
			"i":      "10",
			"b":      "true",
			"sub.f":  "123.456",
			"sub.vs": "1,2,3",
		})

		Convey("test Get", func() {
			v, err := s.Get("sub.f")
			So(err, ShouldBeNil)
			So(v, ShouldEqual, "123.456")
		})

		Convey("test Set", func() {
			So(s.Set("sub.vs", []int{4, 5, 6}), ShouldBeNil)
			v, err := s.Get("sub.vs")
			So(err, ShouldBeNil)
			So(v, ShouldEqual, "4,5,6")
		})

		Convey("test Sub", func() {
			s, err := s.Sub("sub")
			So(err, ShouldBeNil)
			So(s, ShouldNotBeNil)

			v, err := s.Get("f")
			So(err, ShouldBeNil)
			So(v, ShouldEqual, "123.456")
			v, err = s.Get("vs")
			So(err, ShouldBeNil)
			So(v, ShouldEqual, "1,2,3")
		})

		Convey("test Unmarshal", func() {
			type MySubStruct struct {
				F float64
				V []int `hconf:"vs"`
			}
			type MyStruct struct {
				I   int
				B   bool
				Sub MySubStruct
			}

			ms := &MyStruct{}
			So(s.Unmarshal(ms), ShouldBeNil)
			So(ms.I, ShouldEqual, 10)
			So(ms.B, ShouldBeTrue)
			So(ms.Sub.F, ShouldAlmostEqual, 123.456)
			So(ms.Sub.V, ShouldResemble, []int{1, 2, 3})

			mss := &MySubStruct{}
			sub, err := s.Sub("sub")
			So(err, ShouldBeNil)
			So(sub.Unmarshal(mss), ShouldBeNil)
			So(mss.F, ShouldAlmostEqual, 123.456)
			So(mss.V, ShouldResemble, []int{1, 2, 3})
		})
	})
}

func TestInterfaceStorage(t *testing.T) {
	Convey("test interface storage", t, func() {
		s := NewInterfaceStorage(map[string]interface{}{
			"i": 10,
			"b": true,
			"sub": map[string]interface{}{
				"f":  123.456,
				"vs": []interface{}{1, 2, 3},
			},
		})

		Convey("test Get", func() {
			v, err := s.Get("sub.f")
			So(err, ShouldBeNil)
			So(v, ShouldAlmostEqual, 123.456)
		})

		Convey("test Set", func() {
			So(s.Set("sub.vs", []interface{}{4, 5, 6}), ShouldBeNil)
			v, err := s.Get("sub.vs")
			So(err, ShouldBeNil)
			So(v, ShouldResemble, []interface{}{4, 5, 6})
		})

		Convey("test Sub", func() {
			s, err := s.Sub("sub")
			So(err, ShouldBeNil)
			So(s, ShouldNotBeNil)

			v, err := s.Get("f")
			So(err, ShouldBeNil)
			So(v, ShouldEqual, 123.456)
			v, err = s.Get("vs")
			So(err, ShouldBeNil)
			So(v, ShouldResemble, []interface{}{1, 2, 3})
		})

		Convey("test Unmarshal", func() {
			type MySubStruct struct {
				F float64
				V []int `hconf:"vs"`
			}
			type MyStruct struct {
				I   int
				B   bool
				Sub MySubStruct
			}

			ms := &MyStruct{}
			So(s.Unmarshal(ms), ShouldBeNil)
			So(ms.I, ShouldEqual, 10)
			So(ms.B, ShouldBeTrue)
			So(ms.Sub.F, ShouldAlmostEqual, 123.456)
			So(ms.Sub.V, ShouldResemble, []int{1, 2, 3})

			mss := &MySubStruct{}
			sub, err := s.Sub("sub")
			So(err, ShouldBeNil)
			So(sub.Unmarshal(mss), ShouldBeNil)
			So(mss.F, ShouldAlmostEqual, 123.456)
			So(mss.V, ShouldResemble, []int{1, 2, 3})
		})
	})
}
