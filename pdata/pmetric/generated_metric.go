package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type Metric struct {
	orig  *internal.Metric
	state *internal.State
}

func newMetric(orig *internal.Metric, state *internal.State) Metric {
	_ = "STUB: not implemented"
	return *new(Metric)
}

func NewMetric() Metric { _ = "STUB: not implemented"; return *new(Metric) }

func (ms Metric) MoveTo(dest Metric) { _ = "STUB: not implemented"; return }

func (ms Metric) Name() string { _ = "STUB: not implemented"; return "" }

func (ms Metric) SetName(v string) { _ = "STUB: not implemented"; return }

func (ms Metric) Description() string { _ = "STUB: not implemented"; return "" }

func (ms Metric) SetDescription(v string) { _ = "STUB: not implemented"; return }

func (ms Metric) Unit() string { _ = "STUB: not implemented"; return "" }

func (ms Metric) SetUnit(v string) { _ = "STUB: not implemented"; return }

func (ms Metric) Type() MetricType { _ = "STUB: not implemented"; return *new(MetricType) }

func (ms Metric) Gauge() Gauge { _ = "STUB: not implemented"; return *new(Gauge) }

func (ms Metric) SetEmptyGauge() Gauge { _ = "STUB: not implemented"; return *new(Gauge) }

func (ms Metric) Sum() Sum { _ = "STUB: not implemented"; return *new(Sum) }

func (ms Metric) SetEmptySum() Sum { _ = "STUB: not implemented"; return *new(Sum) }

func (ms Metric) Histogram() Histogram { _ = "STUB: not implemented"; return *new(Histogram) }

func (ms Metric) SetEmptyHistogram() Histogram { _ = "STUB: not implemented"; return *new(Histogram) }

func (ms Metric) ExponentialHistogram() ExponentialHistogram {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogram)
}

func (ms Metric) SetEmptyExponentialHistogram() ExponentialHistogram {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogram)
}

func (ms Metric) Summary() Summary { _ = "STUB: not implemented"; return *new(Summary) }

func (ms Metric) SetEmptySummary() Summary { _ = "STUB: not implemented"; return *new(Summary) }

func (ms Metric) Metadata() pcommon.Map { _ = "STUB: not implemented"; return *new(pcommon.Map) }

func (ms Metric) CopyTo(dest Metric) { _ = "STUB: not implemented"; return }
