package configgrpc

import (
	"context"

	"google.golang.org/grpc"
)

type wrappedServerStream struct {
	grpc.ServerStream

	wrappedCtx context.Context
}

func (w *wrappedServerStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func wrapServerStream(wrappedCtx context.Context, stream grpc.ServerStream) *wrappedServerStream {
	_ = "STUB: not implemented"
	return nil
}
