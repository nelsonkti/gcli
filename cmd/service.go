// Package cmd
// @Author fzy
// @Date 2022-04-14 15:04:03
package cmd

import (
	"fmt"
	"gcli/util/xprintf"
	"github.com/spf13/cobra"
)

func init() {

	Cmd.AddCommand(serviceCommand)

	serviceCommand.PersistentFlags().StringVar(&name, "n", "", "请输入 [服务] 中文名称")
}

var serviceCommand = &cobra.Command{
	Use:   "make:service",
	Short: "create a new service file",
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
