package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ScopeProfiles struct {
	orig  *internal.ScopeProfiles
	state *internal.State
}

func newScopeProfiles(orig *internal.ScopeProfiles, state *internal.State) ScopeProfiles {
	_ = "STUB: not implemented"
	return *new(ScopeProfiles)
}

func NewScopeProfiles() ScopeProfiles { _ = "STUB: not implemented"; return *new(ScopeProfiles) }

func (ms ScopeProfiles) MoveTo(dest ScopeProfiles) { _ = "STUB: not implemented"; return }

func (ms ScopeProfiles) Scope() pcommon.InstrumentationScope {
	_ = "STUB: not implemented"
	return *new(pcommon.InstrumentationScope)
}

func (ms ScopeProfiles) Profiles() ProfilesSlice {
	_ = "STUB: not implemented"
	return *new(ProfilesSlice)
}

func (ms ScopeProfiles) SchemaUrl() string { _ = "STUB: not implemented"; return "" }

func (ms ScopeProfiles) SetSchemaUrl(v string) { _ = "STUB: not implemented"; return }

func (ms ScopeProfiles) CopyTo(dest ScopeProfiles) { _ = "STUB: not implemented"; return }
