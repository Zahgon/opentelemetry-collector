package otelcol

import (
	"github.com/spf13/cobra"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
)

type componentWithStability struct {
	Name      component.Type
	Module    string
	Stability map[string]string
}

type componentWithoutStability struct {
	Scheme string `yaml:",omitempty"`
	Module string
}

type componentsOutput struct {
	BuildInfo  component.BuildInfo
	Receivers  []componentWithStability
	Processors []componentWithStability
	Exporters  []componentWithStability
	Connectors []componentWithStability
	Extensions []componentWithStability
	Providers  []componentWithoutStability
	Converters []componentWithoutStability `yaml:",omitempty"`
}

func newComponentsCommand(set CollectorSettings) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func canonicalFactoryKeys[T component.Factory](factories map[component.Type]T) []component.Type {
	_ = "STUB: not implemented"
	return nil
}

func sortFactoriesByType[T component.Factory](factories map[component.Type]T) []T {
	_ = "STUB: not implemented"
	return nil
}

func sortProvidersByScheme(providerModules map[string]string, provFactories []confmap.ProviderFactory, set confmap.ProviderSettings) []componentWithoutStability {
	_ = "STUB: not implemented"
	return nil
}

func sortConverterModules(modules []string) []componentWithoutStability {
	_ = "STUB: not implemented"
	return nil
}
