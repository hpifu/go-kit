package hconf

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestWatch(t *testing.T) {
	Convey("test watch", t, func() {
		CreateFile()
		h, err := NewHConfWithFile("test.json")
		So(err, ShouldBeNil)
		h.RegisterOnChangeHandler(func(h *HConf) {
			time.Sleep(3 * time.Second)
			fmt.Println("hello world")
		})
		So(h.Watch(), ShouldBeNil)
		time.Sleep(60 * time.Second)
		h.Stop()
		DeleteFile()
	})
}
