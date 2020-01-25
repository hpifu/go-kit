package stream

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNext(t *testing.T) {
	Convey("test next", t, func() {
		s := Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Filter(func(v interface{}) bool {
			return v.(int)%2 == 0
		})
		for v := s.Next(); v != nil; v = s.Next() {
			fmt.Println(v)
		}
	})
}

func TestIntermediate(t *testing.T) {
	Convey("test intermediate", t, func() {
		{
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Map(func(v interface{}) interface{} {
				return v.(int) * v.(int)
			}).ToSlice(), ShouldResemble, []interface{}{
				1, 4, 9, 16, 25, 36, 49, 64, 81,
			})
		}
		{
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Filter(func(v interface{}) bool {
				return v.(int)%2 == 0
			}).ToSlice(), ShouldResemble, []interface{}{2, 4, 6, 8})
		}
		{
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Limit(3).ToSlice(), ShouldResemble, []interface{}{1, 2, 3})
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Limit(0).ToSlice(), ShouldBeNil)
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Limit(10).ToSlice(), ShouldResemble, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Skip(0).ToSlice(), ShouldResemble, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Skip(6).ToSlice(), ShouldResemble, []interface{}{7, 8, 9})
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Skip(9).ToSlice(), ShouldBeNil)
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Skip(3).Limit(3).ToSlice(), ShouldResemble, []interface{}{4, 5, 6})
		}
		{
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).TakeWhile(func(v interface{}) bool {
				return v.(int) < 5
			}).ToSlice(), ShouldResemble, []interface{}{1, 2, 3, 4})
		}
		{
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).DropWhile(func(v interface{}) bool {
				return v.(int) < 5
			}).ToSlice(), ShouldResemble, []interface{}{5, 6, 7, 8, 9})
		}
		{
			So(Of(4, 1, 2, 1, 2, 4, 3, 3).Distinct().ToSlice(), ShouldResemble, []interface{}{4, 1, 2, 3})
		}
		{
			So(Of(6, 4, 7, 3, 2, 9, 1, 5, 8).Sorted(func(v1 interface{}, v2 interface{}) int {
				return v1.(int) - v2.(int)
			}).ToSlice(), ShouldResemble, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})
		}
	})
}

func TestTerminal(t *testing.T) {
	Convey("test terminal", t, func() {
		{
			Of(1, 2, 3, 4, 5, 6, 7, 8, 9).ForEach(func(v interface{}) {
				fmt.Println(v)
			})
		}
		{
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).AnyMatch(func(v interface{}) bool {
				return v.(int) > 10
			}), ShouldBeFalse)
		}
		{
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).AllMatch(func(v interface{}) bool {
				return v.(int) < 10
			}), ShouldBeTrue)
		}
		{
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).NoneMatch(func(v interface{}) bool {
				return v.(int) > 10
			}), ShouldBeTrue)
		}
		{
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Count(), ShouldEqual, 9)
		}
		{
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Max(func(v1 interface{}, v2 interface{}) int {
				return v1.(int) - v2.(int)
			}), ShouldEqual, 9)
		}
		{
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Min(func(v1 interface{}, v2 interface{}) int {
				return v1.(int) - v2.(int)
			}), ShouldEqual, 1)
		}
		{
			So(Of(1, 2, 3, 4, 5, 6, 7, 8, 9).Reduce(func(v1 interface{}, v2 interface{}) interface{} {
				return v1.(int) + v2.(int)
			}, 0), ShouldEqual, 45)
		}
	})
}
