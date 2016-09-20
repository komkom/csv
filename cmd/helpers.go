package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/komkom/csv/render"
	"github.com/spf13/cobra"
)

var crange string
var print bool
var start int
var end *int
var csvPath string

func setup(cmd *cobra.Command) {
	cmd.Flags().StringVar(&csvPath, "path", "", "path to a csv encoded file.")
	cmd.Flags().StringVar(&crange, "range", "", "rows to match eg `[0,10]` form row 0 to 9")
	cmd.Flags().BoolVar(&print, "print", false, "prints the csv file as a table.")
}

func run(filter render.Filter) {

	if len(crange) > 0 {
		crange = strings.Replace(crange, `[`, ``, -1)
		crange = strings.Replace(crange, `]`, ``, -1)
		ranges := strings.Split(crange, `,`)
		if len(ranges) != 2 {
			panic(fmt.Errorf(`match format incorrect`))
		}

		rs := strings.Trim(ranges[0], ` `)
		re := strings.Trim(ranges[1], ` `)

		s, err := strconv.ParseInt(rs, 10, 64)
		if err != nil {
			panic(err)
		}
		e, err := strconv.ParseInt(re, 10, 64)
		if err != nil {
			panic(err)
		}

		start = int(s)
		t := int(e)
		end = &t
	}

	input := os.Stdin
	if len(csvPath) > 0 {
		f, err := os.Open(csvPath)
		if err != nil {
			panic(err)
		}
		input = f
	}

	var output render.Output
	output = render.NewCsvOutput()
	if print {
		output = render.NewTableOutput()
	}

	err := render.StartReadingCSV(input, filter, output, start, end)
	if err != nil {
		panic(err)
	}
}
