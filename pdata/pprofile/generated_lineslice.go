package pprofile

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type LineSlice struct {
	orig  *[]*internal.Line
	state *internal.State
}

func newLineSlice(orig *[]*internal.Line, state *internal.State) LineSlice {
	_ = "STUB: not implemented"
	return *new(LineSlice)
}

func NewLineSlice() LineSlice { _ = "STUB: not implemented"; return *new(LineSlice) }

func (es LineSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es LineSlice) At(i int) Line { _ = "STUB: not implemented"; return *new(Line) }

func (es LineSlice) All() iter.Seq2[int, Line] { _ = "STUB: not implemented"; return nil }

func (es LineSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es LineSlice) AppendEmpty() Line { _ = "STUB: not implemented"; return *new(Line) }

func (es LineSlice) MoveAndAppendTo(dest LineSlice) { _ = "STUB: not implemented"; return }

func (es LineSlice) RemoveIf(f func(Line) bool) { _ = "STUB: not implemented"; return }

func (es LineSlice) CopyTo(dest LineSlice) { _ = "STUB: not implemented"; return }

func (es LineSlice) Sort(less func(a, b Line) bool) { _ = "STUB: not implemented"; return }
