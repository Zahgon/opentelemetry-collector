package request

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

func MarshalMetrics(ctx context.Context, ld pmetric.Metrics) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnmarshalMetrics(buf []byte) (context.Context, pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(pmetric.Metrics), nil
}
