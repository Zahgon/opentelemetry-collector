package plog

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type ResourceLogsSlice struct {
	orig  *[]*internal.ResourceLogs
	state *internal.State
}

func newResourceLogsSlice(orig *[]*internal.ResourceLogs, state *internal.State) ResourceLogsSlice {
	_ = "STUB: not implemented"
	return *new(ResourceLogsSlice)
}

func NewResourceLogsSlice() ResourceLogsSlice {
	_ = "STUB: not implemented"
	return *new(ResourceLogsSlice)
}

func (es ResourceLogsSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es ResourceLogsSlice) At(i int) ResourceLogs {
	_ = "STUB: not implemented"
	return *new(ResourceLogs)
}

func (es ResourceLogsSlice) All() iter.Seq2[int, ResourceLogs] {
	_ = "STUB: not implemented"
	return nil
}

func (es ResourceLogsSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es ResourceLogsSlice) AppendEmpty() ResourceLogs {
	_ = "STUB: not implemented"
	return *new(ResourceLogs)
}

func (es ResourceLogsSlice) MoveAndAppendTo(dest ResourceLogsSlice) {
	_ = "STUB: not implemented"
	return
}

func (es ResourceLogsSlice) RemoveIf(f func(ResourceLogs) bool) { _ = "STUB: not implemented"; return }

func (es ResourceLogsSlice) CopyTo(dest ResourceLogsSlice) { _ = "STUB: not implemented"; return }

func (es ResourceLogsSlice) Sort(less func(a, b ResourceLogs) bool) {
	_ = "STUB: not implemented"
	return
}
