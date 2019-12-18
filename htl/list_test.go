package htl

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestListPush(t *testing.T) {
	Convey("test push", t, func() {
		l := NewList()

		for i := 0; i < 3; i++ {
			l.PushFront(i)
			So(l.Len(), ShouldEqual, i+1)
			So(l.Front(), ShouldEqual, i)
		}
		for i := 3; i < 6; i++ {
			l.PushBack(i)
			So(l.Len(), ShouldEqual, i+1)
			So(l.Back(), ShouldEqual, i)
		}

		fmt.Println(l.String())
	})
}

func TestListPop(t *testing.T) {
	Convey("test pop", t, func() {
		l := NewList()

		for i := 0; i < 6; i++ {
			l.PushBack(i)
		}
		fmt.Println(l.String())

		for i := 0; i < 6; i++ {
			So(l.Len(), ShouldEqual, 6-i)
			So(l.Front(), ShouldEqual, i)
			v := l.PopFront()
			So(v, ShouldEqual, i)
		}
	})
}

func TestListIterator(t *testing.T) {
	Convey("test iterator", t, func() {
		l := NewList()

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
