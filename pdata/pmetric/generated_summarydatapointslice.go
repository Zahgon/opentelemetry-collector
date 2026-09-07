package pmetric

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type SummaryDataPointSlice struct {
	orig  *[]*internal.SummaryDataPoint
	state *internal.State
}

func newSummaryDataPointSlice(orig *[]*internal.SummaryDataPoint, state *internal.State) SummaryDataPointSlice {
	_ = "STUB: not implemented"
	return *new(SummaryDataPointSlice)
}

func NewSummaryDataPointSlice() SummaryDataPointSlice {
	_ = "STUB: not implemented"
	return *new(SummaryDataPointSlice)
}

func (es SummaryDataPointSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es SummaryDataPointSlice) At(i int) SummaryDataPoint {
	_ = "STUB: not implemented"
	return *new(SummaryDataPoint)
}

func (es SummaryDataPointSlice) All() iter.Seq2[int, SummaryDataPoint] {
	_ = "STUB: not implemented"
	return nil
}

func (es SummaryDataPointSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es SummaryDataPointSlice) AppendEmpty() SummaryDataPoint {
	_ = "STUB: not implemented"
	return *new(SummaryDataPoint)
}

func (es SummaryDataPointSlice) MoveAndAppendTo(dest SummaryDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func (es SummaryDataPointSlice) RemoveIf(f func(SummaryDataPoint) bool) {
	_ = "STUB: not implemented"
	return
}

func (es SummaryDataPointSlice) CopyTo(dest SummaryDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func (es SummaryDataPointSlice) Sort(less func(a, b SummaryDataPoint) bool) {
	_ = "STUB: not implemented"
	return
}
