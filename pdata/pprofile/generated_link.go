package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type Link struct {
	orig  *internal.Link
	state *internal.State
}

func newLink(orig *internal.Link, state *internal.State) Link {
	_ = "STUB: not implemented"
	return *new(Link)
}

func NewLink() Link { _ = "STUB: not implemented"; return *new(Link) }

func (ms Link) MoveTo(dest Link) { _ = "STUB: not implemented"; return }

func (ms Link) TraceID() pcommon.TraceID { _ = "STUB: not implemented"; return *new(pcommon.TraceID) }

func (ms Link) SetTraceID(v pcommon.TraceID) { _ = "STUB: not implemented"; return }

func (ms Link) SpanID() pcommon.SpanID { _ = "STUB: not implemented"; return *new(pcommon.SpanID) }

func (ms Link) SetSpanID(v pcommon.SpanID) { _ = "STUB: not implemented"; return }

func (ms Link) CopyTo(dest Link) { _ = "STUB: not implemented"; return }
