//go:generate mdatagen metadata.yaml

package fileprovider

import (
	"context"

	"go.opentelemetry.io/collector/confmap"
)

const schemeName = "file"

type provider struct{}

func NewFactory() confmap.ProviderFactory {
	_ = "STUB: not implemented"
	return *new(confmap.ProviderFactory)
}

func newProvider(confmap.ProviderSettings) confmap.Provider {
	_ = "STUB: not implemented"
	return *new(confmap.Provider)
}

func (fmp *provider) Retrieve(_ context.Context, uri string, _ confmap.WatcherFunc) (*confmap.Retrieved, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*provider) Scheme() string { _ = "STUB: not implemented"; return "" }

func (*provider) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }
