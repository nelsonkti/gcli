// Package cmd
// @Author fzy
// @Date 2022-04-14 15:04:03
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
	Short: "create a new repository file",
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
