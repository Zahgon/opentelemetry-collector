package confmaptest

import (
	"go.opentelemetry.io/collector/confmap"
)

func NewNopProviderSettings() confmap.ProviderSettings {
	_ = "STUB: not implemented"
	return *new(confmap.ProviderSettings)
}
