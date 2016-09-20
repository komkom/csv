# csv command line tool
view create and transform csv files.

## usage

csv --path somefile.csv --print
display a pretty csv file on the screen

```
  AD ID | AD SUBJECT |            AD BODY             | AD PRICE
+-------+------------+--------------------------------+----------+
  12345 | subject 0  | body asdf sdfdfs               | price 0
        |            | asdfdsfasdfasdf   asdfasdfasdf |
        |            |   asdfasdfsdf  asdfasdfsadfasd |
        |            |   asdl;kfjsdf  sdfjaslk        |
  12345 | subject 1  | body asdf sdfdfs               | price 1
        |            | asdfdsfasdfasdf   asdfasdfasdf |
        |            |   asdfasdfsdf  asdfasdfsadfasd |
        |            |   asdl;kfjsdf  sdfjaslk        |
```

csv --path somefile.csv --range [1,10]
displays the header and row 1 to 9 of the file.

..add --print so that you can actually see something
csv --path somefile.csv --range [1,10] --print 

use the pipe
csv --path | csv index | csv match "query" -- print

this adds an index as the first column then filters out all the columns which do not contain `query`. 







csv index --path somefile.csv
add an index to 




## give it a try

go get github.com/komkom/csvdisplay/...
