package plog

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ResourceLogs struct {
	orig  *internal.ResourceLogs
	state *internal.State
}

func newResourceLogs(orig *internal.ResourceLogs, state *internal.State) ResourceLogs {
	_ = "STUB: not implemented"
	return *new(ResourceLogs)
}

func NewResourceLogs() ResourceLogs { _ = "STUB: not implemented"; return *new(ResourceLogs) }

func (ms ResourceLogs) MoveTo(dest ResourceLogs) { _ = "STUB: not implemented"; return }

func (ms ResourceLogs) Resource() pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}

func (ms ResourceLogs) ScopeLogs() ScopeLogsSlice {
	_ = "STUB: not implemented"
	return *new(ScopeLogsSlice)
}

func (ms ResourceLogs) SchemaUrl() string { _ = "STUB: not implemented"; return "" }

func (ms ResourceLogs) SetSchemaUrl(v string) { _ = "STUB: not implemented"; return }

func (ms ResourceLogs) CopyTo(dest ResourceLogs) { _ = "STUB: not implemented"; return }
