package cmd

import (
	"fmt"
	"gcli/util/xprintf"
	"github.com/spf13/cobra"
)

const Version = "1.0"


func init() {
	Cmd.AddCommand(version)
}

var version = &cobra.Command{
	Use:   "version",
	Short: "print gcli version",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(xprintf.Blue("gcli version "+Version))
	},
}