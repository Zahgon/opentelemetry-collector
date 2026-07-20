package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type Metrics internal.MetricsWrapper

func newMetrics(orig *internal.ExportMetricsServiceRequest, state *internal.State) Metrics {
	_ = "STUB: not implemented"
	return *new(Metrics)
}

func NewMetrics() Metrics { _ = "STUB: not implemented"; return *new(Metrics) }

func (ms Metrics) MoveTo(dest Metrics) { _ = "STUB: not implemented"; return }

func (ms Metrics) ResourceMetrics() ResourceMetricsSlice {
	_ = "STUB: not implemented"
	return *new(ResourceMetricsSlice)
}

func (ms Metrics) CopyTo(dest Metrics) { _ = "STUB: not implemented"; return }

func (ms Metrics) getOrig() *internal.ExportMetricsServiceRequest {
	_ = "STUB: not implemented"
	return nil
}

func (ms Metrics) getState() *internal.State { _ = "STUB: not implemented"; return nil }
