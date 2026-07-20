package ptraceotlp

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type ExportRequest struct {
	orig  *internal.ExportTraceServiceRequest
	state *internal.State
}

func NewExportRequest() ExportRequest { _ = "STUB: not implemented"; return *new(ExportRequest) }

func NewExportRequestFromTraces(td ptrace.Traces) ExportRequest {
	_ = "STUB: not implemented"
	return *new(ExportRequest)
}

func (ms ExportRequest) MarshalProto() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ms ExportRequest) UnmarshalProto(data []byte) error { _ = "STUB: not implemented"; return nil }

func (ms ExportRequest) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ms ExportRequest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (ms ExportRequest) Traces() ptrace.Traces {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces)
}
