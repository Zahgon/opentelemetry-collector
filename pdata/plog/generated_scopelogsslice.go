package plog

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type ScopeLogsSlice struct {
	orig  *[]*internal.ScopeLogs
	state *internal.State
}

func newScopeLogsSlice(orig *[]*internal.ScopeLogs, state *internal.State) ScopeLogsSlice {
	_ = "STUB: not implemented"
	return *new(ScopeLogsSlice)
}

func NewScopeLogsSlice() ScopeLogsSlice { _ = "STUB: not implemented"; return *new(ScopeLogsSlice) }

func (es ScopeLogsSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es ScopeLogsSlice) At(i int) ScopeLogs { _ = "STUB: not implemented"; return *new(ScopeLogs) }

func (es ScopeLogsSlice) All() iter.Seq2[int, ScopeLogs] { _ = "STUB: not implemented"; return nil }

func (es ScopeLogsSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es ScopeLogsSlice) AppendEmpty() ScopeLogs { _ = "STUB: not implemented"; return *new(ScopeLogs) }

func (es ScopeLogsSlice) MoveAndAppendTo(dest ScopeLogsSlice) { _ = "STUB: not implemented"; return }

func (es ScopeLogsSlice) RemoveIf(f func(ScopeLogs) bool) { _ = "STUB: not implemented"; return }

func (es ScopeLogsSlice) CopyTo(dest ScopeLogsSlice) { _ = "STUB: not implemented"; return }

func (es ScopeLogsSlice) Sort(less func(a, b ScopeLogs) bool) { _ = "STUB: not implemented"; return }
