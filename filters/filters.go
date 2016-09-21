package filters

import (
	"fmt"
	"regexp"
)

type IndexFilter struct {
	count int
}

func (f *IndexFilter) Header(header []string) (headerout []string, err error) {
	return append([]string{`idx`}, header...), nil
}

func (f *IndexFilter) Record(record []string) (recordout []string, err error) {
	record = append(record, ``)
	copy(record[1:len(record)], record[0:len(record)-1])
	record[0] = fmt.Sprintf("%v", f.count)
	f.count += 1
	return record, nil
}

type MatchFilter struct {
	column  *int
	regexes []*regexp.Regexp
	invert  bool
}

func NewMatchFilter(rregexes []string, column *int, invert bool) *MatchFilter {

	var regexes []*regexp.Regexp
	for _, rg := range rregexes {
		regexes = append(regexes, regexp.MustCompile(rg))
	}

	return &MatchFilter{column, regexes, invert}
}

func (m *MatchFilter) Header(header []string) (headerout []string, err error) {
	return header, nil
}

func (m *MatchFilter) Record(record []string) (recordout []string, err error) {

	if m.column != nil {
		c := *m.column
		if c >= len(record) {
			return nil, fmt.Errorf("no such column")
		}

		for _, r := range m.regexes {
			if r.FindAllString(record[c], 1) != nil {
				if !m.invert {
					return record, nil
				} else {
					return nil, fmt.Errorf("has match")
				}
			}
		}
	} else {

		for _, c := range record {
			for _, r := range m.regexes {
				if r.FindAllString(c, 1) != nil {
					if !m.invert {
						return record, nil
					} else {
						return nil, fmt.Errorf("has match")
					}
				}
			}
		}
	}

	if !m.invert {
		return record, fmt.Errorf("not matched")
	} else {
		return record, nil
	}
}
