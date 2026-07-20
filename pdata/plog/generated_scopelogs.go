package plog

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ScopeLogs struct {
	orig  *internal.ScopeLogs
	state *internal.State
}

func newScopeLogs(orig *internal.ScopeLogs, state *internal.State) ScopeLogs {
	_ = "STUB: not implemented"
	return *new(ScopeLogs)
}

func NewScopeLogs() ScopeLogs { _ = "STUB: not implemented"; return *new(ScopeLogs) }

func (ms ScopeLogs) MoveTo(dest ScopeLogs) { _ = "STUB: not implemented"; return }

func (ms ScopeLogs) Scope() pcommon.InstrumentationScope {
	_ = "STUB: not implemented"
	return *new(pcommon.InstrumentationScope)
}

func (ms ScopeLogs) LogRecords() LogRecordSlice {
	_ = "STUB: not implemented"
	return *new(LogRecordSlice)
}

func (ms ScopeLogs) SchemaUrl() string { _ = "STUB: not implemented"; return "" }

func (ms ScopeLogs) SetSchemaUrl(v string) { _ = "STUB: not implemented"; return }

func (ms ScopeLogs) CopyTo(dest ScopeLogs) { _ = "STUB: not implemented"; return }
