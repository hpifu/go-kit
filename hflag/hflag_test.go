package hflag

import (
	"fmt"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHFlag(t *testing.T) {
	Convey("test flag", t, func() {
		flagSet := NewFlagSet("test flag")
		i := flagSet.Int("i", 10, "int flag")
		//i := new(int)
		//flagSet.IntVar(i, "i", 10, "int flag")
		f := flagSet.Float64("f", 11.11, "float flag")
		s := flagSet.String("s", "hello world", "string flag")
		d := flagSet.Duration("d", time.Duration(30)*time.Second, "string flag")
		b := flagSet.Bool("b", false, "bool flag")
		vi := flagSet.IntSlice("vi", []int{1, 2, 3, 4, 5}, "int slice flag")

		Convey("check default Value", func() {
			So(*i, ShouldEqual, 10)
			So(*f, ShouldAlmostEqual, 11.11)
			So(*s, ShouldEqual, "hello world")
			So(*d, ShouldEqual, time.Duration(30)*time.Second)
			So(*b, ShouldBeFalse)
			So(*vi, ShouldResemble, []int{1, 2, 3, 4, 5})
		})

		Convey("parse case success", func() {
			err := flagSet.Parse(strings.Split("-b -i 100 -f=12.12 --s golang --d=20s -vi 6,7,8,9", " "))
			So(err, ShouldBeNil)
			So(*i, ShouldEqual, 100)
			So(*f, ShouldAlmostEqual, 12.12)
			So(*s, ShouldEqual, "golang")
			So(*d, ShouldEqual, time.Duration(20)*time.Second)
			So(*b, ShouldBeTrue)
			So(*vi, ShouldResemble, []int{6, 7, 8, 9})
		})

		Convey("parse case unexpected Value", func() {
			// -i 后面期望之后一个参数，100 会被当做 -i 的参数，101 会被当成位置参数，继续解析
			err := flagSet.Parse(strings.Split("-b -i 100 101 -f 12.12 -s golang -d 20s", " "))
			So(err, ShouldBeNil)
			So(*b, ShouldBeTrue)
			So(*i, ShouldEqual, 100)
			So(*f, ShouldEqual, 12.12)
			So(*s, ShouldEqual, "golang")
			So(*d, ShouldEqual, 20*time.Second)
			So(flagSet.NFlag(), ShouldEqual, 6)
			So(flagSet.NArg(), ShouldEqual, 1)
			So(flagSet.Args(), ShouldResemble, []string{
				"101",
			})
		})

		Convey("set case", func() {
			err := flagSet.Set("i", "120")
			So(err, ShouldBeNil)
			So(*i, ShouldEqual, 120)
		})

		Convey("visit case", func() {
			// 遍历所有设置过得选项
			flagSet.Visit(func(f *Flag) {
				fmt.Println(f.Name)
			})

			// 遍历所有选项
			flagSet.VisitAll(func(f *Flag) {
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

func TestHFlagParse(t *testing.T) {
	Convey("test case1", t, func() {
		flagSet := NewFlagSet("test flag")
		So(flagSet.AddFlag("int-option", "usage", Shorthand("i"), Type("int"), Required(), DefaultValue("10")), ShouldBeNil)
		So(flagSet.AddFlag("str-option", "usage", Shorthand("s"), Required()), ShouldBeNil)
		So(flagSet.AddFlag("key", "usage", Shorthand("k"), Type("float"), Required()), ShouldBeNil)
		So(flagSet.AddFlag("all", "usage", Shorthand("a"), Type("bool"), Required()), ShouldBeNil)
		So(flagSet.AddFlag("user", "usage", Shorthand("u"), Type("bool"), Required()), ShouldBeNil)
		So(flagSet.AddFlag("password", "usage", Shorthand("p"), Type("string"), DefaultValue("654321")), ShouldBeNil)
		So(flagSet.AddFlag("vs", "usage", Shorthand("v"), Type("[]string"), DefaultValue("dog,cat")), ShouldBeNil)
		So(flagSet.AddPosFlag("pos1", "usage", Type("string")), ShouldBeNil)
		So(flagSet.AddPosFlag("pos2", "usage", Type("string")), ShouldBeNil)

		Convey("check default value", func() {
			So(flagSet.GetInt("int-option"), ShouldEqual, 10)
			So(flagSet.GetString("str-option"), ShouldEqual, "")
			So(flagSet.GetFloat("key"), ShouldEqual, 0.0)
			So(flagSet.GetBool("all"), ShouldBeFalse)
			So(flagSet.GetBool("user"), ShouldBeFalse)
			So(flagSet.GetString("password"), ShouldEqual, "654321")
			So(flagSet.GetStringSlice("vs"), ShouldResemble, []string{"dog", "cat"})
			So(flagSet.GetString("pos1"), ShouldEqual, "")
			So(flagSet.GetString("pos2"), ShouldEqual, "")
		})

		Convey("parse case 1", func() {
			err := flagSet.Parse([]string{
				"val1",
				"--int-option=123",
				"--str-option", "apple,banana,orange",
				"-k", "3.14",
				"-au",
				"-p123456",
				"val2",
				"-vs", "one,two,three",
			})
			So(err, ShouldBeNil)

			So(flagSet.GetInt("int-option"), ShouldEqual, 123)
			So(flagSet.GetString("str-option"), ShouldEqual, "apple,banana,orange")
			So(flagSet.GetStringSlice("str-option"), ShouldResemble, []string{
				"apple", "banana", "orange",
			})
			So(flagSet.GetFloat("key"), ShouldAlmostEqual, 3.14)
			So(flagSet.GetBool("all"), ShouldBeTrue)
			So(flagSet.GetBool("user"), ShouldBeTrue)
			So(flagSet.GetString("password"), ShouldEqual, "123456")
			So(flagSet.GetStringSlice("vs"), ShouldResemble, []string{"one", "two", "three"})
			So(flagSet.GetString("pos1"), ShouldEqual, "val1")
			So(flagSet.GetString("pos2"), ShouldEqual, "val2")

			So(flagSet.Args(), ShouldResemble, []string{
				"val1", "val2",
			})
			So(flagSet.Arg(0), ShouldEqual, "val1")
			So(flagSet.Arg(1), ShouldEqual, "val2")
		})

		Convey("parse case 2", func() {
			err := flagSet.Parse(strings.Split("--str-option=1,2,3,4 -key=3 -au", " "))
			So(err, ShouldBeNil)
			So(flagSet.GetIntSlice("str-option"), ShouldResemble, []int{1, 2, 3, 4})
			So(flagSet.GetString("a"), ShouldEqual, "true")
			So(flagSet.GetString("u"), ShouldEqual, "true")
		})
	})

	Convey("test case2", t, func() {
		flagSet := NewFlagSet("test flag")
		version := flagSet.Bool("v", false, "print current version")
		configfile := flagSet.String("c", "configs/monitor.json", "config file path")
		So(flagSet.Parse(strings.Split("--v", " ")), ShouldBeNil)
		So(*version, ShouldBeTrue)
		So(*configfile, ShouldEqual, "configs/monitor.json")
	})
}
