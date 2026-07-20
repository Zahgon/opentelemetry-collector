package consumererror

import (
	"go.opentelemetry.io/collector/consumer/consumererror/internal"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type Traces struct {
	internal.Retryable[ptrace.Traces]
}

func NewTraces(err error, data ptrace.Traces) error { _ = "STUB: not implemented"; return nil }

type Logs struct {
	internal.Retryable[plog.Logs]
}

func NewLogs(err error, data plog.Logs) error { _ = "STUB: not implemented"; return nil }

type Metrics struct {
	internal.Retryable[pmetric.Metrics]
}

func NewMetrics(err error, data pmetric.Metrics) error { _ = "STUB: not implemented"; return nil }
