package xexporterhelper

import (
	"context"

	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/sizer"
	"go.opentelemetry.io/collector/pdata/pprofile"
)

func (req *profilesRequest) MergeSplit(_ context.Context, maxSize int, szt exporterhelper.RequestSizerType, r2 Request) ([]Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (req *profilesRequest) mergeTo(dst *profilesRequest, sz sizer.ProfilesSizer) error {
	_ = "STUB: not implemented"
	return nil
}

func (req *profilesRequest) split(maxSize int, sz sizer.ProfilesSizer) ([]Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractProfiles(srcProfiles pprofile.Profiles, capacity int, sz sizer.ProfilesSizer) (pprofile.Profiles, int) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), 0
}

func extractResourceProfiles(srcRP pprofile.ResourceProfiles, capacity int, sz sizer.ProfilesSizer) (pprofile.ResourceProfiles, int) {
	_ = "STUB: not implemented"
	return *new(pprofile.ResourceProfiles), 0
}

func extractScopeProfiles(srcSS pprofile.ScopeProfiles, capacity int, sz sizer.ProfilesSizer) (pprofile.ScopeProfiles, int) {
	_ = "STUB: not implemented"
	return *new(pprofile.ScopeProfiles), 0
}
