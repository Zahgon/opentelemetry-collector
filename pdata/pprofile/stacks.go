package pprofile

import (
	"errors"
)

var errTooManyStackTableEntries = errors.New("too many entries in StackTable")

func SetStack(table StackSlice, st Stack) (int32, error) { _ = "STUB: not implemented"; return 0, nil }
