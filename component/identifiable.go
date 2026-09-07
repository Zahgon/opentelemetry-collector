package component

import (
	"encoding"
	"fmt"
	"regexp"
)

const typeAndNameSeparator = "/"

var (
	typeRegexp = regexp.MustCompile(`^[a-zA-Z][0-9a-zA-Z_]{0,62}$`)

	nameRegexp = regexp.MustCompile(`^[^\pZ\pC\pS]+$`)
)

var _ fmt.Stringer = Type{}

type Type struct {
	name string
}

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

func (t Type) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewType(ty string) (Type, error) { _ = "STUB: not implemented"; return *new(Type), nil }

func MustNewType(strType string) Type { _ = "STUB: not implemented"; return *new(Type) }

var (
	_ fmt.Stringer             = ID{}
	_ encoding.TextMarshaler   = ID{}
	_ encoding.TextUnmarshaler = (*ID)(nil)
)

type ID struct {
	typeVal Type   `mapstructure:"-"`
	nameVal string `mapstructure:"-"`
}

func NewID(typeVal Type) ID { _ = "STUB: not implemented"; return *new(ID) }

func MustNewID(typeVal string) ID { _ = "STUB: not implemented"; return *new(ID) }

func NewIDWithName(typeVal Type, nameVal string) ID { _ = "STUB: not implemented"; return *new(ID) }

func MustNewIDWithName(typeVal, nameVal string) ID { _ = "STUB: not implemented"; return *new(ID) }

func (id ID) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (id ID) Name() string { _ = "STUB: not implemented"; return "" }

func (id ID) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id *ID) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (id ID) String() string { _ = "STUB: not implemented"; return "" }

func validateName(nameStr string) error { _ = "STUB: not implemented"; return nil }
