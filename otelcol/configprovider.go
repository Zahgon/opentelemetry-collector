package otelcol

import (
	"context"

	"go.opentelemetry.io/collector/confmap"
)

type ConfigProvider struct {
	mapResolver *confmap.Resolver
}

type ConfigProviderSettings struct {
	ResolverSettings confmap.ResolverSettings

	_ struct{}
}

func NewConfigProvider(set ConfigProviderSettings) (*ConfigProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cm *ConfigProvider) Get(ctx context.Context, factories Factories) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cm *ConfigProvider) getWithConf(ctx context.Context, factories Factories) (*Config, *confmap.Conf, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (cm *ConfigProvider) Watch() <-chan error { _ = "STUB: not implemented"; return nil }

func (cm *ConfigProvider) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
