/**
** @创建时间 : 2021/12/22 10:39
** @作者 : fzy
 */
package cmd

import (
	"errors"
	"fmt"
	"github.com/nelsonkti/gcli/lib"
	"github.com/nelsonkti/gcli/util/xfile"
	"github.com/nelsonkti/gcli/util/xprintf"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	sourcePath string
	destPath   string
)

func init() {

	Cmd.AddCommand(templateCommand)

	templateCommand.PersistentFlags().StringVar(&sourcePath, "s", "$GOPATH", "source file path")
	templateCommand.PersistentFlags().StringVar(&destPath, "d", "$GOPATH", "create file path")

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

	var projectName string
	destPathSeparator := strings.LastIndex(destPath, string(os.PathSeparator))
	if destPathSeparator > 0 {
		projectName = destPath[destPathSeparator+1:]
	}

	_, err = ioutil.ReadDir(destPath)
	if err != nil {
		if err = os.MkdirAll(destPath, 0755); err != nil {
			return err
		}
	}

	for _, name := range box.List() {

		if strings.HasPrefix(name, ".git/") || strings.HasPrefix(name, ".idea/") ||
			strings.HasPrefix(name, "log/") || strings.HasPrefix(name, "config.json") {
			continue
		}

		tmpl, _ := box.FindString(name)

		i := strings.LastIndex(name, string(os.PathSeparator))

		if i > 0 {
			dir := name[:i]
			if err = os.MkdirAll(filepath.Join(destPath, dir), 0755); err != nil {
				return err
			}
		}

		name = fmt.Sprintf("%s.%s", name, "tmpl")

		if err = xfile.WriteFile(filepath.Join(destPath, name), tmpl, projectName); err != nil {
			return err
		}

	}

	return nil
}
