/**
** @创建时间 : 2021/12/22 16:32
** @作者 : fzy
 */
package xstring

import "unicode"

// 判断是否有大写
func IsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

// 判断是否有小写
func IsLower(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
