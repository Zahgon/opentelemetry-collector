package plogotlp

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/plog"
)

type ExportRequest struct {
	orig  *internal.ExportLogsServiceRequest
	state *internal.State
}

func NewExportRequest() ExportRequest { _ = "STUB: not implemented"; return *new(ExportRequest) }

func NewExportRequestFromLogs(ld plog.Logs) ExportRequest {
	_ = "STUB: not implemented"
	return *new(ExportRequest)
}

func (ms ExportRequest) MarshalProto() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ms ExportRequest) UnmarshalProto(data []byte) error { _ = "STUB: not implemented"; return nil }

func (ms ExportRequest) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ms ExportRequest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (ms ExportRequest) Logs() plog.Logs { _ = "STUB: not implemented"; return *new(plog.Logs) }
