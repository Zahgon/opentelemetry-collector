package batchprocessor

import (
	"go.opentelemetry.io/collector/pdata/plog"
)

func splitLogs(size int, src plog.Logs) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

func resourceLRC(rs plog.ResourceLogs) (count int) { _ = "STUB: not implemented"; return 0 }
