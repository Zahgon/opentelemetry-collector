package ptrace

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type Span struct {
	orig  *internal.Span
	state *internal.State
}

func newSpan(orig *internal.Span, state *internal.State) Span {
	_ = "STUB: not implemented"
	return *new(Span)
}

func NewSpan() Span { _ = "STUB: not implemented"; return *new(Span) }

func (ms Span) MoveTo(dest Span) { _ = "STUB: not implemented"; return }

func (ms Span) TraceID() pcommon.TraceID { _ = "STUB: not implemented"; return *new(pcommon.TraceID) }

func (ms Span) SetTraceID(v pcommon.TraceID) { _ = "STUB: not implemented"; return }

func (ms Span) SpanID() pcommon.SpanID { _ = "STUB: not implemented"; return *new(pcommon.SpanID) }

func (ms Span) SetSpanID(v pcommon.SpanID) { _ = "STUB: not implemented"; return }

func (ms Span) TraceState() pcommon.TraceState {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceState)
}

func (ms Span) ParentSpanID() pcommon.SpanID {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID)
}

func (ms Span) SetParentSpanID(v pcommon.SpanID) { _ = "STUB: not implemented"; return }

func (ms Span) Flags() uint32 { _ = "STUB: not implemented"; return 0 }

func (ms Span) SetFlags(v uint32) { _ = "STUB: not implemented"; return }

func (ms Span) Name() string { _ = "STUB: not implemented"; return "" }

func (ms Span) SetName(v string) { _ = "STUB: not implemented"; return }

func (ms Span) Kind() SpanKind { _ = "STUB: not implemented"; return *new(SpanKind) }

func (ms Span) SetKind(v SpanKind) { _ = "STUB: not implemented"; return }

func (ms Span) StartTimestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms Span) SetStartTimestamp(v pcommon.Timestamp) { _ = "STUB: not implemented"; return }

func (ms Span) EndTimestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms Span) SetEndTimestamp(v pcommon.Timestamp) { _ = "STUB: not implemented"; return }

func (ms Span) Attributes() pcommon.Map { _ = "STUB: not implemented"; return *new(pcommon.Map) }

func (ms Span) DroppedAttributesCount() uint32 { _ = "STUB: not implemented"; return 0 }

func (ms Span) SetDroppedAttributesCount(v uint32) { _ = "STUB: not implemented"; return }

func (ms Span) Events() SpanEventSlice { _ = "STUB: not implemented"; return *new(SpanEventSlice) }

func (ms Span) DroppedEventsCount() uint32 { _ = "STUB: not implemented"; return 0 }

func (ms Span) SetDroppedEventsCount(v uint32) { _ = "STUB: not implemented"; return }

func (ms Span) Links() SpanLinkSlice { _ = "STUB: not implemented"; return *new(SpanLinkSlice) }

func (ms Span) DroppedLinksCount() uint32 { _ = "STUB: not implemented"; return 0 }

func (ms Span) SetDroppedLinksCount(v uint32) { _ = "STUB: not implemented"; return }

func (ms Span) Status() Status { _ = "STUB: not implemented"; return *new(Status) }

func (ms Span) CopyTo(dest Span) { _ = "STUB: not implemented"; return }
