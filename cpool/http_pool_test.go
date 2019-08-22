package connpool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	Convey("test connect pool", t, func() {
		hp := NewHttpPool(2, 500*time.Millisecond, 500*time.Millisecond)
		client := hp.Get()

		req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
		So(err, ShouldBeNil)
		res, err := client.Do(req)
		So(err, ShouldBeNil)
		buf, err := ioutil.ReadAll(res.Body)
		So(err, ShouldBeNil)
		fmt.Println(string(buf))
		defer res.Body.Close()

		hp.Put(client)
	})
}
