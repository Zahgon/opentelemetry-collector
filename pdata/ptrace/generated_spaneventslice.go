package ptrace

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type SpanEventSlice struct {
	orig  *[]*internal.SpanEvent
	state *internal.State
}

func newSpanEventSlice(orig *[]*internal.SpanEvent, state *internal.State) SpanEventSlice {
	_ = "STUB: not implemented"
	return *new(SpanEventSlice)
}

func NewSpanEventSlice() SpanEventSlice { _ = "STUB: not implemented"; return *new(SpanEventSlice) }

func (es SpanEventSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es SpanEventSlice) At(i int) SpanEvent { _ = "STUB: not implemented"; return *new(SpanEvent) }

func (es SpanEventSlice) All() iter.Seq2[int, SpanEvent] { _ = "STUB: not implemented"; return nil }

func (es SpanEventSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es SpanEventSlice) AppendEmpty() SpanEvent { _ = "STUB: not implemented"; return *new(SpanEvent) }

func (es SpanEventSlice) MoveAndAppendTo(dest SpanEventSlice) { _ = "STUB: not implemented"; return }

func (es SpanEventSlice) RemoveIf(f func(SpanEvent) bool) { _ = "STUB: not implemented"; return }

func (es SpanEventSlice) CopyTo(dest SpanEventSlice) { _ = "STUB: not implemented"; return }

func (es SpanEventSlice) Sort(less func(a, b SpanEvent) bool) { _ = "STUB: not implemented"; return }
