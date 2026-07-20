package plogotlp

import (
	"context"

	"google.golang.org/grpc"

	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/internal/otelgrpc"
)

type GRPCClient interface {
	Export(ctx context.Context, request ExportRequest, opts ...grpc.CallOption) (ExportResponse, error)

	unexported()
}

func NewGRPCClient(cc *grpc.ClientConn) GRPCClient {
	_ = "STUB: not implemented"
	return *new(GRPCClient)
}

type grpcClient struct {
	rawClient otelgrpc.LogsServiceClient
}

func (c *grpcClient) Export(ctx context.Context, request ExportRequest, opts ...grpc.CallOption) (ExportResponse, error) {
	_ = "STUB: not implemented"
	return *new(ExportResponse), nil
}

func (c *grpcClient) unexported() { _ = "STUB: not implemented"; return }

type GRPCServer interface {
	Export(context.Context, ExportRequest) (ExportResponse, error)

	unexported()
}

var _ GRPCServer = (*UnimplementedGRPCServer)(nil)

type UnimplementedGRPCServer struct{}

func (*UnimplementedGRPCServer) Export(context.Context, ExportRequest) (ExportResponse, error) {
	_ = "STUB: not implemented"
	return *new(ExportResponse), nil
}

func (*UnimplementedGRPCServer) unexported() { _ = "STUB: not implemented"; return }

func RegisterGRPCServer(s *grpc.Server, srv GRPCServer) { _ = "STUB: not implemented"; return }

type rawLogsServer struct {
	srv GRPCServer
}

func (s rawLogsServer) Export(ctx context.Context, request *internal.ExportLogsServiceRequest) (*internal.ExportLogsServiceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
