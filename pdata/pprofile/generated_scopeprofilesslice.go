package pprofile

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type ScopeProfilesSlice struct {
	orig  *[]*internal.ScopeProfiles
	state *internal.State
}

func newScopeProfilesSlice(orig *[]*internal.ScopeProfiles, state *internal.State) ScopeProfilesSlice {
	_ = "STUB: not implemented"
	return *new(ScopeProfilesSlice)
}

func NewScopeProfilesSlice() ScopeProfilesSlice {
	_ = "STUB: not implemented"
	return *new(ScopeProfilesSlice)
}

func (es ScopeProfilesSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es ScopeProfilesSlice) At(i int) ScopeProfiles {
	_ = "STUB: not implemented"
	return *new(ScopeProfiles)
}

func (es ScopeProfilesSlice) All() iter.Seq2[int, ScopeProfiles] {
	_ = "STUB: not implemented"
	return nil
}

func (es ScopeProfilesSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es ScopeProfilesSlice) AppendEmpty() ScopeProfiles {
	_ = "STUB: not implemented"
	return *new(ScopeProfiles)
}

func (es ScopeProfilesSlice) MoveAndAppendTo(dest ScopeProfilesSlice) {
	_ = "STUB: not implemented"
	return
}

func (es ScopeProfilesSlice) RemoveIf(f func(ScopeProfiles) bool) {
	_ = "STUB: not implemented"
	return
}

func (es ScopeProfilesSlice) CopyTo(dest ScopeProfilesSlice) { _ = "STUB: not implemented"; return }

func (es ScopeProfilesSlice) Sort(less func(a, b ScopeProfiles) bool) {
	_ = "STUB: not implemented"
	return
}
