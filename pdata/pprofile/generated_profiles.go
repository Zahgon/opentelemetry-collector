package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type Profiles internal.ProfilesWrapper

func newProfiles(orig *internal.ExportProfilesServiceRequest, state *internal.State) Profiles {
	_ = "STUB: not implemented"
	return *new(Profiles)
}

func NewProfiles() Profiles { _ = "STUB: not implemented"; return *new(Profiles) }

func (ms Profiles) MoveTo(dest Profiles) { _ = "STUB: not implemented"; return }

func (ms Profiles) ResourceProfiles() ResourceProfilesSlice {
	_ = "STUB: not implemented"
	return *new(ResourceProfilesSlice)
}

func (ms Profiles) Dictionary() ProfilesDictionary {
	_ = "STUB: not implemented"
	return *new(ProfilesDictionary)
}

func (ms Profiles) CopyTo(dest Profiles) { _ = "STUB: not implemented"; return }

func (ms Profiles) getOrig() *internal.ExportProfilesServiceRequest {
	_ = "STUB: not implemented"
	return nil
}

func (ms Profiles) getState() *internal.State { _ = "STUB: not implemented"; return nil }
