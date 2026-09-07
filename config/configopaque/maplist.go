package configopaque

import (
	"iter"

	"go.opentelemetry.io/collector/confmap"
)

type Pair struct {
	Name  string `mapstructure:"name"`
	Value String `mapstructure:"value"`

	_ struct{}
}

type MapList []Pair

var _ confmap.Unmarshaler = (*MapList)(nil)

func (ml *MapList) Unmarshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

var _ confmap.Validator = MapList(nil)

func (ml MapList) Validate() error { _ = "STUB: not implemented"; return nil }

var _ iter.Seq2[string, String] = MapList(nil).Iter

func (ml MapList) Iter(yield func(name string, value String) bool) {
	_ = "STUB: not implemented"
	return
}

func (ml MapList) Get(name string) (val String, ok bool) {
	_ = "STUB: not implemented"
	return *new(String), false
}

func (ml *MapList) Set(name string, val String) { _ = "STUB: not implemented"; return }
