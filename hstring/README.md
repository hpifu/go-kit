# hstring

## 简介

hstring 封装了一些常用的字符串算法，提供统一的接口

## 用法

### 字符串转化

提供字符串和以下类型以及其对应的 `slice` 类型的相互转化，

``` go
"bool", "int", "uint", "int64", "int32", "int16", "int8",
"uint64", "uint32", "uint16", "uint8", "float64", "float32",
"time.Duration", "time.Time", "net.IP"
```

主要提供如下几个接口

- `To<type>(str string) (<type>, error)` 函数族: 字符串像其他类型的转化
- `<type>To(v <type>) string` 函数族: 其他类型向字符串的转化
- `ToString(v interface{}) string`: 任意类型向字符串的转化
- `ToInterface(str string, v interface{}) error`: 字符串转任意类型
- `SetValue(v reflect.Value, str string) error`: 用字符串给 reflect.Value 赋值
