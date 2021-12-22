/**
** @创建时间 : 2021/12/22 17:40
** @作者 : fzy
 */
package cmd

import (
	"errors"
	"fmt"
	"gcli/util/xfile"
	"gcli/util/xstring"
	"github.com/gobuffalo/packr/v2"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var name string
var zhName string
var defaultPath = "./"
var tmplPath = "./templates/make_templates/"
var tmplFileName = "tmp.go.tmpl"

func makeFile(args []string, createType string) error {
	if name == "" {
		return errors.New("请输入服务名称, 查看详情 -help")
	}

	file := strings.Replace(args[0], "\\", "/", -1)

	fileName := getFileName(file)
	path := getPath(file)
	packName := getPackName(path)

	fmt.Println(packName)

	box := packr.New(tmplPath, tmplPath)
	tmpl, _ := box.FindString(tmplFileName)

	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	fileName = addSuffix(fileName, fmt.Sprintf("%s%s", "_", createType))

	hostname, _ := os.Hostname()
	tmplData := map[string]interface{}{
		"Name":       fmt.Sprintf("%s %s", name, zhName),
		"CreateTime": time.Now().Format("2006-01-02 15:04:05"),
		"Author":     hostname,
		"PackName":   packName,
		"StructName": xstring.Case2Camel(fileName),
	}

	if err := xfile.WriteFile(filepath.Join(path, addSuffix(fileName, ".go")), tmpl, tmplData); err != nil {
		return err
	}

	return nil
}

func getFileName(file string) string {
	var fileName string
	i := strings.LastIndex(file, string(os.PathSeparator))

	if xstring.IsUpper(file) {
		fileName = xstring.Camel2Case(file[i+1:])
	}

	return fileName
}

func getPath(file string) string {
	i := strings.LastIndex(file, string(os.PathSeparator))

	if i > 0 {
		return strings.ToLower(file[:i])
	}

	return defaultPath
}

func getPackName(path string) string {
	i := strings.LastIndex(path, string(os.PathSeparator))

	if i > 0 && path != defaultPath {
		return path[i+1:]
	}

	return "main"
}

func addSuffix(name string, suffix string) string {

	if !strings.Contains(name, suffix) {
		return name + suffix
	}

	return name
}
