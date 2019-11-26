package hflag

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	"time"
)

func TestHFlag(t *testing.T) {
	Convey("test flag", t, func() {
		flagSet := NewFlagSet()
		i := flagSet.Int("i", 10, "int flag")
		f := flagSet.Float("f", 11.11, "float flag")
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

		//Convey("parse case unexpected value", func() {
		//	// -i 后面期望之后一个参数，但是提供了两个，解析会立即停止，剩下的参数会写入到 args 中
		//	err := flagSet.Parse(strings.Split("-b -i 100 101 -f 12.12 -s golang -d 20s", " "))
		//	So(err, ShouldBeNil)
		//	So(*b, ShouldBeTrue)
		//	So(*i, ShouldEqual, 100)
		//	So(*f, ShouldEqual, 11.11)          // not override
		//	So(*s, ShouldEqual, "hello world")  // not override
		//	So(*d, ShouldEqual, 30*time.Second) // not override
		//	So(flagSet.NFlag(), ShouldEqual, 2)
		//	So(flagSet.NArg(), ShouldEqual, 7)
		//	So(flagSet.Args(), ShouldResemble, []string{
		//		"101", "-f", "12.12", "-s", "golang", "-d", "20s",
		//	})
		//})
		//
		//Convey("set case", func() {
		//	err := flagSet.Set("i", "120")
		//	So(err, ShouldBeNil)
		//	So(*i, ShouldEqual, 120)
		//})
		//
		//Convey("visit case", func() {
		//	// 遍历所有设置过得选项
		//	flagSet.Visit(func(f *flag.Flag) {
		//		fmt.Println(f.Name)
		//	})
		//
		//	// 遍历所有选项
		//	flagSet.VisitAll(func(f *flag.Flag) {
		//		fmt.Println(f.Name)
		//	})
		//})
		//
		//Convey("lookup case", func() {
		//	f := flagSet.Lookup("i")
		//	So(f.Name, ShouldEqual, "i")
		//	So(f.DefValue, ShouldEqual, "10")
		//	So(f.Usage, ShouldEqual, "int flag")
		//	So(f.Value.String(), ShouldEqual, "10")
		//})
		//
		//Convey("print defaults", func() {
		//	flagSet.PrintDefaults()
		//})
	})
}

func TestHFlagParse(t *testing.T) {
	Convey("test case1", t, func() {
		flagSet := NewFlagSet()
		So(flagSet.addFlag("int-option", "i", "help", "int", true, "0"), ShouldBeNil)
		So(flagSet.addFlag("str-option", "s", "help", "string", true, ""), ShouldBeNil)
		So(flagSet.addFlag("key", "k", "help", "float", true, ""), ShouldBeNil)
		So(flagSet.addFlag("all", "a", "help", "bool", true, ""), ShouldBeNil)
		So(flagSet.addFlag("user", "u", "help", "bool", true, ""), ShouldBeNil)
		So(flagSet.addFlag("password", "p", "help", "string", false, "654321"), ShouldBeNil)
		So(flagSet.addPosFlag("pos1", "help", "string", ""), ShouldBeNil)
		So(flagSet.addPosFlag("pos2", "help", "string", ""), ShouldBeNil)
		err := flagSet.Parse([]string{
			"pos1",
			"--int-option=123",
			"--str-option", "hello world",
			"-k", "3.14",
			"-au",
			"-p123456",
			"pos2",
		})
		So(err, ShouldBeNil)
		flagSet.Usage()
	})

	Convey("test case2", t, func() {
		flagSet := NewFlagSet()
		version := flagSet.Bool("v", false, "print current version")
		configfile := flagSet.String("c", "configs/monitor.json", "config file path")
		So(flagSet.Parse(strings.Split("--v", " ")), ShouldBeNil)
		So(*version, ShouldBeTrue)
		So(*configfile, ShouldEqual, "configs/monitor.json")
	})
}
