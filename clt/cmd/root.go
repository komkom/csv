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

var RootCmd = &cobra.Command{
	Use:   "csvdispaly",
	Short: "Print csv file",

	Run: func(cmd *cobra.Command, args []string) {

		var start int
		var end int

		if len(match) > 0 {
			match = strings.Replace(match, `[`, ``, -1)
			match = strings.Replace(match, `]`, ``, -1)
			ranges := strings.Split(match, `,`)
			if len(ranges) != 2 {
				panic(fmt.Errorf(`match format incorrect`))
			}

			rs := strings.Trim(ranges[0], ` `)
			re := strings.Trim(ranges[1], ` `)

			//fmt.Println(rs + ` ` + re)

			s, err := strconv.ParseInt(rs, 10, 64)
			if err != nil {
				panic(err)
			}
			e, err := strconv.ParseInt(re, 10, 64)
			if err != nil {
				panic(err)
			}

			start = int(s)
			end = int(e)
		}

		if len(args) != 1 {
			panic(fmt.Errorf(`expected csv arg as only argument.`))
		}

		err := display.Render(args[0], start, end)
		if err != nil {
			panic(err)
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
}
