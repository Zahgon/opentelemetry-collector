package pprofile

import (
	"errors"
)

func FromLocationIndices(table LocationSlice, record Stack) (LocationSlice, error) {
	_ = "STUB: not implemented"
	return *new(LocationSlice), nil
}

var errTooManyLocationTableEntries = errors.New("too many entries in LocationTable")

func SetLocation(table LocationSlice, loc Location) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
