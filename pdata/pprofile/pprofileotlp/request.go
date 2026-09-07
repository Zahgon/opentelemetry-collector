package pprofileotlp

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pprofile"
)

type ExportRequest struct {
	orig  *internal.ExportProfilesServiceRequest
	state *internal.State
}

func NewExportRequest() ExportRequest { _ = "STUB: not implemented"; return *new(ExportRequest) }

func NewExportRequestFromProfiles(td pprofile.Profiles) ExportRequest {
	_ = "STUB: not implemented"
	return *new(ExportRequest)
}

func (ms ExportRequest) MarshalProto() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ms ExportRequest) UnmarshalProto(data []byte) error { _ = "STUB: not implemented"; return nil }

func (ms ExportRequest) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ms ExportRequest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (ms ExportRequest) Profiles() pprofile.Profiles {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles)
}
