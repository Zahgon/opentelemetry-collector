package pmetricotlp

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type ExportRequest struct {
	orig  *internal.ExportMetricsServiceRequest
	state *internal.State
}

func NewExportRequest() ExportRequest { _ = "STUB: not implemented"; return *new(ExportRequest) }

func NewExportRequestFromMetrics(md pmetric.Metrics) ExportRequest {
	_ = "STUB: not implemented"
	return *new(ExportRequest)
}

func (ms ExportRequest) MarshalProto() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ms ExportRequest) UnmarshalProto(data []byte) error { _ = "STUB: not implemented"; return nil }

func (ms ExportRequest) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ms ExportRequest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (ms ExportRequest) Metrics() pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}
