package hflag

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestArgumentParserParse(t *testing.T) {
	Convey("test case1", t, func() {
		flag := NewFlagSet()
		So(flag.addFlag("int-option", "i", "help", "int", true, "0"), ShouldBeNil)
		So(flag.addFlag("str-option", "s", "help", "string", true, ""), ShouldBeNil)
		So(flag.addFlag("key", "k", "help", "float", true, ""), ShouldBeNil)
		So(flag.addFlag("all", "a", "help", "bool", true, ""), ShouldBeNil)
		So(flag.addFlag("user", "u", "help", "bool", true, ""), ShouldBeNil)
		So(flag.addFlag("password", "p", "help", "string", false, "654321"), ShouldBeNil)
		So(flag.addPosFlag("pos1", "help", "string", ""), ShouldBeNil)
		So(flag.addPosFlag("pos2", "help", "string", ""), ShouldBeNil)
		err := flag.parse([]string{
			"pos1",
			"--int-option=123",
			"--str-option", "hello world",
			"-k", "3.14",
			"-au",
			"-p123456",
			"pos2",
		})
		So(err, ShouldBeNil)
		flag.Usage()
	})

	Convey("test case2", t, func() {
		flag := NewFlagSet()
		version := flag.Bool("v", false, "print current version")
		configfile := flag.String("c", "configs/monitor.json", "config file path")
		So(flag.parse(strings.Split("--v", " ")), ShouldBeNil)
		So(*version, ShouldBeTrue)
		So(*configfile, ShouldEqual, "configs/monitor.json")
	})
}
