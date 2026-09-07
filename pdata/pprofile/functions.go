package pprofile

import (
	"errors"
)

var errTooManyFunctionTableEntries = errors.New("too many entries in FunctionTable")

func SetFunction(table FunctionSlice, fn Function) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
