package pcommon

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type Slice internal.SliceWrapper

func newSlice(orig *[]internal.AnyValue, state *internal.State) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

func NewSlice() Slice { _ = "STUB: not implemented"; return *new(Slice) }

func (es Slice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es Slice) At(i int) Value { _ = "STUB: not implemented"; return *new(Value) }

func (es Slice) All() iter.Seq2[int, Value] { _ = "STUB: not implemented"; return nil }

func (es Slice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es Slice) AppendEmpty() Value { _ = "STUB: not implemented"; return *new(Value) }

func (es Slice) MoveAndAppendTo(dest Slice) { _ = "STUB: not implemented"; return }

func (es Slice) RemoveIf(f func(Value) bool) { _ = "STUB: not implemented"; return }

func (es Slice) CopyTo(dest Slice) { _ = "STUB: not implemented"; return }

func (ms Slice) getOrig() *[]internal.AnyValue { _ = "STUB: not implemented"; return nil }

func (ms Slice) getState() *internal.State { _ = "STUB: not implemented"; return nil }
