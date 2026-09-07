package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type Gauge struct {
	orig  *internal.Gauge
	state *internal.State
}

func newGauge(orig *internal.Gauge, state *internal.State) Gauge {
	_ = "STUB: not implemented"
	return *new(Gauge)
}

func NewGauge() Gauge { _ = "STUB: not implemented"; return *new(Gauge) }

func (ms Gauge) MoveTo(dest Gauge) { _ = "STUB: not implemented"; return }

func (ms Gauge) DataPoints() NumberDataPointSlice {
	_ = "STUB: not implemented"
	return *new(NumberDataPointSlice)
}

func (ms Gauge) CopyTo(dest Gauge) { _ = "STUB: not implemented"; return }
