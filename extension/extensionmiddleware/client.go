package extensionmiddleware

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
)

type HTTPClient interface {
	GetHTTPRoundTripper(context.Context) (WrapHTTPRoundTripperFunc, error)
}

type GRPCClient interface {
	GetGRPCClientOptions(context.Context) ([]grpc.DialOption, error)
}

var _ HTTPClient = (*GetHTTPRoundTripperFunc)(nil)

type GetHTTPRoundTripperFunc func(context.Context) (WrapHTTPRoundTripperFunc, error)

func (f GetHTTPRoundTripperFunc) GetHTTPRoundTripper(ctx context.Context) (WrapHTTPRoundTripperFunc, error) {
	_ = "STUB: not implemented"
	return *new(WrapHTTPRoundTripperFunc), nil
}

var _ GRPCClient = (*GetGRPCClientOptionsFunc)(nil)

type GetGRPCClientOptionsFunc func(context.Context) ([]grpc.DialOption, error)

func (f GetGRPCClientOptionsFunc) GetGRPCClientOptions(ctx context.Context) ([]grpc.DialOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type WrapHTTPRoundTripperFunc = func(context.Context, http.RoundTripper) (http.RoundTripper, error)
