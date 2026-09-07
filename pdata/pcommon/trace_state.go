package pcommon

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type TraceState internal.TraceStateWrapper

func NewTraceState() TraceState { _ = "STUB: not implemented"; return *new(TraceState) }

func (ms TraceState) getOrig() *string { _ = "STUB: not implemented"; return nil }

func (ms TraceState) getState() *internal.State { _ = "STUB: not implemented"; return nil }

func (ms TraceState) AsRaw() string { _ = "STUB: not implemented"; return "" }

func (ms TraceState) FromRaw(v string) { _ = "STUB: not implemented"; return }

func (ms TraceState) MoveTo(dest TraceState) { _ = "STUB: not implemented"; return }

func (ms TraceState) CopyTo(dest TraceState) { _ = "STUB: not implemented"; return }
