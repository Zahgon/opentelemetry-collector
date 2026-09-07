package pmetric

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type ExemplarSlice struct {
	orig  *[]internal.Exemplar
	state *internal.State
}

func newExemplarSlice(orig *[]internal.Exemplar, state *internal.State) ExemplarSlice {
	_ = "STUB: not implemented"
	return *new(ExemplarSlice)
}

func NewExemplarSlice() ExemplarSlice { _ = "STUB: not implemented"; return *new(ExemplarSlice) }

func (es ExemplarSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es ExemplarSlice) At(i int) Exemplar { _ = "STUB: not implemented"; return *new(Exemplar) }

func (es ExemplarSlice) All() iter.Seq2[int, Exemplar] { _ = "STUB: not implemented"; return nil }

func (es ExemplarSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es ExemplarSlice) AppendEmpty() Exemplar { _ = "STUB: not implemented"; return *new(Exemplar) }

func (es ExemplarSlice) MoveAndAppendTo(dest ExemplarSlice) { _ = "STUB: not implemented"; return }

func (es ExemplarSlice) RemoveIf(f func(Exemplar) bool) { _ = "STUB: not implemented"; return }

func (es ExemplarSlice) CopyTo(dest ExemplarSlice) { _ = "STUB: not implemented"; return }
