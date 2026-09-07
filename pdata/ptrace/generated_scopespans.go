package ptrace

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ScopeSpans struct {
	orig  *internal.ScopeSpans
	state *internal.State
}

func newScopeSpans(orig *internal.ScopeSpans, state *internal.State) ScopeSpans {
	_ = "STUB: not implemented"
	return *new(ScopeSpans)
}

func NewScopeSpans() ScopeSpans { _ = "STUB: not implemented"; return *new(ScopeSpans) }

func (ms ScopeSpans) MoveTo(dest ScopeSpans) { _ = "STUB: not implemented"; return }

func (ms ScopeSpans) Scope() pcommon.InstrumentationScope {
	_ = "STUB: not implemented"
	return *new(pcommon.InstrumentationScope)
}

func (ms ScopeSpans) Spans() SpanSlice { _ = "STUB: not implemented"; return *new(SpanSlice) }

func (ms ScopeSpans) SchemaUrl() string { _ = "STUB: not implemented"; return "" }

func (ms ScopeSpans) SetSchemaUrl(v string) { _ = "STUB: not implemented"; return }

func (ms ScopeSpans) CopyTo(dest ScopeSpans) { _ = "STUB: not implemented"; return }
