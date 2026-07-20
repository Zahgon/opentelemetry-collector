package configauth

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/extensionauth"
)

var (
	errAuthenticatorNotFound = errors.New("authenticator not found")
	errNotHTTPClient         = errors.New("requested authenticator is not a HTTP client authenticator")
	errNotGRPCClient         = errors.New("requested authenticator is not a gRPC client authenticator")
	errNotServer             = errors.New("requested authenticator is not a server authenticator")
)

type Config struct {
	AuthenticatorID component.ID `mapstructure:"authenticator,omitempty"`

	_ struct{}
}

func (a Config) GetServerAuthenticator(_ context.Context, extensions map[component.ID]component.Component) (extensionauth.Server, error) {
	_ = "STUB: not implemented"
	return *new(extensionauth.Server), nil
}

func (a Config) GetHTTPClientAuthenticator(_ context.Context, extensions map[component.ID]component.Component) (extensionauth.HTTPClient, error) {
	_ = "STUB: not implemented"
	return *new(extensionauth.HTTPClient), nil
}

func (a Config) GetGRPCClientAuthenticator(_ context.Context, extensions map[component.ID]component.Component) (extensionauth.GRPCClient, error) {
	_ = "STUB: not implemented"
	return *new(extensionauth.GRPCClient), nil
}
