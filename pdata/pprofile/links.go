package pprofile

import (
	"errors"
)

var errTooManyLinkTableEntries = errors.New("too many entries in LinkTable")

func SetLink(table LinkSlice, li Link) (int32, error) { _ = "STUB: not implemented"; return 0, nil }
