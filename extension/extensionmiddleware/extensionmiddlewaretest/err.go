package extensionmiddlewaretest

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensionmiddleware"
)

var (
	_ extension.Extension            = (*baseExtension)(nil)
	_ extensionmiddleware.HTTPClient = (*baseExtension)(nil)
	_ extensionmiddleware.GRPCClient = (*baseExtension)(nil)
	_ extensionmiddleware.HTTPServer = (*baseExtension)(nil)
	_ extensionmiddleware.GRPCServer = (*baseExtension)(nil)
)

type baseExtension struct {
	component.StartFunc
	component.ShutdownFunc
	extensionmiddleware.GetHTTPHandlerFunc
	extensionmiddleware.GetGRPCServerOptionsFunc
	extensionmiddleware.GetHTTPRoundTripperFunc
	extensionmiddleware.GetGRPCClientOptionsFunc
}

func NewErr(err error) extension.Extension {
	_ = "STUB: not implemented"
	return *new(extension.Extension)
}
