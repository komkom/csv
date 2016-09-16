# csv display
command line tool to display ranges of rows in a csv encoded file.

## usage

go run main.go pathtocsv --match [1,10]

displays the header and row 1 to 10 of the file.

## give it a try

go get github.com/komkom/csvdisplay/...

this is basically a thin wrapper around github.com/olekukonko/tablewriter, so thanks.
