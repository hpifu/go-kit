package rule

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheck(t *testing.T) {
	Convey("test check", t, func() {
		So(Check([][3]interface{}{
			{"邮箱", "hatlonely@foxmail.com", []Rule{ValidEmail, AtMost(64), AtLeast(12)}},
			{"电话", "12345674567", []Rule{ValidPhone}},
		}), ShouldNotBeNil)
	})
}

func TestIn(t *testing.T) {
	Convey("test in", t, func() {
		So(In(1, 2, 3)(1), ShouldBeNil)
		So(In(1, 2, 3)(2), ShouldBeNil)
		So(In(1, 2, 3)(4), ShouldNotBeNil)
		So(In("dog", "cat", "pig")("apple"), ShouldNotBeNil)
	})
}

func TestGreaterLess(t *testing.T) {
	Convey("test greater less", t, func() {
		So(LessThan(10)(9), ShouldBeNil)
		So(GreaterThan(10)(11), ShouldBeNil)
		So(GreaterEqual(10)(10), ShouldBeNil)
		So(GreaterEqual(10)(11), ShouldBeNil)
		So(LessEqual(10)(10), ShouldBeNil)
		So(LessEqual(10)(9), ShouldBeNil)
		So(LessThan(10)(10), ShouldNotBeNil)
		So(LessThan(10)(11), ShouldNotBeNil)
		So(GreaterThan(10)(10), ShouldNotBeNil)
		So(GreaterThan(10)(9), ShouldNotBeNil)
	})
}

func TestValidEmail(t *testing.T) {
	Convey("test email", t, func() {
		So(ValidEmail("hatlonely@foxmail.com"), ShouldBeNil)
		So(ValidEmail("hat.lonely._%+-abc@fox.-mail.com"), ShouldBeNil)
		So(ValidEmail("hatlonely"), ShouldNotBeNil)
		So(ValidEmail("hatlonely$@foxmail.com"), ShouldNotBeNil)
		So(ValidEmail("hatlonely@fox+mail.com"), ShouldNotBeNil)
	})
}

func TestValidPhone(t *testing.T) {
	Convey("test phone", t, func() {
		So(ValidPhone("13112341234"), ShouldBeNil)
		So(ValidPhone("131123456789"), ShouldNotBeNil)
		So(ValidPhone("1311234567"), ShouldNotBeNil)
	})
}

func TestValidCode(t *testing.T) {
	Convey("test code", t, func() {
		So(ValidCode("123456"), ShouldBeNil)
		So(ValidCode("abcdef"), ShouldNotBeNil)
	})
}

func TestValidBirthday(t *testing.T) {
	Convey("test birthday", t, func() {
		So(ValidBirthday("2019-02-12"), ShouldBeNil)
		So(ValidBirthday("1950-02-12"), ShouldBeNil)
		So(ValidBirthday("20190212"), ShouldNotBeNil)
		So(ValidBirthday("2019/02/12"), ShouldNotBeNil)
		So(ValidBirthday("2019-2-2"), ShouldNotBeNil)
		So(ValidBirthday("2030-01-01"), ShouldNotBeNil)
		So(ValidBirthday("1900-01-01"), ShouldNotBeNil)
	})
}

func TestAtLeast(t *testing.T) {
	Convey("test at least", t, func() {
		So(AtLeast8Characters("12345678"), ShouldBeNil)
		So(AtLeast8Characters("123456789"), ShouldBeNil)
		So(AtLeast8Characters("1234567"), ShouldNotBeNil)
	})
}

func TestAtMost(t *testing.T) {
	Convey("test at most", t, func() {
		So(AtMost32Characters("12345678901234567890123456789012"), ShouldBeNil)
		So(AtMost32Characters("1234567890123456789012345678901"), ShouldBeNil)
		So(AtMost32Characters("123456789012345678901234567890123"), ShouldNotBeNil)
	})
}
