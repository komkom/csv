package cmd

import (
	"github.com/komkom/csvdisplay/filters"
	"github.com/spf13/cobra"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Add an index column as first column",
	Run: func(cmd *cobra.Command, args []string) {

		idxFilter := &filters.IndexFilter{}
		run(idxFilter)
	},
}

func init() {
	RootCmd.AddCommand(indexCmd)
	setup(indexCmd)
}
