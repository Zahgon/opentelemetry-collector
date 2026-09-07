package pprofile

import (
	"errors"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

var errTooManyStringTableEntries = errors.New("too many entries in StringTable")

func SetString(table pcommon.StringSlice, val string) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
