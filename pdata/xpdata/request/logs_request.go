package request

import (
	"context"

	"go.opentelemetry.io/collector/pdata/plog"
)

func MarshalLogs(ctx context.Context, ld plog.Logs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnmarshalLogs(buf []byte) (context.Context, plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(plog.Logs), nil
}
