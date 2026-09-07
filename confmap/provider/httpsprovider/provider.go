//go:generate mdatagen metadata.yaml

package httpsprovider

import (
	"go.opentelemetry.io/collector/confmap"
)

func NewFactory() confmap.ProviderFactory {
	_ = "STUB: not implemented"
	return *new(confmap.ProviderFactory)
}

func newProvider(set confmap.ProviderSettings) confmap.Provider {
	_ = "STUB: not implemented"
	return *new(confmap.Provider)
}
