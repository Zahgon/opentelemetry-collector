package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type Sum struct {
	orig  *internal.Sum
	state *internal.State
}

func newSum(orig *internal.Sum, state *internal.State) Sum {
	_ = "STUB: not implemented"
	return *new(Sum)
}

func NewSum() Sum { _ = "STUB: not implemented"; return *new(Sum) }

func (ms Sum) MoveTo(dest Sum) { _ = "STUB: not implemented"; return }

func (ms Sum) DataPoints() NumberDataPointSlice {
	_ = "STUB: not implemented"
	return *new(NumberDataPointSlice)
}

func (ms Sum) AggregationTemporality() AggregationTemporality {
	_ = "STUB: not implemented"
	return *new(AggregationTemporality)
}

func (ms Sum) SetAggregationTemporality(v AggregationTemporality) {
	_ = "STUB: not implemented"
	return
}

func (ms Sum) IsMonotonic() bool { _ = "STUB: not implemented"; return false }

func (ms Sum) SetIsMonotonic(v bool) { _ = "STUB: not implemented"; return }

func (ms Sum) CopyTo(dest Sum) { _ = "STUB: not implemented"; return }
