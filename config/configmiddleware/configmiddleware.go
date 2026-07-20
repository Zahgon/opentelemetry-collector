package configmiddleware

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/extensionmiddleware"
)

var (
	errMiddlewareNotFound = errors.New("middleware not found")
	errNotHTTPServer      = errors.New("requested extension is not an HTTP server middleware")
	errNotGRPCServer      = errors.New("requested extension is not a gRPC server middleware")
	errNotHTTPClient      = errors.New("requested extension is not an HTTP client middleware")
	errNotGRPCClient      = errors.New("requested extension is not a gRPC client middleware")
)

func (m Config) GetHTTPClientRoundTripper(ctx context.Context, extensions map[component.ID]component.Component) (extensionmiddleware.WrapHTTPRoundTripperFunc, error) {
	_ = "STUB: not implemented"
	return *new(extensionmiddleware.WrapHTTPRoundTripperFunc), nil
}

func (m Config) GetHTTPServerHandler(ctx context.Context, extensions map[component.ID]component.Component) (extensionmiddleware.WrapHTTPHandlerFunc, error) {
	_ = "STUB: not implemented"
	return *new(extensionmiddleware.WrapHTTPHandlerFunc), nil
}

func (m Config) GetGRPCClientOptions(ctx context.Context, extensions map[component.ID]component.Component) ([]grpc.DialOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m Config) GetGRPCServerOptions(ctx context.Context, extensions map[component.ID]component.Component) ([]grpc.ServerOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
