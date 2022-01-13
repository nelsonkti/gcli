/**
** @创建时间 : 2021/12/22 10:39
** @作者 : fzy
 */
package cmd

import (
	"errors"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/nelsonkti/gcli/lib"
	"github.com/nelsonkti/gcli/util/xprintf"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	sourcePath  string
	destPath    string
	replaceName string
)

func init() {

	Cmd.AddCommand(templateCommand)

	templateCommand.PersistentFlags().StringVar(&sourcePath, "s", "$GOPATH", "source file path")
	templateCommand.PersistentFlags().StringVar(&destPath, "d", "$GOPATH", "create file path")
	templateCommand.PersistentFlags().StringVar(&replaceName, "r", "", "replace name")

}

var templateCommand = &cobra.Command{
	Use:   "template",
	Short: "template service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := conversionTemplate()

		if err != nil {
			fmt.Println(xprintf.Red(err.Error()))
			return
		}
	},
}

func conversionTemplate() error {
	if lib.Version() < "1.13" {
		return errors.New("the current version is relatively low")
	}

	_, err := ioutil.ReadDir(sourcePath)
	if err != nil {
		return err
	}

	box := packr.New(sourcePath, sourcePath)

	_, err = ioutil.ReadDir(destPath)
	if err != nil {
		if err = os.MkdirAll(destPath, 0755); err != nil {
			return err
		}
	}

	newAbsPath := "/" + box.ResolutionDir + "/"

	for _, name := range box.List() {

		if strings.HasPrefix(name, ".git/") || strings.HasPrefix(name, ".idea/") ||
			strings.HasPrefix(name, "log/") || name == "config.yaml" || name == "config.json" {
			continue
		}

		tmpl, _ := box.FindString(newAbsPath + name)

		i := strings.LastIndex(name, string(os.PathSeparator))

		if i > 0 {
			dir := name[:i]
			if err = os.MkdirAll(filepath.Join(destPath, dir), 0755); err != nil {
				return err
			}
		}

		tmpl = PareTmpl(name, tmpl, replaceName)

		name = fmt.Sprintf("%s.%s", name, "tmpl")

		if strings.HasSuffix(name, ".example.tmpl") {
			if err = WriteFile(fmt.Sprintf("%s/%s", destPath, strings.Replace(name, ".example", "", 1)), tmpl); err != nil {
				return err
			}
		}

		if err = WriteFile(fmt.Sprintf("%s/%s", destPath, name), tmpl); err != nil {
			return err
		}

	}

	return nil
}

func WriteFile(path, tmpl string) (err error) {
	return ioutil.WriteFile(path, []byte(tmpl), 0755)
}

func PareTmpl(name, tmpl, replaceName string) string {
	switch name {
	case "README.md":
		break
	case "go.mod":
		tmpl = pareTmpl(tmpl, replaceName, "{{.Name}}")
	default:
		tmpl = pareTmpl(tmpl, replaceName, "{{.ShortPath}}{{.Name}}")
	}

	return tmpl
}

func pareTmpl(tmpl, replaceName string, format string) string {
	if replaceName == "" {
		return tmpl
	}
	return strings.Replace(tmpl, replaceName, format, -1)
}

func exaClone() {

}
