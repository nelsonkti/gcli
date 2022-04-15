/**
** @创建时间 : 2022/3/23 09:25
** @作者 : fzy
 */
package string

import "strings"

// 去差集
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

//求交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// str 包含数组
func ContainArray(slice1 []string, substr string) bool {
	if len(slice1) == 0 || substr == "" {
		return false
	}
	for _, s := range slice1 {
		if strings.Contains(substr, s) {
			return true
		}
	}
	return false
}
