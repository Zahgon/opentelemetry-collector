package pprofile

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type SampleSlice struct {
	orig  *[]*internal.Sample
	state *internal.State
}

func newSampleSlice(orig *[]*internal.Sample, state *internal.State) SampleSlice {
	_ = "STUB: not implemented"
	return *new(SampleSlice)
}

func NewSampleSlice() SampleSlice { _ = "STUB: not implemented"; return *new(SampleSlice) }

func (es SampleSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es SampleSlice) At(i int) Sample { _ = "STUB: not implemented"; return *new(Sample) }

func (es SampleSlice) All() iter.Seq2[int, Sample] { _ = "STUB: not implemented"; return nil }

func (es SampleSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es SampleSlice) AppendEmpty() Sample { _ = "STUB: not implemented"; return *new(Sample) }

func (es SampleSlice) MoveAndAppendTo(dest SampleSlice) { _ = "STUB: not implemented"; return }

func (es SampleSlice) RemoveIf(f func(Sample) bool) { _ = "STUB: not implemented"; return }

func (es SampleSlice) CopyTo(dest SampleSlice) { _ = "STUB: not implemented"; return }

func (es SampleSlice) Sort(less func(a, b Sample) bool) { _ = "STUB: not implemented"; return }
