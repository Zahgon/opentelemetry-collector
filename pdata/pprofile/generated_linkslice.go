package pprofile

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type LinkSlice struct {
	orig  *[]*internal.Link
	state *internal.State
}

func newLinkSlice(orig *[]*internal.Link, state *internal.State) LinkSlice {
	_ = "STUB: not implemented"
	return *new(LinkSlice)
}

func NewLinkSlice() LinkSlice { _ = "STUB: not implemented"; return *new(LinkSlice) }

func (es LinkSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es LinkSlice) At(i int) Link { _ = "STUB: not implemented"; return *new(Link) }

func (es LinkSlice) All() iter.Seq2[int, Link] { _ = "STUB: not implemented"; return nil }

func (es LinkSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es LinkSlice) AppendEmpty() Link { _ = "STUB: not implemented"; return *new(Link) }

func (es LinkSlice) MoveAndAppendTo(dest LinkSlice) { _ = "STUB: not implemented"; return }

func (es LinkSlice) RemoveIf(f func(Link) bool) { _ = "STUB: not implemented"; return }

func (es LinkSlice) CopyTo(dest LinkSlice) { _ = "STUB: not implemented"; return }

func (es LinkSlice) Sort(less func(a, b Link) bool) { _ = "STUB: not implemented"; return }
