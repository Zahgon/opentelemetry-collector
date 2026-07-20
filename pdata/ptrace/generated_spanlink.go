package ptrace

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type SpanLink struct {
	orig  *internal.SpanLink
	state *internal.State
}

func newSpanLink(orig *internal.SpanLink, state *internal.State) SpanLink {
	_ = "STUB: not implemented"
	return *new(SpanLink)
}

func NewSpanLink() SpanLink { _ = "STUB: not implemented"; return *new(SpanLink) }

func (ms SpanLink) MoveTo(dest SpanLink) { _ = "STUB: not implemented"; return }

func (ms SpanLink) TraceID() pcommon.TraceID {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID)
}

func (ms SpanLink) SetTraceID(v pcommon.TraceID) { _ = "STUB: not implemented"; return }

func (ms SpanLink) SpanID() pcommon.SpanID { _ = "STUB: not implemented"; return *new(pcommon.SpanID) }

func (ms SpanLink) SetSpanID(v pcommon.SpanID) { _ = "STUB: not implemented"; return }

func (ms SpanLink) TraceState() pcommon.TraceState {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceState)
}

func (ms SpanLink) Attributes() pcommon.Map { _ = "STUB: not implemented"; return *new(pcommon.Map) }

func (ms SpanLink) DroppedAttributesCount() uint32 { _ = "STUB: not implemented"; return 0 }

func (ms SpanLink) SetDroppedAttributesCount(v uint32) { _ = "STUB: not implemented"; return }

func (ms SpanLink) Flags() uint32 { _ = "STUB: not implemented"; return 0 }

func (ms SpanLink) SetFlags(v uint32) { _ = "STUB: not implemented"; return }

func (ms SpanLink) CopyTo(dest SpanLink) { _ = "STUB: not implemented"; return }
