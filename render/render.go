package render

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
	"unicode"

	"github.com/olekukonko/tablewriter"
)

type Output interface {
	AddHeader(header []string)
	AddRecord(record []string)
	Flush()
}

type Filter interface {
	Header(header []string) (headerout []string, err error)
	Record(idx int, record []string) (recordout []string, err error)
}

type NilOutput struct{}

func (n *NilOutput) AddHeader(header []string) {}
func (n *NilOutput) AddRecord(record []string) {}
func (n *NilOutput) Flush()                    {}

type NilFilter struct {
	count int
}

type IndexedRecord struct {
	idx    int
	record []string
}

func (f *NilFilter) Header(header []string) (headerout []string, err error) {
	return header, nil
}

func (f *NilFilter) Record(idx int, record []string) (recordout []string, err error) {
	return record, nil
}

type TableOutput struct {
	table       *tablewriter.Table
	insertCount int
}

func NewTableOutput() *TableOutput {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)

	return &TableOutput{table: table}
}

func (t *TableOutput) AddHeader(header []string) {
	t.table.SetHeader(header)
}

func (t *TableOutput) AddRecord(record []string) {
	t.table.Append(record)

	if t.insertCount >= 100 {
		t.table.Render()
		t.insertCount = 0
	}

	t.insertCount += 1
}

func (t TableOutput) Flush() {
	t.table.Render()
}

type CsvOutput struct {
	writer      *csv.Writer
	insertCount int
}

func NewCsvOutput() *CsvOutput {
	writer := csv.NewWriter(os.Stdout)
	return &CsvOutput{writer: writer}
}

func (c *CsvOutput) AddHeader(header []string) {
	c.writer.Write(header)
}

func (c *CsvOutput) AddRecord(record []string) {
	c.writer.Write(record)

	if c.insertCount >= 100 {
		c.writer.Flush()
		c.insertCount = 0
	}

	c.insertCount += 1
}

func (c CsvOutput) Flush() {
	c.writer.Flush()
}

func StartReadingCSV(reader io.Reader, filter Filter, output Output, start int, end *int) error {

	if end != nil && start > (*end) {
		return fmt.Errorf(`records start > end`)
	}

	if filter == nil {
		filter = &NilFilter{}
	}

	done := make(chan int)
	msgs := make(chan interface{})
	go func() {
		readCSV(reader, start, end, msgs)
		close(done)
	}()

	var hasHeader bool
	var hasRun bool
	var idx int
	for {

		select {
		case <-time.After(time.Second):
			if !hasRun {
				return fmt.Errorf("no input")
			}

		case m := <-msgs:
			switch o := m.(type) {
			case []string:
				if !hasHeader {

					h, err := filter.Header(o)
					if err != nil {
						return err
					}

					output.AddHeader(h)
					hasHeader = true
					continue
				}
			case IndexedRecord:
				if !hasHeader {
					return fmt.Errorf(`no header found.`)
				}

				r, err := filter.Record(o.idx, o.record)
				if err == nil {
					output.AddRecord(r)
				}

				idx += 1

			case error:
				return o
			}
		case <-done:
			goto exit

		}

		hasRun = true
	}
exit:

	output.Flush()

	return nil
}

func readCSV(reader io.Reader, start int, end *int, msgs chan interface{}) {
	r := csv.NewReader(reader)
	idx := 0
	var header []string
	hasHeader := false
	for {
		if end != nil && *end <= idx {
			break
		}

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			msgs <- err
			return
		}

		if !hasHeader {
			hasHeader = true
			// parse the header
			h, consume := createHeader(record)
			header = h

			msgs <- h
			if consume {
				continue
			}
		}

		if idx >= start {

			if len(record) != len(header) {
				msgs <- fmt.Errorf("len(record) != len(header)")
				return
			}

			msgs <- IndexedRecord{idx, record}
			//records = append(records, record)
		}

		idx++
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
