package filters

//type Records [][]string
type Header []string
type FilterAction func(header Header, record []string) (trec *[]string)
