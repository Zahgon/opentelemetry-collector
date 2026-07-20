package ptrace

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type SpanEvent struct {
	orig  *internal.SpanEvent
	state *internal.State
}

func newSpanEvent(orig *internal.SpanEvent, state *internal.State) SpanEvent {
	_ = "STUB: not implemented"
	return *new(SpanEvent)
}

func NewSpanEvent() SpanEvent { _ = "STUB: not implemented"; return *new(SpanEvent) }

func (ms SpanEvent) MoveTo(dest SpanEvent) { _ = "STUB: not implemented"; return }

func (ms SpanEvent) Timestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms SpanEvent) SetTimestamp(v pcommon.Timestamp) { _ = "STUB: not implemented"; return }

func (ms SpanEvent) Name() string { _ = "STUB: not implemented"; return "" }

func (ms SpanEvent) SetName(v string) { _ = "STUB: not implemented"; return }

func (ms SpanEvent) Attributes() pcommon.Map { _ = "STUB: not implemented"; return *new(pcommon.Map) }

func (ms SpanEvent) DroppedAttributesCount() uint32 { _ = "STUB: not implemented"; return 0 }

func (ms SpanEvent) SetDroppedAttributesCount(v uint32) { _ = "STUB: not implemented"; return }

func (ms SpanEvent) CopyTo(dest SpanEvent) { _ = "STUB: not implemented"; return }
