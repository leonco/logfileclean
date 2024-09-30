package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type datePatternString struct {
	regexp     string
	dateFormat string
}

type datePattern struct {
	regexp     *regexp.Regexp
	dateFormat string
}

var datePatterns = []datePatternString{
	{"(?P<date>[0-9]{4}-[0-9]{2}-[0-9]{2})", "2006-01-02"},
	{"(?P<date>[0-9]{8})", "20060102"},
}

type stringParser interface {
	ParseString(filename string) (entry Entry, err error)
}

type Parser struct {
	Format       string
	DatePatterns []datePattern
}

func NewParser(format string) *Parser {
	preparedFormat := regexp.QuoteMeta(format)
	placeholder := "#date#"
	var reSlices []datePattern
	for _, r := range datePatterns {
		reSlices = append(reSlices, datePattern{
			regexp:     regexp.MustCompile(strings.Replace(preparedFormat, placeholder, r.regexp, -1)),
			dateFormat: r.dateFormat})
	}

	return &Parser{format, reSlices}
}

func (p *Parser) ParseString(filename string) (entry *Entry, err error) {
	for _, p := range p.DatePatterns {
		if m := p.regexp.FindStringSubmatch(filename); m != nil {
			for i, name := range p.regexp.SubexpNames() {
				if i != 0 && name == "date" {
					if t, err := time.Parse(p.dateFormat, m[i]); err == nil {
						return &Entry{filename, t}, nil
					}
				}
			}
		}
	}
	return nil, fmt.Errorf("No date found in %s", filename)
}
