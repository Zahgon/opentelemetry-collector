package pmetric

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type SummaryDataPointValueAtQuantileSlice struct {
	orig  *[]*internal.SummaryDataPointValueAtQuantile
	state *internal.State
}

func newSummaryDataPointValueAtQuantileSlice(orig *[]*internal.SummaryDataPointValueAtQuantile, state *internal.State) SummaryDataPointValueAtQuantileSlice {
	_ = "STUB: not implemented"
	return *new(SummaryDataPointValueAtQuantileSlice)
}

func NewSummaryDataPointValueAtQuantileSlice() SummaryDataPointValueAtQuantileSlice {
	_ = "STUB: not implemented"
	return *new(SummaryDataPointValueAtQuantileSlice)
}

func (es SummaryDataPointValueAtQuantileSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es SummaryDataPointValueAtQuantileSlice) At(i int) SummaryDataPointValueAtQuantile {
	_ = "STUB: not implemented"
	return *new(SummaryDataPointValueAtQuantile)
}

func (es SummaryDataPointValueAtQuantileSlice) All() iter.Seq2[int, SummaryDataPointValueAtQuantile] {
	_ = "STUB: not implemented"
	return nil
}

func (es SummaryDataPointValueAtQuantileSlice) EnsureCapacity(newCap int) {
	_ = "STUB: not implemented"
	return
}

func (es SummaryDataPointValueAtQuantileSlice) AppendEmpty() SummaryDataPointValueAtQuantile {
	_ = "STUB: not implemented"
	return *new(SummaryDataPointValueAtQuantile)
}

func (es SummaryDataPointValueAtQuantileSlice) MoveAndAppendTo(dest SummaryDataPointValueAtQuantileSlice) {
	_ = "STUB: not implemented"
	return
}

func (es SummaryDataPointValueAtQuantileSlice) RemoveIf(f func(SummaryDataPointValueAtQuantile) bool) {
	_ = "STUB: not implemented"
	return
}

func (es SummaryDataPointValueAtQuantileSlice) CopyTo(dest SummaryDataPointValueAtQuantileSlice) {
	_ = "STUB: not implemented"
	return
}

func (es SummaryDataPointValueAtQuantileSlice) Sort(less func(a, b SummaryDataPointValueAtQuantile) bool) {
	_ = "STUB: not implemented"
	return
}
