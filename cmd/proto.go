package cmd

import (
	"fmt"
	"github.com/nelsonkti/gcli/util/xprintf"
	"github.com/spf13/cobra"
)

func init() {

	Cmd.AddCommand(protoCommand)

	protoCommand.PersistentFlags().StringVar(&name, "n", "", "请输入 [protobuf] 中文名称")

}

var protoCommand = &cobra.Command{
	Use:   "make:proto",
	Short: "create a new protobuf file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		zhName = "protobuf"
		err := makeFile(args, "proto")

		if err != nil {
			fmt.Println(xprintf.Red(err.Error()))
			return
		}
	},
}
