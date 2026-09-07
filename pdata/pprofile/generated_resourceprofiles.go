package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ResourceProfiles struct {
	orig  *internal.ResourceProfiles
	state *internal.State
}

func newResourceProfiles(orig *internal.ResourceProfiles, state *internal.State) ResourceProfiles {
	_ = "STUB: not implemented"
	return *new(ResourceProfiles)
}

func NewResourceProfiles() ResourceProfiles {
	_ = "STUB: not implemented"
	return *new(ResourceProfiles)
}

func (ms ResourceProfiles) MoveTo(dest ResourceProfiles) { _ = "STUB: not implemented"; return }

func (ms ResourceProfiles) Resource() pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}

func (ms ResourceProfiles) ScopeProfiles() ScopeProfilesSlice {
	_ = "STUB: not implemented"
	return *new(ScopeProfilesSlice)
}

func (ms ResourceProfiles) SchemaUrl() string { _ = "STUB: not implemented"; return "" }

func (ms ResourceProfiles) SetSchemaUrl(v string) { _ = "STUB: not implemented"; return }

func (ms ResourceProfiles) CopyTo(dest ResourceProfiles) { _ = "STUB: not implemented"; return }
