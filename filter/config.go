package filter

import (
	"regexp"
)

type Config struct {
	Strict string `mapstructure:"strict"`
	Regex  string `mapstructure:"regexp"`

	_ struct{}
}

func (c Config) Validate() error { _ = "STUB: not implemented"; return nil }

type combinedFilter struct {
	stricts map[any]struct{}
	regexes []*regexp.Regexp
}

func CreateFilter(configs []Config) Filter { _ = "STUB: not implemented"; return *new(Filter) }

func (cf *combinedFilter) Matches(toMatch any) bool { _ = "STUB: not implemented"; return false }
