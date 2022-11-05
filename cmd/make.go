// Package cmd
// @Author fzy
// @Date 2022-04-14 15:04:03
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
var bakType = "main"
var typeNameMap = map[string]string{
	"model":      "模型",
	"repository": "仓库",
	"service":    "服务",
}

func makeFile(args []string, createType string) error {
	typeName := typeNameMap[createType]
	if typeName == "" {
		typeName = "服务"
	}

	if len(args) == 0 {
		return errors.New(fmt.Sprintf("请输入创建%s名称, 查看详情 -help", typeName))
	}

	if name == "" {
		return errors.New(fmt.Sprintf("请输入%s中文名称, 查看详情 -help", typeName))
	}

	bakType = createType

	file := strings.Replace(args[0], "\\", "/", -1)
	fileName := getFileName(file)
	path := getPath(file)
	packageName := getPackageName(path)

	box := packr.New(tmplPath, tmplPath)
	tmpl, _ := box.FindString(tmplFileName)

	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	fileName = addSuffix(fileName, fmt.Sprintf("%s%s", "_", createType))

	hostname, _ := os.Hostname()
	tmplData := map[string]interface{}{
		"Name":        fmt.Sprintf("%s %s", name, zhName),
		"CreateTime":  time.Now().Format("2006-01-02 15:04:05"),
		"Author":      hostname,
		"PackageName": packageName,
		"StructName":  xstring.Case2Camel(fileName),
	}

	if err := xfile.WriteFile(filepath.Join(path, addSuffix(fileName, ".go")), tmpl, tmplData); err != nil {
		return err
	}

	return nil
}

func getFileName(file string) string {
	var fileName string
	i := strings.LastIndex(file, string(os.PathSeparator))

	fileName = file
	if i > -1 {
		fileName = file[i+1:]
	}

	if xstring.IsUpper(fileName) {
		fileName = xstring.Camel2Case(fileName)
	}

	return fileName
}

func getPath(file string) string {
	i := strings.LastIndex(file, string(os.PathSeparator))

	if i > 0 {
		file := file[:i+1]

		if xstring.IsUpper(file) {
			file = xstring.Camel2Case(file)
		}
		if file[0:1] != string(os.PathSeparator) && file[0:1] != "." {
			file = fmt.Sprintf(".%s%s", string(os.PathSeparator), file)
		}
		if strings.Contains(file, "//") {
			file = strings.Replace(file, "//", "/", -1)
		}
		return file
	}

	return defaultPath
}

func getPackageName(path string) string {
	i := strings.LastIndex(path, string(os.PathSeparator))

	if i > 0 && path != defaultPath {
		packageName := path[:i]
		if strings.Contains(packageName, string(os.PathSeparator)) {
			j := strings.LastIndex(packageName, string(os.PathSeparator))
			return packageName[j+1 : i]
		}
		return packageName
	}

	return bakType
}

func addSuffix(name string, suffix string) string {

	if !strings.Contains(name, suffix) {
		return name + suffix
	}

	return name
}
