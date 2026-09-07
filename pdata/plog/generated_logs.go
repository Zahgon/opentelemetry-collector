package plog

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type Logs internal.LogsWrapper

func newLogs(orig *internal.ExportLogsServiceRequest, state *internal.State) Logs {
	_ = "STUB: not implemented"
	return *new(Logs)
}

func NewLogs() Logs { _ = "STUB: not implemented"; return *new(Logs) }

func (ms Logs) MoveTo(dest Logs) { _ = "STUB: not implemented"; return }

func (ms Logs) ResourceLogs() ResourceLogsSlice {
	_ = "STUB: not implemented"
	return *new(ResourceLogsSlice)
}

func (ms Logs) CopyTo(dest Logs) { _ = "STUB: not implemented"; return }

func (ms Logs) getOrig() *internal.ExportLogsServiceRequest { _ = "STUB: not implemented"; return nil }

func (ms Logs) getState() *internal.State { _ = "STUB: not implemented"; return nil }
