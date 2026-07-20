package pmetric

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type ExponentialHistogramDataPointSlice struct {
	orig  *[]*internal.ExponentialHistogramDataPoint
	state *internal.State
}

func newExponentialHistogramDataPointSlice(orig *[]*internal.ExponentialHistogramDataPoint, state *internal.State) ExponentialHistogramDataPointSlice {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogramDataPointSlice)
}

func NewExponentialHistogramDataPointSlice() ExponentialHistogramDataPointSlice {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogramDataPointSlice)
}

func (es ExponentialHistogramDataPointSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es ExponentialHistogramDataPointSlice) At(i int) ExponentialHistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogramDataPoint)
}

func (es ExponentialHistogramDataPointSlice) All() iter.Seq2[int, ExponentialHistogramDataPoint] {
	_ = "STUB: not implemented"
	return nil
}

func (es ExponentialHistogramDataPointSlice) EnsureCapacity(newCap int) {
	_ = "STUB: not implemented"
	return
}

func (es ExponentialHistogramDataPointSlice) AppendEmpty() ExponentialHistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogramDataPoint)
}

func (es ExponentialHistogramDataPointSlice) MoveAndAppendTo(dest ExponentialHistogramDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func (es ExponentialHistogramDataPointSlice) RemoveIf(f func(ExponentialHistogramDataPoint) bool) {
	_ = "STUB: not implemented"
	return
}

func (es ExponentialHistogramDataPointSlice) CopyTo(dest ExponentialHistogramDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func (es ExponentialHistogramDataPointSlice) Sort(less func(a, b ExponentialHistogramDataPoint) bool) {
	_ = "STUB: not implemented"
	return
}
