package pprofile

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type MappingSlice struct {
	orig  *[]*internal.Mapping
	state *internal.State
}

func newMappingSlice(orig *[]*internal.Mapping, state *internal.State) MappingSlice {
	_ = "STUB: not implemented"
	return *new(MappingSlice)
}

func NewMappingSlice() MappingSlice { _ = "STUB: not implemented"; return *new(MappingSlice) }

func (es MappingSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es MappingSlice) At(i int) Mapping { _ = "STUB: not implemented"; return *new(Mapping) }

func (es MappingSlice) All() iter.Seq2[int, Mapping] { _ = "STUB: not implemented"; return nil }

func (es MappingSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es MappingSlice) AppendEmpty() Mapping { _ = "STUB: not implemented"; return *new(Mapping) }

func (es MappingSlice) MoveAndAppendTo(dest MappingSlice) { _ = "STUB: not implemented"; return }

func (es MappingSlice) RemoveIf(f func(Mapping) bool) { _ = "STUB: not implemented"; return }

func (es MappingSlice) CopyTo(dest MappingSlice) { _ = "STUB: not implemented"; return }

func (es MappingSlice) Sort(less func(a, b Mapping) bool) { _ = "STUB: not implemented"; return }
