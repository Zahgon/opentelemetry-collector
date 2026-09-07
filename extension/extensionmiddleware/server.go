package extensionmiddleware

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
)

type HTTPServer interface {
	GetHTTPHandler(_ context.Context) (WrapHTTPHandlerFunc, error)
}

type GRPCServer interface {
	GetGRPCServerOptions(context.Context) ([]grpc.ServerOption, error)
}

var _ HTTPServer = (*GetHTTPHandlerFunc)(nil)

type GetHTTPHandlerFunc func(_ context.Context) (WrapHTTPHandlerFunc, error)

func (f GetHTTPHandlerFunc) GetHTTPHandler(ctx context.Context) (WrapHTTPHandlerFunc, error) {
	_ = "STUB: not implemented"
	return *new(WrapHTTPHandlerFunc), nil
}

var _ GRPCServer = (*GetGRPCServerOptionsFunc)(nil)

type GetGRPCServerOptionsFunc func(context.Context) ([]grpc.ServerOption, error)

func (f GetGRPCServerOptionsFunc) GetGRPCServerOptions(ctx context.Context) ([]grpc.ServerOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type WrapHTTPHandlerFunc = func(context.Context, http.Handler) (http.Handler, error)
