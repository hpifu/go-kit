package hhttp

import (
	"fmt"
	"github.com/hpifu/go-account/pkg/account"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func Test(t *testing.T) {
	Convey("test", t, func() {
		hc := NewHttpClient(20, 200*time.Millisecond, 200*time.Millisecond)
		req := &account.SignInReq{
			Username: "hatlonely@foxmail.com",
			Password: "12345678",
		}
		res := &account.SignInRes{}
		err := hc.POST("http://127.0.0.1:16060/signin", req, res)
		fmt.Println(res)
		So(err, ShouldBeNil)
	})
}
