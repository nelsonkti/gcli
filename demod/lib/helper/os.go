/**
** @创建时间 : 2022/3/30 10:28
** @作者 : fzy
 */
package helper

import (
	"path"
	"runtime"
	"strings"
)

// 当前文件名称
func CurrentFileName() string {
	_, fullFilename, _, _ := runtime.Caller(1)
	return strings.Replace(path.Base(fullFilename), ".go", "", -1)
}
