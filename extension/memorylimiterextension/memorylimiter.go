package memorylimiterextension

import (
	"context"
	"net/http"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/extensionmiddleware"
	"go.opentelemetry.io/collector/internal/memorylimiter"
)

var (
	_ extensionmiddleware.GRPCServer = (*memoryLimiterExtension)(nil)
	_ extensionmiddleware.HTTPServer = (*memoryLimiterExtension)(nil)
)

type memoryLimiterExtension struct {
	memLimiter *memorylimiter.MemoryLimiter
}

func newMemoryLimiter(cfg *Config, logger *zap.Logger) (*memoryLimiterExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ml *memoryLimiterExtension) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (ml *memoryLimiterExtension) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (ml *memoryLimiterExtension) MustRefuse() bool { _ = "STUB: not implemented"; return false }

func (ml *memoryLimiterExtension) GetHTTPHandler(_ context.Context) (extensionmiddleware.WrapHTTPHandlerFunc, error) {
	_ = "STUB: not implemented"
	return *new(extensionmiddleware.WrapHTTPHandlerFunc), nil
}

func (ml *memoryLimiterExtension) wrapHTTPHandler(_ context.Context, base http.Handler) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (ml *memoryLimiterExtension) GetGRPCServerOptions(_ context.Context) ([]grpc.ServerOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
