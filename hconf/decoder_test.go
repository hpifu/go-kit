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

func TestXmlDecoder(t *testing.T) {
	Convey("test yaml decoder", t, func() {
		d, err := NewDecoder("xml")
		So(err, ShouldBeNil)
		storage, err := d.Decode([]byte(`
<?xml version="1.0" encoding="UTF-8"?>

<i>10</i>
<b>true</b>
<s>hello</s>
<a></a>
<sub>
<f>123.456</f>abc
<vs>
  <v>1</v>
  <v>2</v>
  <v>3</v>
  <v>4</v>
</vs>
</sub>
<songs>
  <song>
    <name>Thunder Road</name>
    <duration>4m49s</duration>
  </song>
  <song>
    <name>Stairway to Heaven</name>
    <duration>8m03s</duration>
  </song>
</songs>
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
		v0, err := storage.Get("sub.vs[0]")
		So(err, ShouldBeNil)
		So(v0, ShouldEqual, "1")
		v1, err := storage.Get("sub.vs[1]")
		So(err, ShouldBeNil)
		So(v1, ShouldEqual, "2")
		v, err := storage.Get("v")
		So(err, ShouldNotBeNil)
		So(v, ShouldBeNil)

		song0name, err := storage.Get("songs[0].name")
		So(err, ShouldBeNil)
		So(song0name, ShouldEqual, "Thunder Road")
		song1name, err := storage.Get("songs[1].name")
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

func TestIniDecoder(t *testing.T) {
	Convey("test toml decoder", t, func() {
		d, err := NewDecoder("ini")
		So(err, ShouldBeNil)
		storage, err := d.Decode([]byte(`
i = 10
b = true
s = hello

[sub]
f = 123.456
vs = 1,2

; comment
[sub.song]
name = "Thunder Road"
duration = 4m49s
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
		v0, err := storage.Get("sub.vs")
		So(err, ShouldBeNil)
		So(v0, ShouldEqual, "1,2")
		name, err := storage.Get("sub.song.name")
		So(err, ShouldBeNil)
		So(name, ShouldEqual, "Thunder Road")
		duration, err := storage.Get("sub.song.duration")
		So(err, ShouldBeNil)
		So(duration, ShouldEqual, "4m49s")
		v, err := storage.Get("v")
		So(err, ShouldNotBeNil)
		So(v, ShouldBeNil)
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
