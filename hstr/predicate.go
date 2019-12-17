package hstr

import (
	"regexp"
	"strconv"
)

func IsUpper(ch uint8) bool {
	return ch >= 'A' && ch <= 'Z'
}

func IsLower(ch uint8) bool {
	return ch >= 'a' && ch <= 'z'
}

func IsDigit(ch uint8) bool {
	return ch >= '0' && ch <= '9'
}

func All(str string, op func(uint8) bool) bool {
	for i := range str {
		if !op(str[i]) {
			return false
		}
	}

	return true
}

func Any(str string, op func(uint8) bool) bool {
	for i := range str {
		if op(str[i]) {
			return true
		}
	}

	return false
}

func IsIntNum(str string) bool {
	return All(str, IsDigit)
}

func IsNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err != nil
}

var identifierRegex = regexp.MustCompile(`\w[0-9\w]+`)

func IsIdentifier(str string) bool {
	return identifierRegex.Match([]byte(str))
}
