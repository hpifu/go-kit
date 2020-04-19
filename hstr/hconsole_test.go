package hstr

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConsole(t *testing.T) {
	Convey("test console", t, func() {
		s := NewFontStyle(
			FormatSetBold,
			FormatSetUnderline,
			ForegroundRed,
			BackgroundBlack,
		)
		fmt.Println(s.Render("hello world\n hello world"))
		fmt.Println("abc")
	})
}
