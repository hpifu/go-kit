package rule

import (
	"fmt"
	"regexp"
	"time"
)

var EmailRegex *regexp.Regexp
var PhoneRegex *regexp.Regexp
var CodeRegex *regexp.Regexp

func init() {
	PhoneRegex = regexp.MustCompile(`^1[345789][0-9]{9}$`)
	EmailRegex = regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)
	CodeRegex = regexp.MustCompile(`^[0-9]{6}$`)
}

type Rule func(interface{}) error

func Check(items [][3]interface{}) error {
	for _, item := range items {
		key := item[0].(string)
		val := item[1]
		rules := item[2].([]Rule)
		for _, r := range rules {
			if err := r(val); err != nil {
				return fmt.Errorf("key: [%v], val: [%v], err: [%v]", key, val, err)
			}
		}
	}

	return nil
}

func In(sets map[interface{}]struct{}) Rule {
	return func(v interface{}) error {
		if _, ok := sets[v]; !ok {
			return fmt.Errorf("%v 必须在 %v 中", v, sets)
		}
		return nil
	}
}

func Required(v interface{}) error {
	if len(v.(string)) == 0 {
		return fmt.Errorf("必要字段")
	}

	return nil
}

func AtLeast(i int) Rule {
	return func(v interface{}) error {
		if len(v.(string)) < i {
			return fmt.Errorf("至少%v个字符", i)
		}
		return nil
	}
}

func AtMost(i int) Rule {
	return func(v interface{}) error {
		if len(v.(string)) > i {
			return fmt.Errorf("至多%v个字符", i)
		}
		return nil
	}
}

var AtLeast8Characters = AtLeast(8)
var AtMost64Characters = AtMost(64)
var AtMost32Characters = AtMost(32)

func ValidEmail(v interface{}) error {
	if !EmailRegex.MatchString(v.(string)) {
		return fmt.Errorf("无效的邮箱")
	}

	return nil
}

func ValidPhone(v interface{}) error {
	if !PhoneRegex.MatchString(v.(string)) {
		return fmt.Errorf("无效的电话号码")
	}

	return nil
}

func ValidCode(v interface{}) error {
	if !CodeRegex.MatchString(v.(string)) {
		return fmt.Errorf("无效的验证码")
	}

	return nil
}

func ValidBirthday(v interface{}) error {
	birthday, err := time.Parse("2006-01-02", v.(string))
	if err != nil {
		return fmt.Errorf("日期格式错误")
	}
	if birthday.After(time.Now()) {
		return fmt.Errorf("日期超过范围")
	}
	if time.Now().Sub(birthday) > 100*365*24*time.Hour {
		return fmt.Errorf("日期超过范围")
	}

	return nil
}
