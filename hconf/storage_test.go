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

func TestHConfParse(t *testing.T) {
	Convey("test conf parse", t, func() {
		Convey("pass case1", func() {
			infos, err := parseKey("key1.key2[3][4].key3", ".")
			So(err, ShouldBeNil)
			So(len(infos), ShouldEqual, 5)
			So(infos[0].key, ShouldEqual, "key1")
			So(infos[0].mod, ShouldEqual, MapMod)
			So(infos[1].key, ShouldEqual, "key2")
			So(infos[1].mod, ShouldEqual, MapMod)
			So(infos[2].idx, ShouldEqual, 3)
			So(infos[2].mod, ShouldEqual, ArrMod)
			So(infos[3].idx, ShouldEqual, 4)
			So(infos[3].mod, ShouldEqual, ArrMod)
			So(infos[4].key, ShouldEqual, "key3")
			So(infos[4].mod, ShouldEqual, MapMod)
		})

		Convey("pass case2", func() {
			infos, err := parseKey("[1][2].key3[4].key5", ".")
			So(err, ShouldBeNil)
			So(len(infos), ShouldEqual, 5)
			So(infos[0].idx, ShouldEqual, 1)
			So(infos[0].mod, ShouldEqual, ArrMod)
			So(infos[1].idx, ShouldEqual, 2)
			So(infos[1].mod, ShouldEqual, ArrMod)
			So(infos[2].key, ShouldEqual, "key3")
			So(infos[2].mod, ShouldEqual, MapMod)
			So(infos[3].idx, ShouldEqual, 4)
			So(infos[3].mod, ShouldEqual, ArrMod)
			So(infos[4].key, ShouldEqual, "key5")
			So(infos[4].mod, ShouldEqual, MapMod)
		})

		Convey("fail case1", func() {
			infos, err := parseKey("[1][key2].key3[4].key5", ".")
			So(err, ShouldNotBeNil)
			So(infos, ShouldBeNil)
		})
	})
}
