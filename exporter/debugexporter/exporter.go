package debugexporter

import (
	"context"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type debugExporter struct {
	verbosity         configtelemetry.Level
	logger            *zap.Logger
	logsMarshaler     plog.Marshaler
	metricsMarshaler  pmetric.Marshaler
	tracesMarshaler   ptrace.Marshaler
	profilesMarshaler pprofile.Marshaler
}

func newDebugExporter(logger *zap.Logger, verbosity configtelemetry.Level) *debugExporter {
	_ = "STUB: not implemented"
	return nil
}

func (s *debugExporter) pushTraces(_ context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *debugExporter) pushMetrics(_ context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *debugExporter) pushLogs(_ context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *debugExporter) pushProfiles(_ context.Context, pd pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}
