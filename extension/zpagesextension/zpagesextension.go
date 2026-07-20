package zpagesextension

import (
	"context"
	"net/http"

	"go.opentelemetry.io/contrib/zpages"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/collector/component"
)

const (
	tracezPath  = "tracez"
	expvarzPath = "expvarz"
)

type zpagesExtension struct {
	config              *Config
	telemetry           component.TelemetrySettings
	zpagesSpanProcessor *zpages.SpanProcessor
	server              *http.Server
	stopCh              chan struct{}
}

type registerableTracerProvider interface {
	RegisterSpanProcessor(SpanProcessor traceSdk.SpanProcessor)

	UnregisterSpanProcessor(SpanProcessor traceSdk.SpanProcessor)
}

func (zpe *zpagesExtension) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (zpe *zpagesExtension) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func newServer(config *Config, telemetry component.TelemetrySettings) *zpagesExtension {
	_ = "STUB: not implemented"
	return nil
}
