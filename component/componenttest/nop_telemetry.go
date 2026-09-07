package componenttest

import (
	"go.opentelemetry.io/collector/component"
)

func NewNopTelemetrySettings() component.TelemetrySettings {
	_ = "STUB: not implemented"
	return *new(component.TelemetrySettings)
}
