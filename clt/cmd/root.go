package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/komkom/csvdisplay/display"
	"github.com/spf13/cobra"
)

var cfgFile string
var match string
var print bool

var RootCmd = &cobra.Command{
	Use:   "csvdispaly",
	Short: "Print csv file",

	Run: func(cmd *cobra.Command, args []string) {

		var start int
		var end *int

		if len(match) > 0 {
			match = strings.Replace(match, `[`, ``, -1)
			match = strings.Replace(match, `]`, ``, -1)
			ranges := strings.Split(match, `,`)
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
		if len(args) > 0 {
			f, err := os.Open(args[0])
			if err != nil {
				panic(err)
			}
			input = f
		}

		if print {
			err := display.StartReadingCSV(input, nil, display.NewTableOutput(), start, end)
			if err != nil {
				panic(err)
			}
		} else {
			err := display.StartReadingCSV(input, nil, display.NewCsvOutput(), start, end)
			if err != nil {
				panic(err)
			}
		}
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.Flags().StringVar(&match, "match", "", "entries to match eg `[10,10]`")
	RootCmd.Flags().BoolVar(&print, "print", false, "prints the csv file as a table.")
}
