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
	},
	"song": [{
		"name": "Thunder Road",
		"duration": "4m49s"
	}, {
		"name": "Stairway to Heaven",
		"duration": "8m03s"
	}]
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

		song0name, err := storage.Get("song[0].name")
		So(err, ShouldBeNil)
		So(song0name, ShouldEqual, "Thunder Road")
		song1name, err := storage.Get("song[1].name")
		So(err, ShouldBeNil)
		So(song1name, ShouldEqual, "Stairway to Heaven")
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

song:
  - name: Thunder Road
    duration: 4m49s
  - name: Stairway to Heaven
    duration: 8m03s
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

		song0name, err := storage.Get("song[0].name")
		So(err, ShouldBeNil)
		So(song0name, ShouldEqual, "Thunder Road")
		song1name, err := storage.Get("song[1].name")
		So(err, ShouldBeNil)
		So(song1name, ShouldEqual, "Stairway to Heaven")
	})
}

func TestTomlDecoder(t *testing.T) {
	Convey("test toml decoder", t, func() {
		d, err := NewDecoder("toml")
		So(err, ShouldBeNil)
		storage, err := d.Decode([]byte(`
i = 10
b = true
s = "hello"

[sub]
f = 123.456
vs = [1, 2]

[[song]]
name = "Thunder Road"
duration = "4m49s"

[[song]]
name = "Stairway to Heaven"
duration = "8m03s"
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

		song0name, err := storage.Get("song[0].name")
		So(err, ShouldBeNil)
		So(song0name, ShouldEqual, "Thunder Road")
		song1name, err := storage.Get("song[1].name")
		So(err, ShouldBeNil)
		So(song1name, ShouldEqual, "Stairway to Heaven")
	})
}

func TestPropertiesDecoder(t *testing.T) {
	Convey("test toml decoder", t, func() {
		d, err := NewDecoder("properties")
		So(err, ShouldBeNil)
		storage, err := d.Decode([]byte(`
i = 10
b = true
# ignore
s = hello

sub.vs = 1,\
		2,\
		3

sub.f = 123.456
`))
		So(err, ShouldBeNil)
		i, err := storage.Get("i")
		So(err, ShouldBeNil)
		So(i, ShouldEqual, "10")
		b, err := storage.Get("b")
		So(err, ShouldBeNil)
		So(b, ShouldEqual, "true")
		s, err := storage.Get("s")
		So(err, ShouldBeNil)
		So(s, ShouldEqual, "hello")
		f, err := storage.Get("sub.f")
		So(err, ShouldBeNil)
		So(f, ShouldEqual, "123.456")
		vs, err := storage.Get("sub.vs")
		So(err, ShouldBeNil)
		So(vs, ShouldEqual, "1,2,3")
		v, err := storage.Get("v")
		So(err, ShouldNotBeNil)
		So(v, ShouldBeNil)
	})
}
