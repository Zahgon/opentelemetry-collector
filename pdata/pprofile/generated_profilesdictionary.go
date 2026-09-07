package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ProfilesDictionary struct {
	orig  *internal.ProfilesDictionary
	state *internal.State
}

func newProfilesDictionary(orig *internal.ProfilesDictionary, state *internal.State) ProfilesDictionary {
	_ = "STUB: not implemented"
	return *new(ProfilesDictionary)
}

func NewProfilesDictionary() ProfilesDictionary {
	_ = "STUB: not implemented"
	return *new(ProfilesDictionary)
}

func (ms ProfilesDictionary) MoveTo(dest ProfilesDictionary) { _ = "STUB: not implemented"; return }

func (ms ProfilesDictionary) MappingTable() MappingSlice {
	_ = "STUB: not implemented"
	return *new(MappingSlice)
}

func (ms ProfilesDictionary) LocationTable() LocationSlice {
	_ = "STUB: not implemented"
	return *new(LocationSlice)
}

func (ms ProfilesDictionary) FunctionTable() FunctionSlice {
	_ = "STUB: not implemented"
	return *new(FunctionSlice)
}

func (ms ProfilesDictionary) LinkTable() LinkSlice {
	_ = "STUB: not implemented"
	return *new(LinkSlice)
}

func (ms ProfilesDictionary) StringTable() pcommon.StringSlice {
	_ = "STUB: not implemented"
	return *new(pcommon.StringSlice)
}

func (ms ProfilesDictionary) AttributeTable() KeyValueAndUnitSlice {
	_ = "STUB: not implemented"
	return *new(KeyValueAndUnitSlice)
}

func (ms ProfilesDictionary) StackTable() StackSlice {
	_ = "STUB: not implemented"
	return *new(StackSlice)
}

func (ms ProfilesDictionary) CopyTo(dest ProfilesDictionary) { _ = "STUB: not implemented"; return }
