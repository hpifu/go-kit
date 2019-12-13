package hstring

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNaming(t *testing.T) {
	Convey("test  naming", t, func() {
		Convey("camel name", func() {
			So(CamelName("hello_world"), ShouldEqual, "helloWorld")
			So(CamelName("HelloWorld"), ShouldEqual, "helloWorld")
			So(CamelName("Hello_World"), ShouldEqual, "helloWorld")
			So(CamelName("hello_WorldWorld"), ShouldEqual, "helloWorldWorld")
			So(CamelName("IPAddress"), ShouldEqual, "ipAddress")
			So(CamelName("HELLOWorld"), ShouldEqual, "helloWorld")
			So(CamelName("helloWorld"), ShouldEqual, "helloWorld")
			So(CamelName("MyIPAddress"), ShouldEqual, "myIPAddress")
			So(CamelName("IP-Address"), ShouldEqual, "ipAddress")
			So(CamelName("IP-AddressAPP"), ShouldEqual, "ipAddressAPP")

		})

		Convey("pascal name", func() {
			So(PascalName("hello_world"), ShouldEqual, "HelloWorld")
			So(PascalName("HelloWorld"), ShouldEqual, "HelloWorld")
			So(PascalName("Hello_World"), ShouldEqual, "HelloWorld")
			So(PascalName("hello_WorldWorld"), ShouldEqual, "HelloWorldWorld")
			So(PascalName("IPAddress"), ShouldEqual, "IPAddress")
			So(PascalName("HELLOWorld"), ShouldEqual, "HELLOWorld")
			So(PascalName("helloWorld"), ShouldEqual, "HelloWorld")
			So(PascalName("MyIPAddress"), ShouldEqual, "MyIPAddress")
			So(PascalName("IP-Address"), ShouldEqual, "IPAddress")
			So(PascalName("IP-AddressAPP"), ShouldEqual, "IPAddressAPP")

		})

		Convey("snake name", func() {
			So(SnakeName("hello_world"), ShouldEqual, "hello_world")
			So(SnakeName("HelloWorld"), ShouldEqual, "hello_world")
			So(SnakeName("Hello_World"), ShouldEqual, "hello_world")
			So(SnakeName("hello_WorldWorld"), ShouldEqual, "hello_world_world")
			So(SnakeName("IPAddress"), ShouldEqual, "ip_address")
			So(SnakeName("HELLOWorld"), ShouldEqual, "hello_world")
			So(SnakeName("helloWorld"), ShouldEqual, "hello_world")
			So(SnakeName("MyIPAddress"), ShouldEqual, "my_ip_address")
			So(SnakeName("IP-Address"), ShouldEqual, "ip_address")
			So(SnakeName("IP-AddressAPP"), ShouldEqual, "ip_address_app")
			So(SnakeNameAllCaps("hello_world"), ShouldEqual, "HELLO_WORLD")
		})

		Convey("kebab name", func() {
			So(KebabName("hello_world"), ShouldEqual, "hello-world")
			So(KebabName("HelloWorld"), ShouldEqual, "hello-world")
			So(KebabName("Hello_World"), ShouldEqual, "hello-world")
			So(KebabName("hello_WorldWorld"), ShouldEqual, "hello-world-world")
			So(KebabName("IPAddress"), ShouldEqual, "ip-address")
			So(KebabName("HELLOWorld"), ShouldEqual, "hello-world")
			So(KebabName("helloWorld"), ShouldEqual, "hello-world")
			So(KebabName("MyIPAddress"), ShouldEqual, "my-ip-address")
			So(KebabName("IP-Address"), ShouldEqual, "ip-address")
			So(KebabName("IP-AddressAPP"), ShouldEqual, "ip-address-app")
			So(KebabNameAllCaps("hello_world"), ShouldEqual, "HELLO-WORLD")
		})
	})
}
