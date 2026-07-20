package plog

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type LogRecordSlice struct {
	orig  *[]*internal.LogRecord
	state *internal.State
}

func newLogRecordSlice(orig *[]*internal.LogRecord, state *internal.State) LogRecordSlice {
	_ = "STUB: not implemented"
	return *new(LogRecordSlice)
}

func NewLogRecordSlice() LogRecordSlice { _ = "STUB: not implemented"; return *new(LogRecordSlice) }

func (es LogRecordSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (es LogRecordSlice) At(i int) LogRecord { _ = "STUB: not implemented"; return *new(LogRecord) }

func (es LogRecordSlice) All() iter.Seq2[int, LogRecord] { _ = "STUB: not implemented"; return nil }

func (es LogRecordSlice) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (es LogRecordSlice) AppendEmpty() LogRecord { _ = "STUB: not implemented"; return *new(LogRecord) }

func (es LogRecordSlice) MoveAndAppendTo(dest LogRecordSlice) { _ = "STUB: not implemented"; return }

func (es LogRecordSlice) RemoveIf(f func(LogRecord) bool) { _ = "STUB: not implemented"; return }

func (es LogRecordSlice) CopyTo(dest LogRecordSlice) { _ = "STUB: not implemented"; return }

func (es LogRecordSlice) Sort(less func(a, b LogRecord) bool) { _ = "STUB: not implemented"; return }
