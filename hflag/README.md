# hflag

## 简介

`hflag` 是被设计用来替代标准的 `flag` 库，提供更强大更灵活的命令行解析功能，相比标准库，`hflag` 有如下特点

- 支持可选参数和必选参数
- 支持参数缩写
- 支持位置参数，位置参数可以出现在任意位置
- 支持 bool 参数简写 (`-aux` 和 `-a -u -x` 等效)
- 支持值参数缩写 (`-p123456` 和 `-p 123456` 等效)
- 更多类型的支持，支持 `net.IP`，`time.Time`，`time.Duration`，`[]int`, `[]string` 的解析
- 更友好的用法帮助
- 提供一套更简洁的 api
- 完全兼容 flag 接口

## 用法

`hflag` 提供两套 `api`，一套完全兼容标准库的 `flag` 接口，另一套类似 `python` 的 `argparse` 先定义 `flag`，在使用时从 `flag` 中获取

### 新接口

``` go
package main

import (
	"fmt"
	"github.com/hpifu/go-kit/hflag"
)

func main() {
	hflag.AddFlag("int", "int flag", hflag.Required(), hflag.Shorthand("i"), hflag.Type("int"), hflag.DefaultValue("123"))
	hflag.AddFlag("str", "str flag", hflag.Shorthand("s"), hflag.Required())
	hflag.AddFlag("int-slice", "int slice flag", hflag.Type("[]int"), hflag.DefaultValue("1,2,3"))
	hflag.AddFlag("ip", "ip flag", hflag.Type("ip"))
	hflag.AddFlag("time", "time flag", hflag.Type("time"), hflag.DefaultValue("2019-11-27"))
	hflag.AddPosFlag("pos", "pos flag")
	if err := hflag.Parse(); err != nil {
		panic(err)
	}

	fmt.Println("int =>", hflag.GetInt("i"))
	fmt.Println("str =>", hflag.GetString("s"))
	fmt.Println("int-slice =>", hflag.GetIntSlice("int-slice"))
	fmt.Println("ip =>", hflag.GetIP("ip"))
	fmt.Println("time =>", hflag.GetTime("time"))
	fmt.Println("pos =>", hflag.GetString("pos"))
}
```

`go run hflag1.go -str abc -ip 192.168.0.1 --int-slice 4,5,6 posflag` 将得到如下输出：

```
int => 123
str => abc
int-slice => [4 5 6]
ip => 192.168.0.1
time => 2019-11-27 00:00:00 +0000 UTC
pos => posflag
```

### flag 接口

``` go
package main

import (
	"fmt"
	"github.com/hpifu/go-kit/hflag"
	"time"
)

func main() {
	i := hflag.Int("int", 123, "int flag")
	s := hflag.String("str", "", "str flag")
	vi := hflag.IntSlice("int-slice", []int{1, 2, 3}, "int slice flag")
	ip := hflag.IP("ip", nil, "ip flag")
	t := hflag.Time("time", time.Now(), "time flag")
	if err := hflag.Parse(); err != nil {
		fmt.Println(hflag.Usage())
		panic(err)
	}

	fmt.Println("int =>", *i)
	fmt.Println("str =>", *s)
	fmt.Println("int-slice =>", *vi)
	fmt.Println("ip =>", *ip)
	fmt.Println("time =>", *t)
}
```

`go run hflag2.go -str abc -ip 192.168.0.1 --int-slice 4,5,6 posflag` 将得到如下输出：

```
int => 123
str => abc
int-slice => [4 5 6]
ip => 192.168.0.1
time => 2019-11-27 00:00:00 +0000 UTC
pos => posflag
```

## 帮助信息

hflag 会自动加入 `-h/--help` 选项，用户可以利用该选项，来查看帮助

```
usage: hflag1 [pos] [-h,help bool] [-i,int int=123] [-int-slice []int=1,2,3] [-ip ip] <-s,str string> [-time time=2019-11-27]

positional options:
      pos          [string]           pos flag

options:
  -h, --help       [bool]             show usage
  -i, --int        [int=123]          int flag
    , --int-slice  [[]int=1,2,3]      int slice flag
    , --ip         [ip]               ip flag
  -s, --str        [string]           str flag
    , --time       [time=2019-11-27]  time flag
```

## 命令行语法

- 以 `-` 或者 `--` 开头的字段会被解析成选项，否则会被解析成位置参数，`-` 和 `--` 是等效的
- 选项中如果包含 `=` 号，会按照 `=` 分割成名称和值 (`-key=val`)
- 非 bool 选项后面的字段会被解析成该选项的值 (`-key val`)
- bool 选项后面的字段如果是一个合法的 bool 值，会被当成该选项的值，否则设置为 `true` (`-b true`)
- 如果一个选项中的所有字母都是 bool 选项的缩写，并且所有的选项都设置为 `true`，(`-aux` 和 `-a -u -x` 等效)
- 允许一个选项缩写后面直接加该选项的值 (`-p123456` 和 `-p 123456`)

支持的数据类型名称和 golang 内部类型的对应关系:

- `int` => `int`
- `string` => `string`
- `float` => `float64`
- `bool` => `bool`
- `time` => `time.Duration`
- `duration` => `time.Time`
- `ip` => `net.IP`
- `[]int` => `[]int`
- `[]string` => `[]string`

常用的数据类型格式：

- 合法的 bool 值: `1`, `t`, `T`, `true`, `TRUE`, `True`, `0`, `f`, `F`, `false`, `FALSE`, `False`
- 合法的 time 值: `2019-11-27`, `2019-11-27T00:00:00`, `2019-11-27T00:00:00Z8:00`, `now`
- []int 值: `1,2,3,4`
- []string: `apple,banana`

## 链接

- 例子: <https://github.com/hpifu/go-kit/tree/master/example>
