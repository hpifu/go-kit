package logger

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func Test(t *testing.T) {
	Convey("test logger", t, func() {
		testlog, err := NewTextLogger("test.log", 24*time.Hour)
		So(err, ShouldEqual, nil)
		testlog.Info("hello world")
		testlog.Warn("hello world")
		testlog.Warn("hello world")
	})
}
