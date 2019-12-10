# hrule

`hrule` 被设计用来实现对一个类中的字段定义规则，比如取值范围，大小写，前后缀等等

## 规则

### 数值类型

- `>`: 大于
- `<`: 小于
- `>=`: 大于等于
- `<=`: 小于等于
- `==`: 等于
- `in(vals...)`: 是 vals 中的一个
- `mod(a,b)`: 对 a 取模得 b
- `range(a,b)`: 在 [a,b] 区间内

### 字符串

- `hasPrefix(prefix)`: 前缀
- `hasSuffix(suffix)`: 后缀
- `==`: 等于
- `in(vals...)`: 是 vals 中的一个
- `contains(sub)`: 包含子字符串
- `regex(re)`: 符合正则表达式
- `atMost(len)`: 字符串最长长度
- `atLeast(len)`: 字符串最短长度
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
- `==`: 等于
- `in(vals...)`: 是 vals 中的一个

### time.Time

- `before`: 在某个时间之前
- `after`: 在某个时间之后
