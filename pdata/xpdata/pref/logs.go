package pref

import (
	"go.opentelemetry.io/collector/pdata/plog"
)

func MarkPipelineOwnedLogs(ld plog.Logs) bool { _ = "STUB: not implemented"; return false }

func RefLogs(ld plog.Logs) { _ = "STUB: not implemented"; return }

func UnrefLogs(ld plog.Logs) { _ = "STUB: not implemented"; return }

func EqualLogs(ld1, ld2 plog.Logs) bool { _ = "STUB: not implemented"; return false }
