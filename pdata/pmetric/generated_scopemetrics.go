package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ScopeMetrics struct {
	orig  *internal.ScopeMetrics
	state *internal.State
}

func newScopeMetrics(orig *internal.ScopeMetrics, state *internal.State) ScopeMetrics {
	_ = "STUB: not implemented"
	return *new(ScopeMetrics)
}

func NewScopeMetrics() ScopeMetrics { _ = "STUB: not implemented"; return *new(ScopeMetrics) }

func (ms ScopeMetrics) MoveTo(dest ScopeMetrics) { _ = "STUB: not implemented"; return }

func (ms ScopeMetrics) Scope() pcommon.InstrumentationScope {
	_ = "STUB: not implemented"
	return *new(pcommon.InstrumentationScope)
}

func (ms ScopeMetrics) Metrics() MetricSlice { _ = "STUB: not implemented"; return *new(MetricSlice) }

func (ms ScopeMetrics) SchemaUrl() string { _ = "STUB: not implemented"; return "" }

func (ms ScopeMetrics) SetSchemaUrl(v string) { _ = "STUB: not implemented"; return }

func (ms ScopeMetrics) CopyTo(dest ScopeMetrics) { _ = "STUB: not implemented"; return }
