package memorylimiterprocessor

import (
	"context"

	"go.opentelemetry.io/otel/metric"

	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/memorylimiterprocessor/internal/metadata"
)

type obsReport struct {
	otelAttrs        metric.MeasurementOption
	telemetryBuilder *metadata.TelemetryBuilder
}

func newObsReport(set processor.Settings) (*obsReport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (or *obsReport) accepted(ctx context.Context, num int, signal pipeline.Signal) {
	_ = "STUB: not implemented"
	return
}

func (or *obsReport) refused(ctx context.Context, num int, signal pipeline.Signal) {
	_ = "STUB: not implemented"
	return
}
