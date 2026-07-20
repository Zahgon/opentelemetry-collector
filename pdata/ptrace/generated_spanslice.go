package ptrace

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type SpanSlice struct {
	orig  *[]*internal.Span
	state *internal.State
}

func newSpanSlice(orig *[]*internal.Span, state *internal.State) SpanSlice {
	_ = "STUB: not implemented"
	return *new(SpanSlice)
}

func NewSpanSlice() SpanSlice { _ = "STUB: not implemented"; return *new(SpanSlice) }

func (es SpanSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es SpanSlice) At(i int) Span { _ = "STUB: not implemented"; return *new(Span) }

func (es SpanSlice) All() iter.Seq2[int, Span] { _ = "STUB: not implemented"; return nil }

func (es SpanSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es SpanSlice) AppendEmpty() Span { _ = "STUB: not implemented"; return *new(Span) }

func (es SpanSlice) MoveAndAppendTo(dest SpanSlice) { _ = "STUB: not implemented"; return }

func (es SpanSlice) RemoveIf(f func(Span) bool) { _ = "STUB: not implemented"; return }

func (es SpanSlice) CopyTo(dest SpanSlice) { _ = "STUB: not implemented"; return }

func (es SpanSlice) Sort(less func(a, b Span) bool) { _ = "STUB: not implemented"; return }
