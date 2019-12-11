# hrule

## 简介

`hrule` 被设计用来实现对一个类中的字段定义规则，比如取值范围，大小写，前后缀等等，典型的两种使用场景比如：

1. 配置文件中配置项的约束
2. 请求包体参数检查

## 用法

### 基本用法

在数据成员的 `tag` 里面定义数据成员的规则，调用 `MustCompile` 传入结构体构造规则对象，调用规则对象的 `Evaluate` 方法，判断 `MyStruct` 对象是否符合定义的规则

``` go
type MyStruct struct {
    I int           `hrule:">=5 & <9"`
    S string        `hrule:"hasPrefix hello & atLeast 10"`
    D time.Duration `hrule:">10s & <1h"`
}

var MyStructRule = MustCompile(&MyStruct{})

So(MyStructRule.Evaluate(&MyStruct{
    I: 6,
    S: "hello world",
    D: time.Duration(3000) * time.Second,
}), ShouldBeTrue)

So(MyStructRule.Evaluate(&MyStruct{
    I: 6,
    S: "hello world",
    D: time.Duration(3600) * time.Second,
}), ShouldBeFalse)

So(MyStructRule.Evaluate(&MyStruct{
    I: 4,
    S: "hello world",
    D: time.Duration(3000) * time.Second,
}), ShouldBeFalse)
```

### 规则拓展

通过 `RegisterXXXRuleGenerator` 函数可以实现对规则的拓展，下面例子展示了检查字符串为 ip 的拓展

``` go
RegisterStringRuleGenerator("isip", func(params string) (Rule, error) {
    return func(v interface{}) bool {
        ip := net.ParseIP(v.(string))
        if ip == nil {
            return false
        }
        return true
    }, nil
})

cond, err := NewCond("isip", reflect.TypeOf(""))
So(err, ShouldBeNil)
So(cond.Evaluate("114.243.208.244"), ShouldBeTrue)
So(cond.Evaluate("113.456.7.8"), ShouldBeFalse)

type MyStruct struct {
    S string `hrule:"isip"`
}
MyStructRule := MustCompile(&MyStruct{})
So(MyStructRule.Evaluate(&MyStruct{S: "114.243.208.244"}), ShouldBeTrue)
```

## 默认规则

### 数值类型

- `>`: 大于
- `<`: 小于
- `>=`: 大于等于
- `<=`: 小于等于
- `in(vals...)`: 是 vals 中的一个
- `mod(a,b)`: 对 a 取模得 b
- `range(a,b)`: 在 [a,b] 区间内

### 字符串

- `hasPrefix <prefix>`: 前缀
- `hasSuffix <suffix>`: 后缀
- `in <vals...>`: 是 vals 中的一个
- `contains <sub>`: 包含子字符串
- `regex <re>`: 符合正则表达式
- `atMost <len>`: 字符串最长长度
- `atLeast <len>`: 字符串最短长度
- `isdigit`: regex([0-9]+)
- `isalnum`: regex([0-9a-zA-Z]+)
- `isalpha`: regex([a-zA-Z]+)
- `isxdigit`: regex(0x[0-9a-f]+)
- `islower`: regex([a-z]+)
- `isupper`: regex([A-Z]+)
- `isEmail`: regex(^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$)
- `isPhone`: regex(^1[345789][0-9]{9}$)
- `isCode`: regex(^[0-9]{6}$)

### time.Duration

- `>`: 大于
- `<`: 小于
- `>=`: 大于等于
- `<=`: 小于等于
- `in(vals...)`: 是 vals 中的一个

### time.Time

- `before`: 在某个时间之前
- `after`: 在某个时间之后
