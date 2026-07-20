package otelconftelemetry

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
	"go.opentelemetry.io/otel/trace/noop"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service/telemetry"
)

const (
	traceContextPropagator = "tracecontext"
	b3Propagator           = "b3"
)

func createTracerProvider(
	ctx context.Context,
	set telemetry.TracerSettings,
	componentConfig component.Config,
) (telemetry.TracerProvider, error) {
	_ = "STUB: not implemented"
	return *new(telemetry.TracerProvider), nil
}

var errUnsupportedPropagator = errors.New("unsupported trace propagator")

type noopNoContextTracer struct {
	embedded.Tracer
}

var noopSpan = noop.Span{}

func (n *noopNoContextTracer) Start(ctx context.Context, _ string, _ ...trace.SpanStartOption) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

type noopNoContextTracerProvider struct {
	embedded.TracerProvider
}

func (n *noopNoContextTracerProvider) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *noopNoContextTracerProvider) Tracer(_ string, _ ...trace.TracerOption) trace.Tracer {
	_ = "STUB: not implemented"
	return *new(trace.Tracer)
}

func textMapPropagatorFromConfig(props []string) (propagation.TextMapPropagator, error) {
	_ = "STUB: not implemented"
	return *new(propagation.TextMapPropagator), nil
}
