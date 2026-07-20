package processortest

import (
	"testing"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/processor"
)

func verifyTracesDoesNotProduceAfterShutdown(t *testing.T, factory processor.Factory, cfg component.Config) {
	_ = "STUB: not implemented"
	return
}

func verifyLogsDoesNotProduceAfterShutdown(t *testing.T, factory processor.Factory, cfg component.Config) {
	_ = "STUB: not implemented"
	return
}

func verifyMetricsDoesNotProduceAfterShutdown(t *testing.T, factory processor.Factory, cfg component.Config) {
	_ = "STUB: not implemented"
	return
}

func VerifyShutdown(t *testing.T, factory processor.Factory, cfg component.Config) {
	_ = "STUB: not implemented"
	return
}
