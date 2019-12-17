package hstr

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAllAny(t *testing.T) {
	Convey("test all and any", t, func() {
		So(All("1234567890", IsDigit), ShouldBeTrue)
		So(All("12345|67890", IsDigit), ShouldBeFalse)
		So(All("abcdefghijklmnopqrstuvwxyz", IsLower), ShouldBeTrue)
		So(All("abcdefghijklmnOpqrstuvwxyz", IsLower), ShouldBeFalse)
		So(Any("abcdefghijklmnOpqrstuvwxyz", IsUpper), ShouldBeTrue)
		So(Any("1234567890@", func(ch uint8) bool {
			return ch == '$' || ch == '@'
		}), ShouldBeTrue)
	})
}

func TestIsFloat(t *testing.T) {
	Convey("test is number", t, func() {
		So(IsFloat("123"), ShouldBeTrue)
		So(IsFloat("-123"), ShouldBeTrue)
		So(IsFloat("123.456"), ShouldBeTrue)
		So(IsFloat("+123.456"), ShouldBeTrue)
		So(IsFloat("-123.456"), ShouldBeTrue)
		So(IsFloat("+123.456E5"), ShouldBeTrue)
		So(IsFloat("-123.456e5"), ShouldBeTrue)
		So(IsFloat("123a"), ShouldBeFalse)
	})
}

func BenchmarkIsFloat(b *testing.B) {
	b.Run("is float", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = IsFloat("-123.456E789")
		}
	})

	b.Run("is float v1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = IsFloatV1("-123.456E789")
		}
	})

	b.Run("is float v2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = IsFloatV2("-123.456E789")
		}
	})
}

func TestIsIdentifier(t *testing.T) {
	Convey("test is identifier", t, func() {
		So(IsIdentifier("abc"), ShouldBeTrue)
		So(IsIdentifier("abc1"), ShouldBeTrue)
		So(IsIdentifier("abc1abc"), ShouldBeTrue)
		So(IsIdentifier("Abc"), ShouldBeTrue)
		So(IsIdentifier("ABC"), ShouldBeTrue)
		So(IsIdentifier("abcAbc"), ShouldBeTrue)
		So(IsIdentifier("abc1Abc"), ShouldBeTrue)
		So(IsIdentifier("abc1Abc1"), ShouldBeTrue)
		So(IsIdentifier("@abc"), ShouldBeFalse)
		So(IsIdentifier("$abc"), ShouldBeFalse)
		So(IsIdentifier("0abc"), ShouldBeFalse)
	})
}

func TestIsEmail(t *testing.T) {
	Convey("test is email", t, func() {
		So(IsEmail("xx@yy.zz.com"), ShouldBeTrue)
		So(IsEmail("abc@def.com"), ShouldBeTrue)
	})
}

func TestIsPhone(t *testing.T) {
	Convey("test is phone", t, func() {
		So(IsPhone("13112345678"), ShouldBeTrue)
		So(IsPhone("1311234567"), ShouldBeFalse)
		So(IsPhone("12312345678"), ShouldBeFalse)
	})
}
