package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type Location struct {
	orig  *internal.Location
	state *internal.State
}

func newLocation(orig *internal.Location, state *internal.State) Location {
	_ = "STUB: not implemented"
	return *new(Location)
}

func NewLocation() Location { _ = "STUB: not implemented"; return *new(Location) }

func (ms Location) MoveTo(dest Location) { _ = "STUB: not implemented"; return }

func (ms Location) MappingIndex() int32 { _ = "STUB: not implemented"; return 0 }

func (ms Location) SetMappingIndex(v int32) { _ = "STUB: not implemented"; return }

func (ms Location) Address() uint64 { _ = "STUB: not implemented"; return 0 }

func (ms Location) SetAddress(v uint64) { _ = "STUB: not implemented"; return }

func (ms Location) Lines() LineSlice { _ = "STUB: not implemented"; return *new(LineSlice) }

func (ms Location) AttributeIndices() pcommon.Int32Slice {
	_ = "STUB: not implemented"
	return *new(pcommon.Int32Slice)
}

func (ms Location) CopyTo(dest Location) { _ = "STUB: not implemented"; return }
