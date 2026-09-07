package extensiontest

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

var NopType = component.MustNewType("nop")

func NewNopSettings(typ component.Type) extension.Settings {
	_ = "STUB: not implemented"
	return *new(extension.Settings)
}

func NewNopFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

type nopConfig struct{}

var nopInstance = &nopExtension{}

type nopExtension struct {
	component.StartFunc
	component.ShutdownFunc
}
