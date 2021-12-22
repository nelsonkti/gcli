/**
** @创建时间 : 2021/12/22 15:52
** @作者 : fzy
 */
package cmd

import (
	"fmt"
	"github.com/nelsonkti/gcli/util/xprintf"
	"github.com/spf13/cobra"
)

func init() {

	Cmd.AddCommand(repositoryCommand)

	repositoryCommand.PersistentFlags().StringVar(&name, "n", "", "请输入 [仓库] 中文名称")

}

var repositoryCommand = &cobra.Command{
	Use:   "make:repository",
	Short: "create a repository file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		zhName = "仓库"
		err := makeFile(args, "repository")

		if err != nil {
			fmt.Println(xprintf.Red(err.Error()))
			return
		}
	},
}
