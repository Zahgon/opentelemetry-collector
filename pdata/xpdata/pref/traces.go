package pref

import (
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func MarkPipelineOwnedTraces(td ptrace.Traces) bool { _ = "STUB: not implemented"; return false }

func RefTraces(td ptrace.Traces) { _ = "STUB: not implemented"; return }

func UnrefTraces(td ptrace.Traces) { _ = "STUB: not implemented"; return }

func EqualTraces(td1, td2 ptrace.Traces) bool { _ = "STUB: not implemented"; return false }
