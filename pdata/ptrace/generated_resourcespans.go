package ptrace

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ResourceSpans struct {
	orig  *internal.ResourceSpans
	state *internal.State
}

func newResourceSpans(orig *internal.ResourceSpans, state *internal.State) ResourceSpans {
	_ = "STUB: not implemented"
	return *new(ResourceSpans)
}

func NewResourceSpans() ResourceSpans { _ = "STUB: not implemented"; return *new(ResourceSpans) }

func (ms ResourceSpans) MoveTo(dest ResourceSpans) { _ = "STUB: not implemented"; return }

func (ms ResourceSpans) Resource() pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}

func (ms ResourceSpans) ScopeSpans() ScopeSpansSlice {
	_ = "STUB: not implemented"
	return *new(ScopeSpansSlice)
}

func (ms ResourceSpans) SchemaUrl() string { _ = "STUB: not implemented"; return "" }

func (ms ResourceSpans) SetSchemaUrl(v string) { _ = "STUB: not implemented"; return }

func (ms ResourceSpans) CopyTo(dest ResourceSpans) { _ = "STUB: not implemented"; return }
