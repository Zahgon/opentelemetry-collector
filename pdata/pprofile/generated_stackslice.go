package pprofile

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type StackSlice struct {
	orig  *[]*internal.Stack
	state *internal.State
}

func newStackSlice(orig *[]*internal.Stack, state *internal.State) StackSlice {
	_ = "STUB: not implemented"
	return *new(StackSlice)
}

func NewStackSlice() StackSlice { _ = "STUB: not implemented"; return *new(StackSlice) }

func (es StackSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es StackSlice) At(i int) Stack { _ = "STUB: not implemented"; return *new(Stack) }

func (es StackSlice) All() iter.Seq2[int, Stack] { _ = "STUB: not implemented"; return nil }

func (es StackSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es StackSlice) AppendEmpty() Stack { _ = "STUB: not implemented"; return *new(Stack) }

func (es StackSlice) MoveAndAppendTo(dest StackSlice) { _ = "STUB: not implemented"; return }

func (es StackSlice) RemoveIf(f func(Stack) bool) { _ = "STUB: not implemented"; return }

func (es StackSlice) CopyTo(dest StackSlice) { _ = "STUB: not implemented"; return }

func (es StackSlice) Sort(less func(a, b Stack) bool) { _ = "STUB: not implemented"; return }
