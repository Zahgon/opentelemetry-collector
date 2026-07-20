package otlpreceiver

import (
	"context"
	"net/http"
	"sync"

	"google.golang.org/grpc"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

type otlpReceiver struct {
	cfg        *Config
	serverGRPC *grpc.Server
	serverHTTP *http.Server

	nextTraces   consumer.Traces
	nextMetrics  consumer.Metrics
	nextLogs     consumer.Logs
	nextProfiles xconsumer.Profiles
	shutdownWG   sync.WaitGroup

	obsrepGRPC *receiverhelper.ObsReport
	obsrepHTTP *receiverhelper.ObsReport

	settings *receiver.Settings
}

func newOtlpReceiver(cfg *Config, set *receiver.Settings) (*otlpReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *otlpReceiver) startGRPCServer(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *otlpReceiver) startHTTPServer(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *otlpReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *otlpReceiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *otlpReceiver) registerTraceConsumer(tc consumer.Traces) { _ = "STUB: not implemented"; return }

func (r *otlpReceiver) registerMetricsConsumer(mc consumer.Metrics) {
	_ = "STUB: not implemented"
	return
}

func (r *otlpReceiver) registerLogsConsumer(lc consumer.Logs) { _ = "STUB: not implemented"; return }

func (r *otlpReceiver) registerProfilesConsumer(tc xconsumer.Profiles) {
	_ = "STUB: not implemented"
	return
}
