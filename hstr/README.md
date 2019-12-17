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

### 命名风格转化

- `CamelName(string) string`: 驼峰命名，常用于变量和函数命名，userLoginCount
- `PascalName(str string) string`: 首字母大写的驼峰，常用于类名，UserLoginCount
- `SnakeName(str string) string`: 下划线命名，python 和 c++ 的主要命名方法，user_login_count
- `KebabName(str string) string`: 中划线命名，html 标签命名，user-login-count
- `SnakeNameAllCaps(str string) string`: 全大写加下划线命名，C语言常量命名，系统环境变量命名，USER_LOGIN_COUNT
- `KebabNameAllCaps(str string) string`: 全大写加中划线，USER-LOGIN-COUNT
