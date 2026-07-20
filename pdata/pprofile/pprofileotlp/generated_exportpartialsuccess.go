package pprofileotlp

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type ExportPartialSuccess struct {
	orig  *internal.ExportProfilesPartialSuccess
	state *internal.State
}

func newExportPartialSuccess(orig *internal.ExportProfilesPartialSuccess, state *internal.State) ExportPartialSuccess {
	_ = "STUB: not implemented"
	return *new(ExportPartialSuccess)
}

func NewExportPartialSuccess() ExportPartialSuccess {
	_ = "STUB: not implemented"
	return *new(ExportPartialSuccess)
}

func (ms ExportPartialSuccess) MoveTo(dest ExportPartialSuccess) { _ = "STUB: not implemented"; return }

func (ms ExportPartialSuccess) RejectedProfiles() int64 { _ = "STUB: not implemented"; return 0 }

func (ms ExportPartialSuccess) SetRejectedProfiles(v int64) { _ = "STUB: not implemented"; return }

func (ms ExportPartialSuccess) ErrorMessage() string { _ = "STUB: not implemented"; return "" }

func (ms ExportPartialSuccess) SetErrorMessage(v string) { _ = "STUB: not implemented"; return }

func (ms ExportPartialSuccess) CopyTo(dest ExportPartialSuccess) { _ = "STUB: not implemented"; return }
