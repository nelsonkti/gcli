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

	Cmd.AddCommand(modelCommand)

	modelCommand.PersistentFlags().StringVar(&name, "n", "", "请输入 [模型] 中文名称")

}

var modelCommand = &cobra.Command{
	Use:   "make:model",
	Short: "create a new model file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		zhName = "模型"

		err := makeFile(args, "model")

		if err != nil {
			fmt.Println(xprintf.Red(err.Error()))
			return
		}
	},
}
