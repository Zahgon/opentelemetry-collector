package extensionauthtest

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensionauth"
)

var (
	_ extension.Extension      = (*nopClient)(nil)
	_ extensionauth.HTTPClient = (*nopClient)(nil)
	_ extensionauth.GRPCClient = (*nopClient)(nil)
)

type nopClient struct {
	component.StartFunc
	component.ShutdownFunc
	extensionauth.ClientRoundTripperFunc
	extensionauth.ClientPerRPCCredentialsFunc
}

func NewNopClient() extension.Extension {
	_ = "STUB: not implemented"
	return *new(extension.Extension)
}
