// Package xstring
// @Author fzy
// @Date 2022-04-14 15:04:03
package xstring

func InArray(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
