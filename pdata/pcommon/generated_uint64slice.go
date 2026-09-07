package pcommon

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type UInt64Slice internal.UInt64SliceWrapper

func (ms UInt64Slice) getOrig() *[]uint64 { _ = "STUB: not implemented"; return nil }

func (ms UInt64Slice) getState() *internal.State { _ = "STUB: not implemented"; return nil }

func NewUInt64Slice() UInt64Slice { _ = "STUB: not implemented"; return *new(UInt64Slice) }

func (ms UInt64Slice) AsRaw() []uint64 { _ = "STUB: not implemented"; return nil }

func (ms UInt64Slice) FromRaw(val []uint64) { _ = "STUB: not implemented"; return }

func (ms UInt64Slice) Len() int { _ = "STUB: not implemented"; return 0 }

func (ms UInt64Slice) At(i int) uint64 { _ = "STUB: not implemented"; return 0 }

func (ms UInt64Slice) All() iter.Seq2[int, uint64] { _ = "STUB: not implemented"; return nil }

func (ms UInt64Slice) SetAt(i int, val uint64) { _ = "STUB: not implemented"; return }

func (ms UInt64Slice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (ms UInt64Slice) Append(elms ...uint64) { _ = "STUB: not implemented"; return }

func (ms UInt64Slice) MoveTo(dest UInt64Slice) { _ = "STUB: not implemented"; return }

func (ms UInt64Slice) MoveAndAppendTo(dest UInt64Slice) { _ = "STUB: not implemented"; return }

func (ms UInt64Slice) RemoveIf(f func(uint64) bool) { _ = "STUB: not implemented"; return }

func (ms UInt64Slice) CopyTo(dest UInt64Slice) { _ = "STUB: not implemented"; return }

func (ms UInt64Slice) Equal(val UInt64Slice) bool { _ = "STUB: not implemented"; return false }

func copyUint64Slice(dst, src []uint64) []uint64 { _ = "STUB: not implemented"; return nil }
