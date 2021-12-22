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

	Cmd.AddCommand(serviceCommand)

	serviceCommand.PersistentFlags().StringVar(&name, "n", "", "请输入 [服务] 中文名称")
}

var serviceCommand = &cobra.Command{
	Use:   "make:service",
	Short: "create a service file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		zhName = "服务"
		err := makeFile(args, "service")

		if err != nil {
			fmt.Println(xprintf.Red(err.Error()))
			return
		}
	},
}
