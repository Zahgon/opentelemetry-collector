package otelcol

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pipeline"
)

type configFingerprint struct {
	nonReceiverHash uint64

	receiverHashes map[component.ID]uint64

	extensions []component.ID

	pipelines map[pipeline.ID]pipelineFingerprint
}

type pipelineFingerprint struct {
	receivers  []component.ID
	processors []component.ID
	exporters  []component.ID
}

func fingerprintForPartialReload(conf *confmap.Conf, cfg *Config) (configFingerprint, error) {
	_ = "STUB: not implemented"
	return *new(configFingerprint), nil
}

func hashSections(conf *confmap.Conf, keys ...string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func hashEntriesByID(conf *confmap.Conf, key string) (map[component.ID]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hashValue(v any) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func receiversOnlyChanged(old, cur configFingerprint, isConnector func(component.ID) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func filterIDs(ids []component.ID, pred func(component.ID) bool) []component.ID {
	_ = "STUB: not implemented"
	return nil
}

func isConnectorID(connectors map[component.ID]component.Config) func(component.ID) bool {
	_ = "STUB: not implemented"
	return nil
}

func changedReceivers(old, cur configFingerprint) map[component.ID]bool {
	_ = "STUB: not implemented"
	return nil
}
