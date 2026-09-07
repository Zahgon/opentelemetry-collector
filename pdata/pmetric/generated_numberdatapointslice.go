package pmetric

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type NumberDataPointSlice struct {
	orig  *[]*internal.NumberDataPoint
	state *internal.State
}

func newNumberDataPointSlice(orig *[]*internal.NumberDataPoint, state *internal.State) NumberDataPointSlice {
	_ = "STUB: not implemented"
	return *new(NumberDataPointSlice)
}

func NewNumberDataPointSlice() NumberDataPointSlice {
	_ = "STUB: not implemented"
	return *new(NumberDataPointSlice)
}

func (es NumberDataPointSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es NumberDataPointSlice) At(i int) NumberDataPoint {
	_ = "STUB: not implemented"
	return *new(NumberDataPoint)
}

func (es NumberDataPointSlice) All() iter.Seq2[int, NumberDataPoint] {
	_ = "STUB: not implemented"
	return nil
}

func (es NumberDataPointSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es NumberDataPointSlice) AppendEmpty() NumberDataPoint {
	_ = "STUB: not implemented"
	return *new(NumberDataPoint)
}

func (es NumberDataPointSlice) MoveAndAppendTo(dest NumberDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func (es NumberDataPointSlice) RemoveIf(f func(NumberDataPoint) bool) {
	_ = "STUB: not implemented"
	return
}

func (es NumberDataPointSlice) CopyTo(dest NumberDataPointSlice) { _ = "STUB: not implemented"; return }

func (es NumberDataPointSlice) Sort(less func(a, b NumberDataPoint) bool) {
	_ = "STUB: not implemented"
	return
}
