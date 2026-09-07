package extensionauthtest

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensionauth"
)

var (
	_ extension.Extension  = (*nopServer)(nil)
	_ extensionauth.Server = (*nopServer)(nil)
)

type nopServer struct {
	component.StartFunc
	component.ShutdownFunc
}

func (n *nopServer) Authenticate(ctx context.Context, _ map[string][]string) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func NewNopServer() extension.Extension {
	_ = "STUB: not implemented"
	return *new(extension.Extension)
}
