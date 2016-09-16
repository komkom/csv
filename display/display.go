package display

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"unicode"

	"github.com/olekukonko/tablewriter"
)

func Render(csvPath string, start int, end int) (err error) {
	header, records, err := createData(csvPath, start, end)
	if err != nil {
		return
	}

	table, err := createTable(records, header)
	if err != nil {
		return
	}

	table.Render()
	return
}

func createData(csvPath string, start int, end int) (header []string, records [][]string, err error) {

	if start > end {
		err = fmt.Errorf(`records start > end`)
		return
	}

	f, err := os.Open(csvPath)
	if err != nil {
		return
	}

	r := csv.NewReader(f)
	idx := 0
	for {
		if end <= idx {
			break
		}

		record, rerr := r.Read()
		if rerr == io.EOF {
			break
		}
		if rerr != nil {
			err = rerr
			return
		}

		if idx == 0 {
			// parse the header
			h, consume := createHeader(record)
			header = h
			if consume {
				continue
			}
		}

		if idx >= start {

			if len(record) != len(header) {
				err = fmt.Errorf("len(record) != len(header)")
			}

			records = append(records, record)
		}

		idx++
	}

	return
}

func createTable(records [][]string, header []string) (table *tablewriter.Table, err error) {

	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetBorder(false)

	for _, r := range records {
		if len(r) != len(header) {
			err = fmt.Errorf("len(record) != len(header)")
			return
		}

		table.Append(r)
	}
	return
}

func createHeader(record []string) (header []string, consume bool) {
	consume = true
	for idx, c := range record {
		if possibleHeaderValue(c) {
			header = append(header, c)
		} else {
			consume = false
			header = append(header, fmt.Sprintf("c%v", idx))
		}
	}

	return
}

func possibleHeaderValue(value string) bool {

	for _, c := range value {
		if !unicode.IsLower(c) && c != '_' {
			return false
		}
	}

	return true
}
