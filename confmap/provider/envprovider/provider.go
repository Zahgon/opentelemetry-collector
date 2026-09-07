//go:generate mdatagen metadata.yaml

package envprovider

import (
	"context"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/confmap"
)

const (
	schemeName = "env"
)

type provider struct {
	logger *zap.Logger
}

func NewFactory() confmap.ProviderFactory {
	_ = "STUB: not implemented"
	return *new(confmap.ProviderFactory)
}

func newProvider(ps confmap.ProviderSettings) confmap.Provider {
	_ = "STUB: not implemented"
	return *new(confmap.Provider)
}

func (emp *provider) Retrieve(_ context.Context, uri string, _ confmap.WatcherFunc) (*confmap.Retrieved, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*provider) Scheme() string { _ = "STUB: not implemented"; return "" }

func (*provider) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func parseEnvVarURI(uri string) (string, *string) { _ = "STUB: not implemented"; return "", nil }
