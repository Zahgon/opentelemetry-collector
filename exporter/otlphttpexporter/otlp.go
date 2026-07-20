package otlphttpexporter

import (
	"context"
	"net/http"

	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/rpc/status"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type baseExporter struct {
	config      *Config
	client      *http.Client
	tracesURL   string
	metricsURL  string
	logsURL     string
	profilesURL string
	logger      *zap.Logger
	settings    component.TelemetrySettings

	userAgent string
}

const (
	headerRetryAfter         = "Retry-After"
	maxHTTPResponseReadBytes = 64 * 1024

	jsonContentType     = "application/json"
	protobufContentType = "application/x-protobuf"
)

func newExporter(cfg component.Config, set exporter.Settings) (*baseExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *baseExporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

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

func (e *baseExporter) export(ctx context.Context, requestURL string, request []byte, partialSuccessHandler partialSuccessHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func isRetryableStatusCode(code int) bool { _ = "STUB: not implemented"; return false }

func readResponseBody(resp *http.Response) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readResponseStatus(resp *http.Response) *status.Status { _ = "STUB: not implemented"; return nil }

func handlePartialSuccessResponse(resp *http.Response, partialSuccessHandler partialSuccessHandler) error {
	_ = "STUB: not implemented"
	return nil
}

type partialSuccessHandler func(bytes []byte, contentType string) error

func (e *baseExporter) tracesPartialSuccessHandler(protoBytes []byte, contentType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *baseExporter) metricsPartialSuccessHandler(protoBytes []byte, contentType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *baseExporter) logsPartialSuccessHandler(protoBytes []byte, contentType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *baseExporter) profilesPartialSuccessHandler(protoBytes []byte, contentType string) error {
	_ = "STUB: not implemented"
	return nil
}
