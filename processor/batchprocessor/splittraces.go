package batchprocessor

import (
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func splitTraces(size int, src ptrace.Traces) ptrace.Traces {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces)
}

func resourceSC(rs ptrace.ResourceSpans) (count int) { _ = "STUB: not implemented"; return 0 }
