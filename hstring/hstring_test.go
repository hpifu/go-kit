package hstring

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"strconv"
	"testing"
	"time"
)

func TestHString(t *testing.T) {
	Convey("test hstring", t, func() {
		{
			v, err := ToBool("true")
			So(err, ShouldBeNil)
			So(v, ShouldBeTrue)
			v, err = ToBool("false")
			So(err, ShouldBeNil)
			So(v, ShouldBeFalse)
		}
		{
			v, err := ToInt("1234")
			So(err, ShouldBeNil)
			So(v, ShouldEqual, 1234)
			v, err = ToInt("-1234")
			So(err, ShouldBeNil)
			So(v, ShouldEqual, -1234)
		}
		{
			v, err := ToInt64(strconv.FormatInt(math.MaxInt64, 10))
			So(err, ShouldBeNil)
			So(v, ShouldEqual, math.MaxInt64)
			v, err = ToInt64(strconv.FormatInt(math.MinInt64, 10))
			So(err, ShouldBeNil)
			So(v, ShouldEqual, math.MinInt64)
			v, err = ToInt64(strconv.FormatUint(math.MaxUint64, 10))
			So(err, ShouldNotBeNil)
		}
		{
			v, err := ToInt32(strconv.FormatInt(math.MaxInt32, 10))
			So(err, ShouldBeNil)
			So(v, ShouldEqual, math.MaxInt32)
			v, err = ToInt32(strconv.FormatInt(math.MinInt32, 10))
			So(err, ShouldBeNil)
			So(v, ShouldEqual, math.MinInt32)
			v, err = ToInt32(strconv.FormatInt(math.MaxInt32+1, 10))
			So(err, ShouldNotBeNil)
		}
		{
			v, err := ToInt16(strconv.FormatInt(math.MaxInt16, 10))
			So(err, ShouldBeNil)
			So(v, ShouldEqual, math.MaxInt16)
			v, err = ToInt16(strconv.FormatInt(math.MinInt16, 10))
			So(err, ShouldBeNil)
			So(v, ShouldEqual, math.MinInt16)
			v, err = ToInt16(strconv.FormatInt(math.MaxInt16+1, 10))
			So(err, ShouldNotBeNil)
		}
		{
			v, err := ToInt8(strconv.FormatInt(math.MaxInt8, 10))
			So(err, ShouldBeNil)
			So(v, ShouldEqual, math.MaxInt8)
			v, err = ToInt8(strconv.FormatInt(math.MinInt8, 10))
			So(err, ShouldBeNil)
			So(v, ShouldEqual, math.MinInt8)
			v, err = ToInt8(strconv.FormatInt(math.MaxInt8+1, 10))
			So(err, ShouldNotBeNil)
		}
		{
			v, err := ToUint64(strconv.FormatUint(math.MaxUint64, 10))
			So(err, ShouldBeNil)
			So(v, ShouldEqual, uint64(math.MaxUint64))
			v, err = ToUint64("-1")
			So(err, ShouldNotBeNil)
		}
		{
			v, err := ToUint32(strconv.FormatUint(math.MaxUint32, 10))
			So(err, ShouldBeNil)
			So(v, ShouldEqual, math.MaxUint32)
			v, err = ToUint32(strconv.FormatUint(math.MaxUint32+1, 10))
			So(err, ShouldNotBeNil)
			v, err = ToUint32("-1")
			So(err, ShouldNotBeNil)
		}
		{
			v, err := ToUint16(strconv.FormatUint(math.MaxUint16, 10))
			So(err, ShouldBeNil)
			So(v, ShouldEqual, math.MaxUint16)
			v, err = ToUint16(strconv.FormatUint(math.MaxUint16+1, 10))
			So(err, ShouldNotBeNil)
			v, err = ToUint16("-1")
			So(err, ShouldNotBeNil)
		}
		{
			v, err := ToUint8(strconv.FormatUint(math.MaxUint8, 10))
			So(err, ShouldBeNil)
			So(v, ShouldEqual, math.MaxUint8)
			v, err = ToUint8(strconv.FormatUint(math.MaxUint8+1, 10))
			So(err, ShouldNotBeNil)
			v, err = ToUint8("-1")
			So(err, ShouldNotBeNil)
		}
		{
			v, err := ToFloat64(fmt.Sprintf("123.456"))
			So(err, ShouldBeNil)
			So(v, ShouldAlmostEqual, 123.456)
		}
		{
			v, err := ToFloat32(fmt.Sprintf("123.456"))
			So(err, ShouldBeNil)
			So(v, ShouldAlmostEqual, float32(123.456))
		}
		{
			v, err := ToDuration("3s")
			So(err, ShouldBeNil)
			So(v, ShouldEqual, 3*time.Second)
		}
		{
			v, err := ToTime("2019-12-21")
			So(err, ShouldBeNil)
			So(v.Year(), ShouldEqual, 2019)
			So(v.Month(), ShouldEqual, 12)
			So(v.Day(), ShouldEqual, 21)
			So(v.Hour(), ShouldEqual, 0)
			So(v.Minute(), ShouldEqual, 0)
			So(v.Second(), ShouldEqual, 0)

			v, err = ToTime("2019-12-21 18:25:36")
			So(err, ShouldBeNil)
			So(v.Year(), ShouldEqual, 2019)
			So(v.Month(), ShouldEqual, 12)
			So(v.Day(), ShouldEqual, 21)
			So(v.Hour(), ShouldEqual, 18)
			So(v.Minute(), ShouldEqual, 25)
			So(v.Second(), ShouldEqual, 36)

			v, err = ToTime("2019-12-21T18:25:36")
			So(err, ShouldBeNil)
			So(v.Year(), ShouldEqual, 2019)
			So(v.Month(), ShouldEqual, 12)
			So(v.Day(), ShouldEqual, 21)
			So(v.Hour(), ShouldEqual, 18)
			So(v.Minute(), ShouldEqual, 25)
			So(v.Second(), ShouldEqual, 36)
		}
		{
			v, err := ToIP("47.244.104.34")
			So(err, ShouldBeNil)
			So(v.String(), ShouldEqual, "47.244.104.34")
		}
		{
			v, err := ToStringSlice("")
			So(v, ShouldResemble, []string{})
			So(err, ShouldBeNil)
		}
	})
}
