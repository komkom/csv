package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "csv",
	Short: "Print csv file",

	Run: func(cmd *cobra.Command, args []string) {
		run(nil)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	setup(RootCmd)
}
