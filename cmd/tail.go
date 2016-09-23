package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Last n rows form a file eg. csv tail 10",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			panic(fmt.Errorf("no tail row count specified"))
		}

		tail, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}

		count := countAction()

		start := count - int(tail)
		if start < 0 {
			start = 0
		}

		run2(nil, start, nil)
	},
}

func init() {
	RootCmd.AddCommand(tailCmd)
	tailCmd.Flags().StringVar(&csvPath, "path", "", "path to a csv encoded file.")
	tailCmd.Flags().BoolVar(&print, "print", false, "prints the csv file as a table.")
}
