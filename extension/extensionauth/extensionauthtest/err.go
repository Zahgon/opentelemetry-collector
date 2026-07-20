package extensionauthtest

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensionauth"
)

var (
	_ extension.Extension      = (*errClient)(nil)
	_ extensionauth.HTTPClient = (*errClient)(nil)
	_ extensionauth.GRPCClient = (*errClient)(nil)
)

type errClient struct {
	component.StartFunc
	component.ShutdownFunc
	extensionauth.ClientPerRPCCredentialsFunc
	extensionauth.ClientRoundTripperFunc
	extensionauth.ServerAuthenticateFunc
}

func NewErr(err error) extension.Extension {
	_ = "STUB: not implemented"
	return *new(extension.Extension)
}
