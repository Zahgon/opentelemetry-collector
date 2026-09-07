package exportertest

import (
	"context"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var (
	errNonPermanent = status.Error(codes.DeadlineExceeded, "non Permanent error")
	errPermanent    = status.Error(codes.Internal, "Permanent error")
)

func randomNonPermanentErrorConsumeDecision() error { _ = "STUB: not implemented"; return nil }

func randomPermanentErrorConsumeDecision() error { _ = "STUB: not implemented"; return nil }

func randomErrorsConsumeDecision() error { _ = "STUB: not implemented"; return nil }

type mockConsumer struct {
	consumer.Traces
	consumer.Logs
	consumer.Metrics
	reqCounter          *requestCounter
	mux                 sync.Mutex
	exportErrorFunction func() error
	receivedTraces      []ptrace.Traces
	receivedMetrics     []pmetric.Metrics
	receivedLogs        []plog.Logs
}

func newMockConsumer(decisionFunc func() error) mockConsumer {
	_ = "STUB: not implemented"
	return *new(mockConsumer)
}

func (r *mockConsumer) ConsumeLogs(_ context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *mockConsumer) ConsumeTraces(_ context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *mockConsumer) ConsumeMetrics(_ context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *mockConsumer) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (r *mockConsumer) processError(err error) { _ = "STUB: not implemented"; return }

func (r *mockConsumer) clear() { _ = "STUB: not implemented"; return }

func (r *mockConsumer) getRequestCounter() *requestCounter { _ = "STUB: not implemented"; return nil }

type requestCounter struct {
	success int
	error   errorCounter
	total   int
}

type errorCounter struct {
	permanent    int
	nonpermanent int
}

func newErrorCounter() errorCounter { _ = "STUB: not implemented"; return *new(errorCounter) }

func newRequestCounter() *requestCounter { _ = "STUB: not implemented"; return nil }

func idFromLogs(data plog.Logs) (string, error) { _ = "STUB: not implemented"; return "", nil }

func idFromTraces(data ptrace.Traces) (string, error) { _ = "STUB: not implemented"; return "", nil }

func idFromMetrics(data pmetric.Metrics) (string, error) { _ = "STUB: not implemented"; return "", nil }
