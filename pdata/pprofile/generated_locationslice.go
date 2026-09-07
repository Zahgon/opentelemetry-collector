package pprofile

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type LocationSlice struct {
	orig  *[]*internal.Location
	state *internal.State
}

func newLocationSlice(orig *[]*internal.Location, state *internal.State) LocationSlice {
	_ = "STUB: not implemented"
	return *new(LocationSlice)
}

func NewLocationSlice() LocationSlice { _ = "STUB: not implemented"; return *new(LocationSlice) }

func (es LocationSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es LocationSlice) At(i int) Location { _ = "STUB: not implemented"; return *new(Location) }

func (es LocationSlice) All() iter.Seq2[int, Location] { _ = "STUB: not implemented"; return nil }

func (es LocationSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es LocationSlice) AppendEmpty() Location { _ = "STUB: not implemented"; return *new(Location) }

func (es LocationSlice) MoveAndAppendTo(dest LocationSlice) { _ = "STUB: not implemented"; return }

func (es LocationSlice) RemoveIf(f func(Location) bool) { _ = "STUB: not implemented"; return }

func (es LocationSlice) CopyTo(dest LocationSlice) { _ = "STUB: not implemented"; return }

func (es LocationSlice) Sort(less func(a, b Location) bool) { _ = "STUB: not implemented"; return }
