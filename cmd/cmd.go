package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "gcli",
	Short: "gcli is a scaffolding tool for project development.",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
