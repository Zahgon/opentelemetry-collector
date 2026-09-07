package exporterhelper

import (
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal"
)

type TimeoutConfig = internal.TimeoutConfig

func NewDefaultTimeoutConfig() TimeoutConfig { _ = "STUB: not implemented"; return *new(TimeoutConfig) }
