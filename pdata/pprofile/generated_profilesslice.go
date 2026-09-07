package pprofile

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type ProfilesSlice struct {
	orig  *[]*internal.Profile
	state *internal.State
}

func newProfilesSlice(orig *[]*internal.Profile, state *internal.State) ProfilesSlice {
	_ = "STUB: not implemented"
	return *new(ProfilesSlice)
}

func NewProfilesSlice() ProfilesSlice { _ = "STUB: not implemented"; return *new(ProfilesSlice) }

func (es ProfilesSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es ProfilesSlice) At(i int) Profile { _ = "STUB: not implemented"; return *new(Profile) }

func (es ProfilesSlice) All() iter.Seq2[int, Profile] { _ = "STUB: not implemented"; return nil }

func (es ProfilesSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es ProfilesSlice) AppendEmpty() Profile { _ = "STUB: not implemented"; return *new(Profile) }

func (es ProfilesSlice) MoveAndAppendTo(dest ProfilesSlice) { _ = "STUB: not implemented"; return }

func (es ProfilesSlice) RemoveIf(f func(Profile) bool) { _ = "STUB: not implemented"; return }

func (es ProfilesSlice) CopyTo(dest ProfilesSlice) { _ = "STUB: not implemented"; return }

func (es ProfilesSlice) Sort(less func(a, b Profile) bool) { _ = "STUB: not implemented"; return }
