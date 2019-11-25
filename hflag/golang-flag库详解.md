# golang flag 库详解

flag 库实现了对命令行参数的解析

## 基本用法

``` go
package main

import (
    "fmt"
    "flag"
)

func main() {
    b := flag.Bool("b", false, "bool flag")
	s := flag.String("s", "hello golang", "string flag")
	flag.Parse()
    fmt.Println("b is", *b)
    fmt.Println("s is", *s)
}
```

上面代码指定了两个选项:

- `bool` 类型的 `b` 选项，默认值为 `false`，帮助信息 `bool flag`
- `string` 类型的 `s` 选项，默认值为 `hello golang`，帮助信息 `string flag`

执行 `go run main.go` 将输出 b 和 s 的值

```
b is false
s is hello golang
```

执行 `go run main.go -b -s "hello world"` 将修改 b 和 s 的值

```
b is true
s is hello world
```

执行 `go run main.go -h` 可以打印帮助信息

```
Usage of main:
  -b	bool flag
  -s string
    	string flag (default "hello golang")
```


## 命令行语法

```
-b -i 100 -f=12.12 --s golang --d=20s
```

- 以 `-` 或者 `--` 开头指定选项名，`-` 和 `--` 是等效的
- 非 `bool` 选项后面需要接一个值赋值 `-flag value`，`-flag=value`，`--falg value`，`--flag=value`
- `bool` 选项后面可以不赋值 `-flag`，`-flag=false`，`--flag`，`--flag=1`
- 解析过程中遇到非选项字段，将立即结束解析，后面的字段会被放入到 `args` 变量中
- 所有的选项必须是已定义的，遇到未知的选项将返回错误

## 主要接口

### CommandLine

```go
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
```

FlagSet 是 flag 的核心解析类，为了方便使用，flag 内部提供了一个 FlagSet 对象 CommandLine，并且为 CommandLine 的所有方法封装了一下直接对外

例如: `flag.Int("i", 10, "int flag")` 其实是 `CommandLine.Int("i", 10, "int flag")`

``` go
func Int(name string, value int, usage string) *int {
	return CommandLine.Int(name, value, usage)
}
```

### 添加选项

``` go
flagSet := flag.NewFlagSet("test flag", flag.ExitOnError)

i := flagSet.Int("i", 10, "int flag")
f := flagSet.Float64("f", 11.11, "float flag")
s := flagSet.String("s", "hello world", "string flag")
d := flagSet.Duration("d", time.Duration(30)*time.Second, "string flag")
b := flagSet.Bool("b", false, "bool flag")

So(*i, ShouldEqual, 10)
So(*f, ShouldAlmostEqual, 11.11)
So(*s, ShouldEqual, "hello world")
So(*d, ShouldEqual, time.Duration(30)*time.Second)
So(*b, ShouldBeFalse)
```

flag 提供了两种方式添加选项

``` go
func Int(name string, value int, usage string) *int
func IntVar(p *int, name string, value int, usage string)
```

- `type` 返回一个对应类型的地址，调用 Parse 之后将会被重新复制
- `typeVar` 没有返回值，传入一个地址，Parse 之后的值写入到传入的地址

### 解析

正常的解析过程

```go
err := flagSet.Parse(strings.Split("-b -i 100 -f=12.12 --s golang --d=20s", " "))
So(err, ShouldBeNil)
So(*i, ShouldEqual, 100)
So(*f, ShouldAlmostEqual, 12.12)
So(*s, ShouldEqual, "golang")
So(*d, ShouldEqual, time.Duration(20)*time.Second)
So(*b, ShouldBeTrue)
```

解析过程中遇到非选项参数

``` go
// -i 后面期望之后一个参数，但是提供了两个，解析会立即停止，剩下的参数会写入到 args 中
err := flagSet.Parse(strings.Split("-b -i 100 101 -f 12.12 -s golang -d 20s", " "))
So(err, ShouldBeNil)
So(*b, ShouldBeTrue)
So(*i, ShouldEqual, 100)
So(*f, ShouldEqual, 11.11)          // not override
So(*s, ShouldEqual, "hello world")  // not override
So(*d, ShouldEqual, 30*time.Second) // not override
```

遇到未知的选项将出错，下面代码将直接返回错误

flag 提供三种错误处理的方式:

- `ContinueOnError`: 通过 Parse 的返回值返回错误
- `ExitOnError`: 调用 os.Exit(2) 直接退出程序，这是默认的处理方式
- `PanicOnError`: 调用 panic 抛出错误

``` go
flagSet.Parse(strings.Split("-xx abc", " "))
```

### 解析状态

``` go
err := flagSet.Parse(strings.Split("-b -i 100 101 -f 12.12 -s golang -d 20s", " "))
So(flagSet.NFlag(), ShouldEqual, 2)
So(flagSet.NArg(), ShouldEqual, 7)
So(flagSet.Args(), ShouldResemble, []string{
    "101", "-f", "12.12", "-s", "golang", "-d", "20s",
})
```

- `NFlag`: 解析了多少个选项
- `NArg`: 剩下多少个参数没有被解析
- `Args`: 返回剩下的参数

### 设置单个选项

``` go
err := flagSet.Set("i", "120")
So(err, ShouldBeNil)
So(*i, ShouldEqual, 120)
```

### 遍历和查询选项

``` go
// 遍历所有设置过得选项
flagSet.Visit(func(f *flag.Flag) {
    fmt.Println(f.Name)
})

// 遍历所有选项
flagSet.VisitAll(func(f *flag.Flag) {
    fmt.Println(f.Name)
})

f := flagSet.Lookup("i")
So(f.Name, ShouldEqual, "i")
So(f.DefValue, ShouldEqual, "10")
So(f.Usage, ShouldEqual, "int flag")
So(f.Value.String(), ShouldEqual, "10")
```

## 设计思路

### 解析过程

``` go
func (f *FlagSet)Parse([]string) error
```

1. 遍历字符串数字
2. 检查当前的字符串，如果以 `-` 或者 `--` 开头说明是选项，否则解析就结束了
3. 检查当前字符串中是否有 `=`，如果有 `=`，直接设置选项的值为等号后面的内容
4. 如果没有 `=`，用下一个字符串作为当前选项的值

### value 的设计

如果我们的设计是在解析之后，用户显示调用类似 `GetInt` 之类的方法，再返回一个值，这样就很容易实现了，但是 Value 的设计难度在于在解析之前需要预先为每一个选项返回一个地址，而选项的类型也不是固定的

``` go
type Value interface {
	String() string
	Set(string) error
}
```

`Value` 被设计成了一个接口，为不同的数据类型实现这个接口，返回给用户的地址就是这个接口的实例数据，解析过程中，可以通过 `Set` 方法修改它的值，这个设计确实还挺巧妙的

## 链接

- 测试代码: <>
