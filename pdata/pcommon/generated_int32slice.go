package pcommon

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type Int32Slice internal.Int32SliceWrapper

func (ms Int32Slice) getOrig() *[]int32 { _ = "STUB: not implemented"; return nil }

func (ms Int32Slice) getState() *internal.State { _ = "STUB: not implemented"; return nil }

func NewInt32Slice() Int32Slice { _ = "STUB: not implemented"; return *new(Int32Slice) }

func (ms Int32Slice) AsRaw() []int32 { _ = "STUB: not implemented"; return nil }

func (ms Int32Slice) FromRaw(val []int32) { _ = "STUB: not implemented"; return }

func (ms Int32Slice) Len() int { _ = "STUB: not implemented"; return 0 }

func (ms Int32Slice) At(i int) int32 { _ = "STUB: not implemented"; return 0 }

func (ms Int32Slice) All() iter.Seq2[int, int32] { _ = "STUB: not implemented"; return nil }

func (ms Int32Slice) SetAt(i int, val int32) { _ = "STUB: not implemented"; return }

func (ms Int32Slice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (ms Int32Slice) Append(elms ...int32) { _ = "STUB: not implemented"; return }

func (ms Int32Slice) MoveTo(dest Int32Slice) { _ = "STUB: not implemented"; return }

func (ms Int32Slice) MoveAndAppendTo(dest Int32Slice) { _ = "STUB: not implemented"; return }

func (ms Int32Slice) RemoveIf(f func(int32) bool) { _ = "STUB: not implemented"; return }

func (ms Int32Slice) CopyTo(dest Int32Slice) { _ = "STUB: not implemented"; return }

func (ms Int32Slice) Equal(val Int32Slice) bool { _ = "STUB: not implemented"; return false }

func copyInt32Slice(dst, src []int32) []int32 { _ = "STUB: not implemented"; return nil }
