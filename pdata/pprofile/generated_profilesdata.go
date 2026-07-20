package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type ProfilesData internal.ProfilesDataWrapper

func newProfilesData(orig *internal.ProfilesData, state *internal.State) ProfilesData {
	_ = "STUB: not implemented"
	return *new(ProfilesData)
}

func NewProfilesData() ProfilesData { _ = "STUB: not implemented"; return *new(ProfilesData) }

func (ms ProfilesData) MoveTo(dest ProfilesData) { _ = "STUB: not implemented"; return }

func (ms ProfilesData) ResourceProfiles() ResourceProfilesSlice {
	_ = "STUB: not implemented"
	return *new(ResourceProfilesSlice)
}

func (ms ProfilesData) Dictionary() ProfilesDictionary {
	_ = "STUB: not implemented"
	return *new(ProfilesDictionary)
}

func (ms ProfilesData) CopyTo(dest ProfilesData) { _ = "STUB: not implemented"; return }

func (ms ProfilesData) getOrig() *internal.ProfilesData { _ = "STUB: not implemented"; return nil }

func (ms ProfilesData) getState() *internal.State { _ = "STUB: not implemented"; return nil }
