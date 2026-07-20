package batchprocessor

import (
	"context"

	"go.opentelemetry.io/otel/metric"

	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/batchprocessor/internal/metadata"
)

type trigger int

const (
	triggerTimeout trigger = iota
	triggerBatchSize
)

type batchProcessorTelemetry struct {
	exportCtx context.Context

	processorAttr    metric.MeasurementOption
	telemetryBuilder *metadata.TelemetryBuilder
}

func newBatchProcessorTelemetry(set processor.Settings, currentMetadataCardinality func() int) (*batchProcessorTelemetry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bpt *batchProcessorTelemetry) record(trigger trigger, sent, bytes int64) {
	_ = "STUB: not implemented"
	return
}
