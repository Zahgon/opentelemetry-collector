package pmetricotlp

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type ExportResponse struct {
	orig  *internal.ExportMetricsServiceResponse
	state *internal.State
}

func newExportResponse(orig *internal.ExportMetricsServiceResponse, state *internal.State) ExportResponse {
	_ = "STUB: not implemented"
	return *new(ExportResponse)
}

func NewExportResponse() ExportResponse { _ = "STUB: not implemented"; return *new(ExportResponse) }

func (ms ExportResponse) MoveTo(dest ExportResponse) { _ = "STUB: not implemented"; return }

func (ms ExportResponse) PartialSuccess() ExportPartialSuccess {
	_ = "STUB: not implemented"
	return *new(ExportPartialSuccess)
}

func (ms ExportResponse) CopyTo(dest ExportResponse) { _ = "STUB: not implemented"; return }
