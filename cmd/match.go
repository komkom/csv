package cmd

import (
	"fmt"

	"github.com/komkom/csv/filters"
	"github.com/spf13/cobra"
)

// matchCmd represents the match command
var matchCmd = &cobra.Command{
	Use:   "match",
	Short: "A filter by match eg. `clt match test1` filters all the rows which do not contain test1 in any column",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic(fmt.Errorf("no match arg found."))
		}

		mf := filters.NewMatchFilter(args)
		run(mf)
	},
}

func init() {
	RootCmd.AddCommand(matchCmd)
	setup(matchCmd)
}
