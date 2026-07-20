package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type Histogram struct {
	orig  *internal.Histogram
	state *internal.State
}

func newHistogram(orig *internal.Histogram, state *internal.State) Histogram {
	_ = "STUB: not implemented"
	return *new(Histogram)
}

func NewHistogram() Histogram { _ = "STUB: not implemented"; return *new(Histogram) }

func (ms Histogram) MoveTo(dest Histogram) { _ = "STUB: not implemented"; return }

func (ms Histogram) DataPoints() HistogramDataPointSlice {
	_ = "STUB: not implemented"
	return *new(HistogramDataPointSlice)
}

func (ms Histogram) AggregationTemporality() AggregationTemporality {
	_ = "STUB: not implemented"
	return *new(AggregationTemporality)
}

func (ms Histogram) SetAggregationTemporality(v AggregationTemporality) {
	_ = "STUB: not implemented"
	return
}

func (ms Histogram) CopyTo(dest Histogram) { _ = "STUB: not implemented"; return }
