package pprofile

import (
	"errors"
)

var errTooManyMappingTableEntries = errors.New("too many entries in MappingTable")

func SetMapping(table MappingSlice, ma Mapping) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
