package cmd

import (
	"fmt"
	"strconv"

	"github.com/komkom/csv/filters"
	"github.com/spf13/cobra"
)

var column string
var invert bool

// matchCmd represents the match command
var matchCmd = &cobra.Command{
	Use:   "match",
	Short: "A filter by match eg. `clt match test1` filters all the rows which do not contain test1 in any column",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic(fmt.Errorf("no match arg found."))
		}

		var c *int

		if len(column) > 0 {
			ri, err := strconv.ParseInt(column, 10, 64)
			if err == nil {
				i := int(ri)
				c = &i
			}
		}

		mf := filters.NewMatchFilter(args, c, invert)
		run(mf)
	},
}

func init() {
	RootCmd.AddCommand(matchCmd)
	setup(matchCmd)
	matchCmd.Flags().StringVar(&column, "column", "", "column to match on.")
	matchCmd.Flags().BoolVar(&invert, "not", false, "filter out matches")
}
