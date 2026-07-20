package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ResourceMetrics struct {
	orig  *internal.ResourceMetrics
	state *internal.State
}

func newResourceMetrics(orig *internal.ResourceMetrics, state *internal.State) ResourceMetrics {
	_ = "STUB: not implemented"
	return *new(ResourceMetrics)
}

func NewResourceMetrics() ResourceMetrics { _ = "STUB: not implemented"; return *new(ResourceMetrics) }

func (ms ResourceMetrics) MoveTo(dest ResourceMetrics) { _ = "STUB: not implemented"; return }

func (ms ResourceMetrics) Resource() pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}

func (ms ResourceMetrics) ScopeMetrics() ScopeMetricsSlice {
	_ = "STUB: not implemented"
	return *new(ScopeMetricsSlice)
}

func (ms ResourceMetrics) SchemaUrl() string { _ = "STUB: not implemented"; return "" }

func (ms ResourceMetrics) SetSchemaUrl(v string) { _ = "STUB: not implemented"; return }

func (ms ResourceMetrics) CopyTo(dest ResourceMetrics) { _ = "STUB: not implemented"; return }
