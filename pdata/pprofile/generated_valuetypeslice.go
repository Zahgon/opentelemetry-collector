package pprofile

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type ValueTypeSlice struct {
	orig  *[]*internal.ValueType
	state *internal.State
}

func newValueTypeSlice(orig *[]*internal.ValueType, state *internal.State) ValueTypeSlice {
	_ = "STUB: not implemented"
	return *new(ValueTypeSlice)
}

func NewValueTypeSlice() ValueTypeSlice { _ = "STUB: not implemented"; return *new(ValueTypeSlice) }

func (es ValueTypeSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es ValueTypeSlice) At(i int) ValueType { _ = "STUB: not implemented"; return *new(ValueType) }

func (es ValueTypeSlice) All() iter.Seq2[int, ValueType] { _ = "STUB: not implemented"; return nil }

func (es ValueTypeSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es ValueTypeSlice) AppendEmpty() ValueType { _ = "STUB: not implemented"; return *new(ValueType) }

func (es ValueTypeSlice) MoveAndAppendTo(dest ValueTypeSlice) { _ = "STUB: not implemented"; return }

func (es ValueTypeSlice) RemoveIf(f func(ValueType) bool) { _ = "STUB: not implemented"; return }

func (es ValueTypeSlice) CopyTo(dest ValueTypeSlice) { _ = "STUB: not implemented"; return }

func (es ValueTypeSlice) Sort(less func(a, b ValueType) bool) { _ = "STUB: not implemented"; return }
