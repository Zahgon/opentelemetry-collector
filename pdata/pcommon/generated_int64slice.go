package pcommon

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type Int64Slice internal.Int64SliceWrapper

func (ms Int64Slice) getOrig() *[]int64 { _ = "STUB: not implemented"; return nil }

func (ms Int64Slice) getState() *internal.State { _ = "STUB: not implemented"; return nil }

func NewInt64Slice() Int64Slice { _ = "STUB: not implemented"; return *new(Int64Slice) }

func (ms Int64Slice) AsRaw() []int64 { _ = "STUB: not implemented"; return nil }

func (ms Int64Slice) FromRaw(val []int64) { _ = "STUB: not implemented"; return }

func (ms Int64Slice) Len() int { _ = "STUB: not implemented"; return 0 }

func (ms Int64Slice) At(i int) int64 { _ = "STUB: not implemented"; return 0 }

func (ms Int64Slice) All() iter.Seq2[int, int64] { _ = "STUB: not implemented"; return nil }

func (ms Int64Slice) SetAt(i int, val int64) { _ = "STUB: not implemented"; return }

func (ms Int64Slice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (ms Int64Slice) Append(elms ...int64) { _ = "STUB: not implemented"; return }

func (ms Int64Slice) MoveTo(dest Int64Slice) { _ = "STUB: not implemented"; return }

func (ms Int64Slice) MoveAndAppendTo(dest Int64Slice) { _ = "STUB: not implemented"; return }

func (ms Int64Slice) RemoveIf(f func(int64) bool) { _ = "STUB: not implemented"; return }

func (ms Int64Slice) CopyTo(dest Int64Slice) { _ = "STUB: not implemented"; return }

func (ms Int64Slice) Equal(val Int64Slice) bool { _ = "STUB: not implemented"; return false }

func copyInt64Slice(dst, src []int64) []int64 { _ = "STUB: not implemented"; return nil }
