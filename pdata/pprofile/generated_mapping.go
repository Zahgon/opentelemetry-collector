package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type Mapping struct {
	orig  *internal.Mapping
	state *internal.State
}

func newMapping(orig *internal.Mapping, state *internal.State) Mapping {
	_ = "STUB: not implemented"
	return *new(Mapping)
}

func NewMapping() Mapping { _ = "STUB: not implemented"; return *new(Mapping) }

func (ms Mapping) MoveTo(dest Mapping) { _ = "STUB: not implemented"; return }

func (ms Mapping) MemoryStart() uint64 { _ = "STUB: not implemented"; return 0 }

func (ms Mapping) SetMemoryStart(v uint64) { _ = "STUB: not implemented"; return }

func (ms Mapping) MemoryLimit() uint64 { _ = "STUB: not implemented"; return 0 }

func (ms Mapping) SetMemoryLimit(v uint64) { _ = "STUB: not implemented"; return }

func (ms Mapping) FileOffset() uint64 { _ = "STUB: not implemented"; return 0 }

func (ms Mapping) SetFileOffset(v uint64) { _ = "STUB: not implemented"; return }

func (ms Mapping) FilenameStrindex() int32 { _ = "STUB: not implemented"; return 0 }

func (ms Mapping) SetFilenameStrindex(v int32) { _ = "STUB: not implemented"; return }

func (ms Mapping) AttributeIndices() pcommon.Int32Slice {
	_ = "STUB: not implemented"
	return *new(pcommon.Int32Slice)
}

func (ms Mapping) CopyTo(dest Mapping) { _ = "STUB: not implemented"; return }
