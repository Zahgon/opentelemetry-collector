package pcommon

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type ByteSlice internal.ByteSliceWrapper

func (ms ByteSlice) getOrig() *[]byte { _ = "STUB: not implemented"; return nil }

func (ms ByteSlice) getState() *internal.State { _ = "STUB: not implemented"; return nil }

func NewByteSlice() ByteSlice { _ = "STUB: not implemented"; return *new(ByteSlice) }

func (ms ByteSlice) AsRaw() []byte { _ = "STUB: not implemented"; return nil }

func (ms ByteSlice) FromRaw(val []byte) { _ = "STUB: not implemented"; return }

func (ms ByteSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (ms ByteSlice) At(i int) byte { _ = "STUB: not implemented"; return 0 }

func (ms ByteSlice) All() iter.Seq2[int, byte] { _ = "STUB: not implemented"; return nil }

func (ms ByteSlice) SetAt(i int, val byte) { _ = "STUB: not implemented"; return }

func (ms ByteSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (ms ByteSlice) Append(elms ...byte) { _ = "STUB: not implemented"; return }

func (ms ByteSlice) MoveTo(dest ByteSlice) { _ = "STUB: not implemented"; return }

func (ms ByteSlice) MoveAndAppendTo(dest ByteSlice) { _ = "STUB: not implemented"; return }

func (ms ByteSlice) RemoveIf(f func(byte) bool) { _ = "STUB: not implemented"; return }

func (ms ByteSlice) CopyTo(dest ByteSlice) { _ = "STUB: not implemented"; return }

func (ms ByteSlice) Equal(val ByteSlice) bool { _ = "STUB: not implemented"; return false }

func copyByteSlice(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }
