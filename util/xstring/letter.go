// Package xstring
// @Author fzy
// @Date 2022-04-14 15:04:03
package xstring

import (
	"unicode"
)

// IsUpper
// @Description: 判断是否有大写
// @param s
// @return bool
func IsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

// IsLower
// @Description: 判断是否有小写
// @param s
// @return bool
func IsLower(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
