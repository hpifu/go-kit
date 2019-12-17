package htl

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBitSet(t *testing.T) {
	Convey("test bit set", t, func() {
		{
			bs := NewBitSet(63)
			So(bs.capacity, ShouldEqual, 63)
			So(len(bs.bits), ShouldEqual, 1)
			So(bs.Has(60), ShouldBeFalse)
			bs.Add(60)
			So(bs.Has(60), ShouldBeTrue)
			So(bs.Has(62), ShouldBeFalse)
			bs.Add(62)
			So(bs.Has(62), ShouldBeTrue)
			s1, _ := bs.MarshalJSON()
			fmt.Println(string(s1))
			{
				var nb BitSet
				So(nb.UnmarshalJSON(s1), ShouldBeNil)
				So(nb.capacity, ShouldEqual, 63)
				So(nb.bits, ShouldResemble, nb.bits)
				s2, _ := nb.MarshalJSON()
				fmt.Println(string(s2))
				So(s1, ShouldResemble, s2)
			}
		}
		{
			bs := NewBitSet(64)
			So(bs.capacity, ShouldEqual, 64)
			So(len(bs.bits), ShouldEqual, 1)
			So(bs.Has(60), ShouldBeFalse)
			bs.Add(60)
			So(bs.Has(60), ShouldBeTrue)
			So(bs.Has(62), ShouldBeFalse)
			bs.Add(62)
			So(bs.Has(62), ShouldBeTrue)
			s1, _ := bs.MarshalJSON()
			fmt.Println(string(s1))
			{
				var nb BitSet
				So(nb.UnmarshalJSON(s1), ShouldBeNil)
				So(nb.capacity, ShouldEqual, 64)
				So(nb.bits, ShouldResemble, nb.bits)
				s2, _ := nb.MarshalJSON()
				fmt.Println(string(s2))
				So(s1, ShouldResemble, s2)
			}
		}
		{
			bs := NewBitSet(65)
			So(bs.capacity, ShouldEqual, 65)
			So(len(bs.bits), ShouldEqual, 2)
			So(bs.Has(60), ShouldBeFalse)
			bs.Add(60)
			So(bs.Has(60), ShouldBeTrue)
			So(bs.Has(62), ShouldBeFalse)
			bs.Add(62)
			So(bs.Has(62), ShouldBeTrue)
			So(bs.Has(64), ShouldBeFalse)
			bs.Add(64)
			So(bs.Has(64), ShouldBeTrue)
			s1, _ := bs.MarshalJSON()
			fmt.Println(string(s1))
			{
				var nb BitSet
				So(nb.UnmarshalJSON(s1), ShouldBeNil)
				So(nb.capacity, ShouldEqual, 65)
				So(nb.bits, ShouldResemble, nb.bits)
				s2, _ := nb.MarshalJSON()
				fmt.Println(string(s2))
				So(s1, ShouldResemble, s2)
			}
		}
	})
}

func BenchmarkSetContains(b *testing.B) {
	bs := NewBitSet(64)
	hs := map[int]struct{}{}
	for _, i := range []int{1, 2, 4, 10} {
		bs.Add(i)
		hs[i] = struct{}{}
	}

	b.Run("bs", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for i := 0; i < 10; i++ {
				_ = bs.Has(i)
			}
		}
	})

	b.Run("ns", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for i := 0; i < 10; i++ {
				_, _ = hs[i]
			}
		}
	})
}
