package main

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/otelcol"
)

type aliasProvider interface{ DeprecatedAlias() component.Type }

func makeModulesMap[T component.Factory](factories map[component.Type]T, modules map[component.Type]string) map[component.Type]string {
	_ = "STUB: not implemented"
	return nil
}

func components() (otelcol.Factories, error) {
	_ = "STUB: not implemented"
	return *new(otelcol.Factories), nil
}
