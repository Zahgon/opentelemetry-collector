package confmap

import (
	"context"
	"errors"
	"regexp"
)

const schemePattern = `[A-Za-z][A-Za-z0-9+.-]+`

var (
	uriRegexp = regexp.MustCompile(`(?s:^(?P<Scheme>` + schemePattern + `):(?P<OpaqueValue>.*)$)`)

	errTooManyRecursiveExpansions = errors.New("too many recursive expansions")
)

func (mr *Resolver) expandValueRecursively(ctx context.Context, value any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (mr *Resolver) expandValue(ctx context.Context, value any) (any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

func (mr *Resolver) findURI(input string) string { _ = "STUB: not implemented"; return "" }

func (mr *Resolver) findAndExpandURI(ctx context.Context, input string) (any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

func (mr *Resolver) expandURI(ctx context.Context, input string) (*Retrieved, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type location struct {
	scheme      string
	opaqueValue string
}

func (c location) asString() string { _ = "STUB: not implemented"; return "" }

func newLocation(uri string) (location, error) {
	_ = "STUB: not implemented"
	return *new(location), nil
}
