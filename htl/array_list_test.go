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
			l.PushBack(i)
			So(l.Len(), ShouldEqual, i+1)
			So(l.Back(), ShouldEqual, i)
		}

		fmt.Println(l.String())
	})
}

func TestArrayListPop(t *testing.T) {
	Convey("test pop", t, func() {
		l := NewArrayList()

		for i := 0; i < 6; i++ {
			l.PushBack(i)
		}
		fmt.Println(l.String())

		for i := 0; i < 6; i++ {
			So(l.Len(), ShouldEqual, 6-i)
			So(l.Back(), ShouldEqual, 6-i-1)
			v := l.PopBack()
			So(v, ShouldEqual, 6-i-1)
		}
	})
}

func TestArrayListIterator(t *testing.T) {
	Convey("test iterator", t, func() {
		l := NewArrayList()

		for i := 0; i < 6; i++ {
			l.PushBack(i)
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
