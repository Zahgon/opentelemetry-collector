package pcommon

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type StringSlice internal.StringSliceWrapper

func (ms StringSlice) getOrig() *[]string { _ = "STUB: not implemented"; return nil }

func (ms StringSlice) getState() *internal.State { _ = "STUB: not implemented"; return nil }

func NewStringSlice() StringSlice { _ = "STUB: not implemented"; return *new(StringSlice) }

func (ms StringSlice) AsRaw() []string { _ = "STUB: not implemented"; return nil }

func (ms StringSlice) FromRaw(val []string) { _ = "STUB: not implemented"; return }

func (ms StringSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (ms StringSlice) At(i int) string { _ = "STUB: not implemented"; return "" }

func (ms StringSlice) All() iter.Seq2[int, string] { _ = "STUB: not implemented"; return nil }

func (ms StringSlice) SetAt(i int, val string) { _ = "STUB: not implemented"; return }

func (ms StringSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (ms StringSlice) Append(elms ...string) { _ = "STUB: not implemented"; return }

func (ms StringSlice) MoveTo(dest StringSlice) { _ = "STUB: not implemented"; return }

func (ms StringSlice) MoveAndAppendTo(dest StringSlice) { _ = "STUB: not implemented"; return }

func (ms StringSlice) RemoveIf(f func(string) bool) { _ = "STUB: not implemented"; return }

func (ms StringSlice) CopyTo(dest StringSlice) { _ = "STUB: not implemented"; return }

func (ms StringSlice) Equal(val StringSlice) bool { _ = "STUB: not implemented"; return false }

func copyStringSlice(dst, src []string) []string { _ = "STUB: not implemented"; return nil }
