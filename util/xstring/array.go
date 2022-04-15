/**
** @创建时间 : 2022/4/15 09:05
** @作者 : fzy
 */
package xstring

func InArray(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
