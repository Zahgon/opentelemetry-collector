package pmetric

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type MetricSlice struct {
	orig  *[]*internal.Metric
	state *internal.State
}

func newMetricSlice(orig *[]*internal.Metric, state *internal.State) MetricSlice {
	_ = "STUB: not implemented"
	return *new(MetricSlice)
}

func NewMetricSlice() MetricSlice { _ = "STUB: not implemented"; return *new(MetricSlice) }

func (es MetricSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es MetricSlice) At(i int) Metric { _ = "STUB: not implemented"; return *new(Metric) }

func (es MetricSlice) All() iter.Seq2[int, Metric] { _ = "STUB: not implemented"; return nil }

func (es MetricSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es MetricSlice) AppendEmpty() Metric { _ = "STUB: not implemented"; return *new(Metric) }

func (es MetricSlice) MoveAndAppendTo(dest MetricSlice) { _ = "STUB: not implemented"; return }

func (es MetricSlice) RemoveIf(f func(Metric) bool) { _ = "STUB: not implemented"; return }

func (es MetricSlice) CopyTo(dest MetricSlice) { _ = "STUB: not implemented"; return }

func (es MetricSlice) Sort(less func(a, b Metric) bool) { _ = "STUB: not implemented"; return }
