//go:generate mdatagen metadata.yaml

package receiverhelper

import (
	"context"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper/internal/metadata"
)

type ObsReport struct {
	spanNamePrefix string
	transport      string
	longLivedCtx   bool
	tracer         trace.Tracer

	otelAttrs        metric.MeasurementOption
	telemetryBuilder *metadata.TelemetryBuilder
}

type ObsReportSettings struct {
	ReceiverID component.ID
	Transport  string

	LongLivedCtx           bool
	ReceiverCreateSettings receiver.Settings

	_ struct{}
}

func NewObsReport(cfg ObsReportSettings) (*ObsReport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newReceiver(cfg ObsReportSettings) (*ObsReport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rec *ObsReport) StartTracesOp(operationCtx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (rec *ObsReport) EndTracesOp(
	receiverCtx context.Context,
	format string,
	numReceivedSpans int,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

func (rec *ObsReport) StartLogsOp(operationCtx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (rec *ObsReport) EndLogsOp(
	receiverCtx context.Context,
	format string,
	numReceivedLogRecords int,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

func (rec *ObsReport) StartMetricsOp(operationCtx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (rec *ObsReport) EndMetricsOp(
	receiverCtx context.Context,
	format string,
	numReceivedPoints int,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

func (rec *ObsReport) StartProfilesOp(operationCtx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (rec *ObsReport) EndProfilesOp(
	receiverCtx context.Context,
	format string,
	numReceivedProfileSamples int,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

func (rec *ObsReport) startOp(receiverCtx context.Context, operationSuffix string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (rec *ObsReport) endOp(
	receiverCtx context.Context,
	format string,
	numReceivedItems int,
	err error,
	signal pipeline.Signal,
) {
	_ = "STUB: not implemented"
	return
}

func (rec *ObsReport) recordMetrics(receiverCtx context.Context, signal pipeline.Signal, numAccepted, numRefused, numFailedErrors int) {
	_ = "STUB: not implemented"
	return
}
