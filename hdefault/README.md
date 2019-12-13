# hdefault

## 简介

hdefault 提供对结构体默认值的设置

## 用法

``` go
type MySubStruct struct {
    Int   int     `hdef:"456"`
    Float float32 `hdef:"123.456"`
}

type MyStruct struct {
    Int      int    `hdef:"123"`
    Str      string `hdef:"hello"`
    IntSlice []int  `hdef:"1,2,3"`
    Ignore   int
    Sub1     MySubStruct
    Sub2     *MySubStruct
}
ms := &MyStruct{}
So(SetDefault(ms), ShouldBeNil)
So(ms.Int, ShouldEqual, 123)
So(ms.Str, ShouldEqual, "hello")
So(ms.Ignore, ShouldEqual, 0)
So(ms.IntSlice, ShouldResemble, []int{1, 2, 3})
So(ms.Sub1.Int, ShouldEqual, 456)
So(ms.Sub1.Float, ShouldAlmostEqual, float32(123.456))
So(ms.Sub2.Int, ShouldEqual, 456)
So(ms.Sub2.Float, ShouldEqual, float32(123.456))
```

目前仅提供一个接口 `SetDefault`，根据结构体的 tag 自动设置结构体的值，支持所有 [hstring](../hstring/README.md) 支持的数据类型

## 链接

- `hstring`: <https://github.com/hpifu/go-kit/tree/master/hstring>