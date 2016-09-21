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
	regexes []*regexp.Regexp
}

func NewMatchFilter(rregexes []string) *MatchFilter {

	var regexes []*regexp.Regexp
	for _, rg := range rregexes {
		regexes = append(regexes, regexp.MustCompile(rg))
	}

	return &MatchFilter{regexes}
}

func (m *MatchFilter) Header(header []string) (headerout []string, err error) {
	return header, nil
}

func (m *MatchFilter) Record(record []string) (recordout []string, err error) {

	for _, c := range record {
		for _, r := range m.regexes {
			if r.FindAllString(c, 1) != nil {
				return record, nil
			}
		}
	}

	return record, fmt.Errorf("not matched")
}
