package pprofile

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type ResourceProfilesSlice struct {
	orig  *[]*internal.ResourceProfiles
	state *internal.State
}

func newResourceProfilesSlice(orig *[]*internal.ResourceProfiles, state *internal.State) ResourceProfilesSlice {
	_ = "STUB: not implemented"
	return *new(ResourceProfilesSlice)
}

func NewResourceProfilesSlice() ResourceProfilesSlice {
	_ = "STUB: not implemented"
	return *new(ResourceProfilesSlice)
}

func (es ResourceProfilesSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es ResourceProfilesSlice) At(i int) ResourceProfiles {
	_ = "STUB: not implemented"
	return *new(ResourceProfiles)
}

func (es ResourceProfilesSlice) All() iter.Seq2[int, ResourceProfiles] {
	_ = "STUB: not implemented"
	return nil
}

func (es ResourceProfilesSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es ResourceProfilesSlice) AppendEmpty() ResourceProfiles {
	_ = "STUB: not implemented"
	return *new(ResourceProfiles)
}

func (es ResourceProfilesSlice) MoveAndAppendTo(dest ResourceProfilesSlice) {
	_ = "STUB: not implemented"
	return
}

func (es ResourceProfilesSlice) RemoveIf(f func(ResourceProfiles) bool) {
	_ = "STUB: not implemented"
	return
}

func (es ResourceProfilesSlice) CopyTo(dest ResourceProfilesSlice) {
	_ = "STUB: not implemented"
	return
}

func (es ResourceProfilesSlice) Sort(less func(a, b ResourceProfiles) bool) {
	_ = "STUB: not implemented"
	return
}
