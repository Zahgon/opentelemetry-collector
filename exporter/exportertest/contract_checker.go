package exportertest

import (
	"testing"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/receiver"
)

const uniqueIDAttrName = "test_id"

type uniqueIDAttrVal string

type CheckConsumeContractParams struct {
	T                    *testing.T
	NumberOfTestElements int
	Signal               pipeline.Signal

	ExporterFactory exporter.Factory
	ExporterConfig  component.Config

	ReceiverFactory receiver.Factory
	ReceiverConfig  component.Config
}

func CheckConsumeContract(params CheckConsumeContractParams) { _ = "STUB: not implemented"; return }

func checkConsumeContractScenario(t *testing.T, params CheckConsumeContractParams, decisionFunc func() error, checkIfTestPassed func(*testing.T, int, requestCounter)) {
	_ = "STUB: not implemented"
	return
}

func checkMetrics(t *testing.T, params CheckConsumeContractParams, mockReceiver component.Component,
	mockConsumer *mockConsumer, checkIfTestPassed func(*testing.T, int, requestCounter),
) {
	_ = "STUB: not implemented"
	return
}

func checkTraces(t *testing.T, params CheckConsumeContractParams, mockReceiver component.Component, mockConsumer *mockConsumer, checkIfTestPassed func(*testing.T, int, requestCounter)) {
	_ = "STUB: not implemented"
	return
}

func checkLogs(t *testing.T, params CheckConsumeContractParams, mockReceiver component.Component, mockConsumer *mockConsumer, checkIfTestPassed func(*testing.T, int, requestCounter)) {
	_ = "STUB: not implemented"
	return
}

func alwaysSucceedsPassed(t *testing.T, allRecordsNumber int, reqCounter requestCounter) {
	_ = "STUB: not implemented"
	return
}

func randomNonPermanentErrorConsumeDecisionPassed(t *testing.T, allRecordsNumber int, reqCounter requestCounter) {
	_ = "STUB: not implemented"
	return
}

func randomPermanentErrorConsumeDecisionPassed(t *testing.T, allRecordsNumber int, reqCounter requestCounter) {
	_ = "STUB: not implemented"
	return
}

func randomErrorConsumeDecisionPassed(t *testing.T, allRecordsNumber int, reqCounter requestCounter) {
	_ = "STUB: not implemented"
	return
}

func createOneLogWithID(id uniqueIDAttrVal) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

func createOneTraceWithID(id uniqueIDAttrVal) ptrace.Traces {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces)
}

func createOneMetricWithID(id uniqueIDAttrVal) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}
