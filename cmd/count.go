package cmd

import (
	"fmt"
	"os"

	"github.com/komkom/csv/render"
	"github.com/spf13/cobra"
)

type CountFilter struct {
	count int
}

func (c CountFilter) Count() int {
	return c.count
}

func (c *CountFilter) Header(header []string) (headerout []string, err error) {
	return nil, nil
}

func (c *CountFilter) Record(record []string) (recordout []string, err error) {
	c.count += 1
	return nil, fmt.Errorf(`dont include this record`)
}

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count the number of rows in a csv file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v\n", countAction())
	},
}

func countAction() int {
	countFilter := &CountFilter{}

	input := os.Stdin
	if len(csvPath) > 0 {
		f, err := os.Open(csvPath)
		if err != nil {
			panic(err)
		}
		input = f
	}

	err := render.StartReadingCSV(input, countFilter, &render.NilOutput{}, start, end)
	if err != nil {
		panic(err)
	}

	return countFilter.Count()
}

func init() {
	RootCmd.AddCommand(countCmd)
	countCmd.Flags().StringVar(&csvPath, "path", "", "path to a csv encoded file.")
}
