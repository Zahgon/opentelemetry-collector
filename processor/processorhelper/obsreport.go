package processorhelper

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/metric"

	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper/internal/metadata"
)

const signalKey = "otel.signal"

type obsReport struct {
	otelAttrs        metric.MeasurementOption
	telemetryBuilder *metadata.TelemetryBuilder
}

func newObsReport(set processor.Settings, signal pipeline.Signal) (*obsReport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (or *obsReport) recordInOut(ctx context.Context, incoming, outgoing int) {
	_ = "STUB: not implemented"
	return
}

func (or *obsReport) recordInternalDuration(ctx context.Context, startTime time.Time) {
	_ = "STUB: not implemented"
	return
}
