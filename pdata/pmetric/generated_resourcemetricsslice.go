package pmetric

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type ResourceMetricsSlice struct {
	orig  *[]*internal.ResourceMetrics
	state *internal.State
}

func newResourceMetricsSlice(orig *[]*internal.ResourceMetrics, state *internal.State) ResourceMetricsSlice {
	_ = "STUB: not implemented"
	return *new(ResourceMetricsSlice)
}

func NewResourceMetricsSlice() ResourceMetricsSlice {
	_ = "STUB: not implemented"
	return *new(ResourceMetricsSlice)
}

func (es ResourceMetricsSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es ResourceMetricsSlice) At(i int) ResourceMetrics {
	_ = "STUB: not implemented"
	return *new(ResourceMetrics)
}

func (es ResourceMetricsSlice) All() iter.Seq2[int, ResourceMetrics] {
	_ = "STUB: not implemented"
	return nil
}

func (es ResourceMetricsSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es ResourceMetricsSlice) AppendEmpty() ResourceMetrics {
	_ = "STUB: not implemented"
	return *new(ResourceMetrics)
}

func (es ResourceMetricsSlice) MoveAndAppendTo(dest ResourceMetricsSlice) {
	_ = "STUB: not implemented"
	return
}

func (es ResourceMetricsSlice) RemoveIf(f func(ResourceMetrics) bool) {
	_ = "STUB: not implemented"
	return
}

func (es ResourceMetricsSlice) CopyTo(dest ResourceMetricsSlice) { _ = "STUB: not implemented"; return }

func (es ResourceMetricsSlice) Sort(less func(a, b ResourceMetrics) bool) {
	_ = "STUB: not implemented"
	return
}
