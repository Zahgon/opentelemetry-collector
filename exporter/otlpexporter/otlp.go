package otlpexporter

import (
	"context"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pmetric/pmetricotlp"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/pprofile/pprofileotlp"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pdata/ptrace/ptraceotlp"
)

type baseExporter struct {
	config *Config

	traceExporter   ptraceotlp.GRPCClient
	metricExporter  pmetricotlp.GRPCClient
	logExporter     plogotlp.GRPCClient
	profileExporter pprofileotlp.GRPCClient
	clientConn      *grpc.ClientConn
	metadata        metadata.MD
	callOptions     []grpc.CallOption

	settings component.TelemetrySettings

	userAgent string
}

func newExporter(cfg component.Config, set exporter.Settings) *baseExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *baseExporter) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *baseExporter) shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *baseExporter) pushTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *baseExporter) pushMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *baseExporter) pushLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *baseExporter) pushProfiles(ctx context.Context, td pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

func processError(err error) error { _ = "STUB: not implemented"; return nil }

func shouldRetry(code codes.Code, retryInfo *errdetails.RetryInfo) bool {
	_ = "STUB: not implemented"
	return false
}
