package hconf

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestJsonDecoder(t *testing.T) {
	Convey("test yaml decoder", t, func() {
		d, err := NewDecoder("json5")
		So(err, ShouldBeNil)
		storage, err := d.Decode([]byte(`{
	"i": 10,
	"b": true,
	"s": "hello",
	"sub": {
		"f": 123.456,
		"vs": [1, 2]
	}
}`))
		So(err, ShouldBeNil)
		i, err := storage.Get("i")
		So(err, ShouldBeNil)
		So(i, ShouldEqual, 10)
		b, err := storage.Get("b")
		So(err, ShouldBeNil)
		So(b, ShouldBeTrue)
		s, err := storage.Get("s")
		So(err, ShouldBeNil)
		So(s, ShouldEqual, "hello")
		f, err := storage.Get("sub.f")
		So(err, ShouldBeNil)
		So(f, ShouldAlmostEqual, 123.456)
		v0, err := storage.Get("sub.vs[0]")
		So(err, ShouldBeNil)
		So(v0, ShouldEqual, 1)
		v1, err := storage.Get("sub.vs[1]")
		So(err, ShouldBeNil)
		So(v1, ShouldEqual, 2)
		v, err := storage.Get("v")
		So(err, ShouldNotBeNil)
		So(v, ShouldBeNil)
	})
}

func TestYamlDecoder(t *testing.T) {
	Convey("test yaml decoder", t, func() {
		d, err := NewDecoder("yaml")
		So(err, ShouldBeNil)
		storage, err := d.Decode([]byte(`
i: 10
b: true
s: hello
sub:
  f: 123.456
  vs:
    - 1
    - 2
`))
		So(err, ShouldBeNil)
		i, err := storage.Get("i")
		So(err, ShouldBeNil)
		So(i, ShouldEqual, 10)
		b, err := storage.Get("b")
		So(err, ShouldBeNil)
		So(b, ShouldBeTrue)
		s, err := storage.Get("s")
		So(err, ShouldBeNil)
		So(s, ShouldEqual, "hello")
		f, err := storage.Get("sub.f")
		So(err, ShouldBeNil)
		So(f, ShouldAlmostEqual, 123.456)
		v0, err := storage.Get("sub.vs[0]")
		So(err, ShouldBeNil)
		So(v0, ShouldEqual, 1)
		v1, err := storage.Get("sub.vs[1]")
		So(err, ShouldBeNil)
		So(v1, ShouldEqual, 2)
		v, err := storage.Get("v")
		So(err, ShouldNotBeNil)
		So(v, ShouldBeNil)
	})
}
