package hflag

import (
	"flag"
	"fmt"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFlag(t *testing.T) {
	Convey("test flag", t, func() {
		flagSet := flag.NewFlagSet("test flag", flag.ExitOnError)
		i := flagSet.Int("i", 10, "int flag")
		f := flagSet.Float64("f", 11.11, "float flag")
		s := flagSet.String("s", "hello world", "string flag")
		d := flagSet.Duration("d", time.Duration(30)*time.Second, "string flag")
		b := flagSet.Bool("b", false, "bool flag")

		Convey("check default value", func() {
			So(*i, ShouldEqual, 10)
			So(*f, ShouldAlmostEqual, 11.11)
			So(*s, ShouldEqual, "hello world")
			So(*d, ShouldEqual, time.Duration(30)*time.Second)
			So(*b, ShouldBeFalse)
		})

		Convey("parse case success", func() {
			// 支持三种格式，'-' 和 '--' 是等效的
			//  -flag		(只支持 bool)
			// 	-flag value	(不支持 bool)
			//	-flag=value
			err := flagSet.Parse(strings.Split("-b -i 100 -f=12.12 --s golang --d=20s", " "))
			So(err, ShouldBeNil)
			So(*i, ShouldEqual, 100)
			So(*f, ShouldAlmostEqual, 12.12)
			So(*s, ShouldEqual, "golang")
			So(*d, ShouldEqual, time.Duration(20)*time.Second)
			So(*b, ShouldBeTrue)
		})

		Convey("parse case unexpected value", func() {
			// -i 后面期望之后一个参数，但是提供了两个，解析会立即停止，剩下的参数会写入到 args 中
			err := flagSet.Parse(strings.Split("-b -i 100 101 -f 12.12 -s golang -d 20s", " "))
			So(err, ShouldBeNil)
			So(*b, ShouldBeTrue)
			So(*i, ShouldEqual, 100)
			So(*f, ShouldEqual, 11.11)          // not override
			So(*s, ShouldEqual, "hello world")  // not override
			So(*d, ShouldEqual, 30*time.Second) // not override
			So(flagSet.NFlag(), ShouldEqual, 2)
			So(flagSet.NArg(), ShouldEqual, 7)
			So(flagSet.Args(), ShouldResemble, []string{
				"101", "-f", "12.12", "-s", "golang", "-d", "20s",
			})
		})

		Convey("set case", func() {
			err := flagSet.Set("i", "120")
			So(err, ShouldBeNil)
			So(*i, ShouldEqual, 120)
		})

		Convey("visit case", func() {
			// 遍历所有设置过得选项
			flagSet.Visit(func(f *flag.Flag) {
				fmt.Println(f.Name)
			})

			// 遍历所有选项
			flagSet.VisitAll(func(f *flag.Flag) {
				fmt.Println(f.Name)
			})
		})

		Convey("lookup case", func() {
			f := flagSet.Lookup("i")
			So(f.Name, ShouldEqual, "i")
			So(f.DefValue, ShouldEqual, "10")
			So(f.Usage, ShouldEqual, "int flag")
			So(f.Value.String(), ShouldEqual, "10")
		})

		Convey("print defaults", func() {
			flagSet.PrintDefaults()
		})
	})
}
