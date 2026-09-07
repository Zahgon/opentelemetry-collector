package ptrace

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type ResourceSpansSlice struct {
	orig  *[]*internal.ResourceSpans
	state *internal.State
}

func newResourceSpansSlice(orig *[]*internal.ResourceSpans, state *internal.State) ResourceSpansSlice {
	_ = "STUB: not implemented"
	return *new(ResourceSpansSlice)
}

func NewResourceSpansSlice() ResourceSpansSlice {
	_ = "STUB: not implemented"
	return *new(ResourceSpansSlice)
}

func (es ResourceSpansSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es ResourceSpansSlice) At(i int) ResourceSpans {
	_ = "STUB: not implemented"
	return *new(ResourceSpans)
}

func (es ResourceSpansSlice) All() iter.Seq2[int, ResourceSpans] {
	_ = "STUB: not implemented"
	return nil
}

func (es ResourceSpansSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es ResourceSpansSlice) AppendEmpty() ResourceSpans {
	_ = "STUB: not implemented"
	return *new(ResourceSpans)
}

func (es ResourceSpansSlice) MoveAndAppendTo(dest ResourceSpansSlice) {
	_ = "STUB: not implemented"
	return
}

func (es ResourceSpansSlice) RemoveIf(f func(ResourceSpans) bool) {
	_ = "STUB: not implemented"
	return
}

func (es ResourceSpansSlice) CopyTo(dest ResourceSpansSlice) { _ = "STUB: not implemented"; return }

func (es ResourceSpansSlice) Sort(less func(a, b ResourceSpans) bool) {
	_ = "STUB: not implemented"
	return
}
