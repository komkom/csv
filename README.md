# csv command line tool
view create and transform csv files.

## usage

```
csv --print --path somefile.csv
```
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
```
csv --path somefile.csv --range [1,10]
```
displays the header and row 1 to 9 of the file.

..add --print so that you can actually see something
```
csv --path somefile.csv --range [1,10] --print 
```

use the pipe
```
csv --path file.csv | csv index | csv match "subject" -- print
```

this adds an index as the first column then filters out all the columns which do not contain `subject`. 
somethink like 
```
IDX | AD ID | AD SUBJECT |            AD BODY             | AD PRICE
+-----+-------+------------+--------------------------------+----------+
   99 | 12345 | subject 99 | body asdf query                | price 99
      |       |            | asdfdsfasdfasdf   asdfasdfasdf |
      |       |            |   asdfasdfsdf  asdfasdfsadfasd |
      |       |            |   asdl;kfjsdf  sdfjaslk        |
```


## give it a try

go get github.com/komkom/csv/...

would be nice to get some more useful filters if anyone finds the time...
