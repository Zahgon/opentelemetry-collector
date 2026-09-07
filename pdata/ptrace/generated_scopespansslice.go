package ptrace

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type ScopeSpansSlice struct {
	orig  *[]*internal.ScopeSpans
	state *internal.State
}

func newScopeSpansSlice(orig *[]*internal.ScopeSpans, state *internal.State) ScopeSpansSlice {
	_ = "STUB: not implemented"
	return *new(ScopeSpansSlice)
}

func NewScopeSpansSlice() ScopeSpansSlice { _ = "STUB: not implemented"; return *new(ScopeSpansSlice) }

func (es ScopeSpansSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es ScopeSpansSlice) At(i int) ScopeSpans { _ = "STUB: not implemented"; return *new(ScopeSpans) }

func (es ScopeSpansSlice) All() iter.Seq2[int, ScopeSpans] { _ = "STUB: not implemented"; return nil }

func (es ScopeSpansSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es ScopeSpansSlice) AppendEmpty() ScopeSpans {
	_ = "STUB: not implemented"
	return *new(ScopeSpans)
}

func (es ScopeSpansSlice) MoveAndAppendTo(dest ScopeSpansSlice) { _ = "STUB: not implemented"; return }

func (es ScopeSpansSlice) RemoveIf(f func(ScopeSpans) bool) { _ = "STUB: not implemented"; return }

func (es ScopeSpansSlice) CopyTo(dest ScopeSpansSlice) { _ = "STUB: not implemented"; return }

func (es ScopeSpansSlice) Sort(less func(a, b ScopeSpans) bool) { _ = "STUB: not implemented"; return }
