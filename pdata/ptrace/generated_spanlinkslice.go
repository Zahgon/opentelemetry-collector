package ptrace

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type SpanLinkSlice struct {
	orig  *[]*internal.SpanLink
	state *internal.State
}

func newSpanLinkSlice(orig *[]*internal.SpanLink, state *internal.State) SpanLinkSlice {
	_ = "STUB: not implemented"
	return *new(SpanLinkSlice)
}

func NewSpanLinkSlice() SpanLinkSlice { _ = "STUB: not implemented"; return *new(SpanLinkSlice) }

func (es SpanLinkSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es SpanLinkSlice) At(i int) SpanLink { _ = "STUB: not implemented"; return *new(SpanLink) }

func (es SpanLinkSlice) All() iter.Seq2[int, SpanLink] { _ = "STUB: not implemented"; return nil }

func (es SpanLinkSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es SpanLinkSlice) AppendEmpty() SpanLink { _ = "STUB: not implemented"; return *new(SpanLink) }

func (es SpanLinkSlice) MoveAndAppendTo(dest SpanLinkSlice) { _ = "STUB: not implemented"; return }

func (es SpanLinkSlice) RemoveIf(f func(SpanLink) bool) { _ = "STUB: not implemented"; return }

func (es SpanLinkSlice) CopyTo(dest SpanLinkSlice) { _ = "STUB: not implemented"; return }

func (es SpanLinkSlice) Sort(less func(a, b SpanLink) bool) { _ = "STUB: not implemented"; return }
