package hflag

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

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
		err := flagSet.parse([]string{
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
		So(flagSet.parse(strings.Split("--v", " ")), ShouldBeNil)
		So(*version, ShouldBeTrue)
		So(*configfile, ShouldEqual, "configs/monitor.json")
	})
}
