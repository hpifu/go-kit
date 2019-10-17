package hhttp

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func Test(t *testing.T) {
	Convey("test1", t, func() {
		hc := NewHttpClient(20, 200*time.Millisecond, 200*time.Millisecond)
		var res string
		err := hc.GET("http://www.baidu.com", nil, nil, nil).String(&res)
		fmt.Println(res)
		So(err, ShouldBeNil)
	})
}
