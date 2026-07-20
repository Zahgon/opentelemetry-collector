package pmetric

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type ScopeMetricsSlice struct {
	orig  *[]*internal.ScopeMetrics
	state *internal.State
}

func newScopeMetricsSlice(orig *[]*internal.ScopeMetrics, state *internal.State) ScopeMetricsSlice {
	_ = "STUB: not implemented"
	return *new(ScopeMetricsSlice)
}

func NewScopeMetricsSlice() ScopeMetricsSlice {
	_ = "STUB: not implemented"
	return *new(ScopeMetricsSlice)
}

func (es ScopeMetricsSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es ScopeMetricsSlice) At(i int) ScopeMetrics {
	_ = "STUB: not implemented"
	return *new(ScopeMetrics)
}

func (es ScopeMetricsSlice) All() iter.Seq2[int, ScopeMetrics] {
	_ = "STUB: not implemented"
	return nil
}

func (es ScopeMetricsSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es ScopeMetricsSlice) AppendEmpty() ScopeMetrics {
	_ = "STUB: not implemented"
	return *new(ScopeMetrics)
}

func (es ScopeMetricsSlice) MoveAndAppendTo(dest ScopeMetricsSlice) {
	_ = "STUB: not implemented"
	return
}

func (es ScopeMetricsSlice) RemoveIf(f func(ScopeMetrics) bool) { _ = "STUB: not implemented"; return }

func (es ScopeMetricsSlice) CopyTo(dest ScopeMetricsSlice) { _ = "STUB: not implemented"; return }

func (es ScopeMetricsSlice) Sort(less func(a, b ScopeMetrics) bool) {
	_ = "STUB: not implemented"
	return
}
