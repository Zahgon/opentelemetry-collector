package entity

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type EntityRefSlice internal.EntityRefSliceWrapper

func newEntityRefSlice(orig *[]*internal.EntityRef, state *internal.State) EntityRefSlice {
	_ = "STUB: not implemented"
	return *new(EntityRefSlice)
}

func NewEntityRefSlice() EntityRefSlice { _ = "STUB: not implemented"; return *new(EntityRefSlice) }

func (es EntityRefSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es EntityRefSlice) At(i int) EntityRef { _ = "STUB: not implemented"; return *new(EntityRef) }

func (es EntityRefSlice) All() iter.Seq2[int, EntityRef] { _ = "STUB: not implemented"; return nil }

func (es EntityRefSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es EntityRefSlice) AppendEmpty() EntityRef { _ = "STUB: not implemented"; return *new(EntityRef) }

func (es EntityRefSlice) MoveAndAppendTo(dest EntityRefSlice) { _ = "STUB: not implemented"; return }

func (es EntityRefSlice) RemoveIf(f func(EntityRef) bool) { _ = "STUB: not implemented"; return }

func (es EntityRefSlice) CopyTo(dest EntityRefSlice) { _ = "STUB: not implemented"; return }

func (es EntityRefSlice) Sort(less func(a, b EntityRef) bool) { _ = "STUB: not implemented"; return }

func (ms EntityRefSlice) getOrig() *[]*internal.EntityRef { _ = "STUB: not implemented"; return nil }

func (ms EntityRefSlice) getState() *internal.State { _ = "STUB: not implemented"; return nil }
