package pprofile

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type KeyValueAndUnitSlice struct {
	orig  *[]*internal.KeyValueAndUnit
	state *internal.State
}

func newKeyValueAndUnitSlice(orig *[]*internal.KeyValueAndUnit, state *internal.State) KeyValueAndUnitSlice {
	_ = "STUB: not implemented"
	return *new(KeyValueAndUnitSlice)
}

func NewKeyValueAndUnitSlice() KeyValueAndUnitSlice {
	_ = "STUB: not implemented"
	return *new(KeyValueAndUnitSlice)
}

func (es KeyValueAndUnitSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es KeyValueAndUnitSlice) At(i int) KeyValueAndUnit {
	_ = "STUB: not implemented"
	return *new(KeyValueAndUnit)
}

func (es KeyValueAndUnitSlice) All() iter.Seq2[int, KeyValueAndUnit] {
	_ = "STUB: not implemented"
	return nil
}

func (es KeyValueAndUnitSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es KeyValueAndUnitSlice) AppendEmpty() KeyValueAndUnit {
	_ = "STUB: not implemented"
	return *new(KeyValueAndUnit)
}

func (es KeyValueAndUnitSlice) MoveAndAppendTo(dest KeyValueAndUnitSlice) {
	_ = "STUB: not implemented"
	return
}

func (es KeyValueAndUnitSlice) RemoveIf(f func(KeyValueAndUnit) bool) {
	_ = "STUB: not implemented"
	return
}

func (es KeyValueAndUnitSlice) CopyTo(dest KeyValueAndUnitSlice) { _ = "STUB: not implemented"; return }

func (es KeyValueAndUnitSlice) Sort(less func(a, b KeyValueAndUnit) bool) {
	_ = "STUB: not implemented"
	return
}
