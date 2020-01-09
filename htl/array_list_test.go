package htl

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestArrayListPush(t *testing.T) {
	Convey("test push", t, func() {
		l := NewArrayList()

		for i := 0; i < 6; i++ {
			l.AddLast(i)
			So(l.Size(), ShouldEqual, i+1)
			So(l.GetLast(), ShouldEqual, i)
		}

		fmt.Println(l.String())
	})
}

func TestArrayListPop(t *testing.T) {
	Convey("test pop", t, func() {
		l := NewArrayList()

		for i := 0; i < 6; i++ {
			l.AddLast(i)
		}
		fmt.Println(l.String())

		for i := 0; i < 6; i++ {
			So(l.Size(), ShouldEqual, 6-i)
			So(l.GetLast(), ShouldEqual, 6-i-1)
			v := l.RemoveLast()
			So(v, ShouldEqual, 6-i-1)
		}
	})
}

func TestArrayListIterator(t *testing.T) {
	Convey("test iterator", t, func() {
		l := NewArrayList()

		for i := 0; i < 6; i++ {
			l.AddLast(i)
		}

		it := l.Iterator()
		for it.HasNext() {
			fmt.Println(it.Next())
		}

		l.ForEach(func(v interface{}) {
			fmt.Println(v)
		})
	})
}
