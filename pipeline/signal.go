package pipeline

import (
	"errors"

	"go.opentelemetry.io/collector/pipeline/internal/globalsignal"
)

type Signal = globalsignal.Signal

var ErrSignalNotSupported = errors.New("telemetry type is not supported")

var (
	SignalTraces  = globalsignal.SignalTraces
	SignalMetrics = globalsignal.SignalMetrics
	SignalLogs    = globalsignal.SignalLogs
)
