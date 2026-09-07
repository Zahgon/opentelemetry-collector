package ptrace

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type Traces internal.TracesWrapper

func newTraces(orig *internal.ExportTraceServiceRequest, state *internal.State) Traces {
	_ = "STUB: not implemented"
	return *new(Traces)
}

func NewTraces() Traces { _ = "STUB: not implemented"; return *new(Traces) }

func (ms Traces) MoveTo(dest Traces) { _ = "STUB: not implemented"; return }

func (ms Traces) ResourceSpans() ResourceSpansSlice {
	_ = "STUB: not implemented"
	return *new(ResourceSpansSlice)
}

func (ms Traces) CopyTo(dest Traces) { _ = "STUB: not implemented"; return }

func (ms Traces) getOrig() *internal.ExportTraceServiceRequest {
	_ = "STUB: not implemented"
	return nil
}

func (ms Traces) getState() *internal.State { _ = "STUB: not implemented"; return nil }
