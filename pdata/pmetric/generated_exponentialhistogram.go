package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type ExponentialHistogram struct {
	orig  *internal.ExponentialHistogram
	state *internal.State
}

func newExponentialHistogram(orig *internal.ExponentialHistogram, state *internal.State) ExponentialHistogram {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogram)
}

func NewExponentialHistogram() ExponentialHistogram {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogram)
}

func (ms ExponentialHistogram) MoveTo(dest ExponentialHistogram) { _ = "STUB: not implemented"; return }

func (ms ExponentialHistogram) DataPoints() ExponentialHistogramDataPointSlice {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogramDataPointSlice)
}

func (ms ExponentialHistogram) AggregationTemporality() AggregationTemporality {
	_ = "STUB: not implemented"
	return *new(AggregationTemporality)
}

func (ms ExponentialHistogram) SetAggregationTemporality(v AggregationTemporality) {
	_ = "STUB: not implemented"
	return
}

func (ms ExponentialHistogram) CopyTo(dest ExponentialHistogram) { _ = "STUB: not implemented"; return }
