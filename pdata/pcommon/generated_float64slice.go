package pcommon

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type Float64Slice internal.Float64SliceWrapper

func (ms Float64Slice) getOrig() *[]float64 { _ = "STUB: not implemented"; return nil }

func (ms Float64Slice) getState() *internal.State { _ = "STUB: not implemented"; return nil }

func NewFloat64Slice() Float64Slice { _ = "STUB: not implemented"; return *new(Float64Slice) }

func (ms Float64Slice) AsRaw() []float64 { _ = "STUB: not implemented"; return nil }

func (ms Float64Slice) FromRaw(val []float64) { _ = "STUB: not implemented"; return }

func (ms Float64Slice) Len() int { _ = "STUB: not implemented"; return 0 }

func (ms Float64Slice) At(i int) float64 { _ = "STUB: not implemented"; return 0 }

func (ms Float64Slice) All() iter.Seq2[int, float64] { _ = "STUB: not implemented"; return nil }

func (ms Float64Slice) SetAt(i int, val float64) { _ = "STUB: not implemented"; return }

func (ms Float64Slice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (ms Float64Slice) Append(elms ...float64) { _ = "STUB: not implemented"; return }

func (ms Float64Slice) MoveTo(dest Float64Slice) { _ = "STUB: not implemented"; return }

func (ms Float64Slice) MoveAndAppendTo(dest Float64Slice) { _ = "STUB: not implemented"; return }

func (ms Float64Slice) RemoveIf(f func(float64) bool) { _ = "STUB: not implemented"; return }

func (ms Float64Slice) CopyTo(dest Float64Slice) { _ = "STUB: not implemented"; return }

func (ms Float64Slice) Equal(val Float64Slice) bool { _ = "STUB: not implemented"; return false }

func copyFloat64Slice(dst, src []float64) []float64 { _ = "STUB: not implemented"; return nil }
