package pipeline

import (
	"regexp"
)

const typeAndNameSeparator = "/"

type ID struct {
	signal Signal `mapstructure:"-"`
	name   string `mapstructure:"-"`
}

func NewID(signal Signal) ID { _ = "STUB: not implemented"; return *new(ID) }

func NewIDWithName(signal Signal, name string) ID { _ = "STUB: not implemented"; return *new(ID) }

func (i ID) Signal() Signal { _ = "STUB: not implemented"; return *new(Signal) }

func (i ID) Name() string { _ = "STUB: not implemented"; return "" }

func (i ID) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (i *ID) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (i ID) String() string { _ = "STUB: not implemented"; return "" }

var nameRegexp = regexp.MustCompile(`^[^\pZ\pC\pS]+$`)

func validateName(nameStr string) error { _ = "STUB: not implemented"; return nil }
