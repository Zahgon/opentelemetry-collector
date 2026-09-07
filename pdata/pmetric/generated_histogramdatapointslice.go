package pmetric

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type HistogramDataPointSlice struct {
	orig  *[]*internal.HistogramDataPoint
	state *internal.State
}

func newHistogramDataPointSlice(orig *[]*internal.HistogramDataPoint, state *internal.State) HistogramDataPointSlice {
	_ = "STUB: not implemented"
	return *new(HistogramDataPointSlice)
}

func NewHistogramDataPointSlice() HistogramDataPointSlice {
	_ = "STUB: not implemented"
	return *new(HistogramDataPointSlice)
}

func (es HistogramDataPointSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es HistogramDataPointSlice) At(i int) HistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(HistogramDataPoint)
}

func (es HistogramDataPointSlice) All() iter.Seq2[int, HistogramDataPoint] {
	_ = "STUB: not implemented"
	return nil
}

func (es HistogramDataPointSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es HistogramDataPointSlice) AppendEmpty() HistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(HistogramDataPoint)
}

func (es HistogramDataPointSlice) MoveAndAppendTo(dest HistogramDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func (es HistogramDataPointSlice) RemoveIf(f func(HistogramDataPoint) bool) {
	_ = "STUB: not implemented"
	return
}

func (es HistogramDataPointSlice) CopyTo(dest HistogramDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func (es HistogramDataPointSlice) Sort(less func(a, b HistogramDataPoint) bool) {
	_ = "STUB: not implemented"
	return
}
