package receivertest

import (
	"context"
	"errors"
	"sync"
	"testing"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/receiver"
)

const UniqueIDAttrName = "test_id"

type UniqueIDAttrVal string

type Generator interface {
	Start()

	Stop()

	Generate() []UniqueIDAttrVal
}

type CheckConsumeContractParams struct {
	T *testing.T

	Factory receiver.Factory
	Signal  pipeline.Signal

	Config component.Config

	Generator Generator

	GenerateCount int

	_ struct{}
}

func CheckConsumeContract(params CheckConsumeContractParams) { _ = "STUB: not implemented"; return }

func checkConsumeContractScenario(params CheckConsumeContractParams, decisionFunc func(ids idSet) error) {
	_ = "STUB: not implemented"
	return
}

type idSet map[UniqueIDAttrVal]bool

func (ds idSet) compare(other idSet) (missingInOther, onlyInOther []UniqueIDAttrVal) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds idSet) merge(other idSet) (duplicates []UniqueIDAttrVal) {
	_ = "STUB: not implemented"
	return nil
}

func (ds idSet) mergeSlice(other []UniqueIDAttrVal) (duplicates []UniqueIDAttrVal) {
	_ = "STUB: not implemented"
	return nil
}

func (ds idSet) union(other idSet) (union idSet, duplicates []UniqueIDAttrVal) {
	_ = "STUB: not implemented"
	return *new(idSet), nil
}

type consumeDecisionFunc func(ids idSet) error

var (
	errNonPermanent = errors.New("non permanent error")
	errPermanent    = errors.New("permanent error")
)

func randomNonPermanentErrorConsumeDecision(idSet) error { _ = "STUB: not implemented"; return nil }

func randomPermanentErrorConsumeDecision(idSet) error { _ = "STUB: not implemented"; return nil }

func randomErrorsConsumeDecision(idSet) error { _ = "STUB: not implemented"; return nil }

type mockConsumer struct {
	t                    *testing.T
	consumeDecisionFunc  consumeDecisionFunc
	mux                  sync.Mutex
	acceptedIDs          idSet
	droppedIDs           idSet
	nonPermanentFailures int
}

func (m *mockConsumer) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (m *mockConsumer) ConsumeTraces(_ context.Context, data ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func idSetFromTraces(data ptrace.Traces) (idSet, error) {
	_ = "STUB: not implemented"
	return *new(idSet), nil
}

func (m *mockConsumer) ConsumeLogs(_ context.Context, data plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func idSetFromLogs(data plog.Logs) (idSet, error) {
	_ = "STUB: not implemented"
	return *new(idSet), nil
}

func (m *mockConsumer) ConsumeMetrics(_ context.Context, data pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func idSetFromMetrics(data pmetric.Metrics) (idSet, error) {
	_ = "STUB: not implemented"
	return *new(idSet), nil
}

func idSetFromDataPoint(ds map[UniqueIDAttrVal]bool, attributes pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mockConsumer) consume(ids idSet) error { _ = "STUB: not implemented"; return nil }

func (m *mockConsumer) acceptedAndDropped() (acceptedAndDropped idSet, duplicates []UniqueIDAttrVal) {
	_ = "STUB: not implemented"
	return *new(idSet), nil
}

func CreateOneLogWithID(id UniqueIDAttrVal) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

func CreateGaugeMetricWithID(id UniqueIDAttrVal) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func CreateSumMetricWithID(id UniqueIDAttrVal) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func CreateSummaryMetricWithID(id UniqueIDAttrVal) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func CreateHistogramMetricWithID(id UniqueIDAttrVal) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func CreateExponentialHistogramMetricWithID(id UniqueIDAttrVal) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func CreateOneSpanWithID(id UniqueIDAttrVal) ptrace.Traces {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces)
}
