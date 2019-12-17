package hconf

import (
	"context"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestConsulProvider(t *testing.T) {
	Convey("test consul provider", t, func() {
		p, err := NewConsulProvider("127.0.0.1:8500", "test_namespace/test_service")
		So(err, ShouldBeNil)
		So(p, ShouldNotBeNil)
		buf, err := p.Get()
		So(err, ShouldBeNil)
		fmt.Println(string(buf))

		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(20*time.Second))
		p.EventLoop(ctx)
		defer cancel()

	out:
		for {
			select {
			case <-p.Events():
				buf, err := p.Get()
				fmt.Println(err)
				fmt.Println(string(buf))
			case err := <-p.Errors():
				fmt.Println(err)
			case <-ctx.Done():
				break out
			}
		}
	})
}
