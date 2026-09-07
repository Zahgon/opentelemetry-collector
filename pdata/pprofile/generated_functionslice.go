package pprofile

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type FunctionSlice struct {
	orig  *[]*internal.Function
	state *internal.State
}

func newFunctionSlice(orig *[]*internal.Function, state *internal.State) FunctionSlice {
	_ = "STUB: not implemented"
	return *new(FunctionSlice)
}

func NewFunctionSlice() FunctionSlice { _ = "STUB: not implemented"; return *new(FunctionSlice) }

func (es FunctionSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es FunctionSlice) At(i int) Function { _ = "STUB: not implemented"; return *new(Function) }

func (es FunctionSlice) All() iter.Seq2[int, Function] { _ = "STUB: not implemented"; return nil }

func (es FunctionSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es FunctionSlice) AppendEmpty() Function { _ = "STUB: not implemented"; return *new(Function) }

func (es FunctionSlice) MoveAndAppendTo(dest FunctionSlice) { _ = "STUB: not implemented"; return }

func (es FunctionSlice) RemoveIf(f func(Function) bool) { _ = "STUB: not implemented"; return }

func (es FunctionSlice) CopyTo(dest FunctionSlice) { _ = "STUB: not implemented"; return }

func (es FunctionSlice) Sort(less func(a, b Function) bool) { _ = "STUB: not implemented"; return }
