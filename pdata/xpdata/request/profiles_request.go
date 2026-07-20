package request

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pprofile"
)

func MarshalProfiles(ctx context.Context, ld pprofile.Profiles) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnmarshalProfiles(buf []byte) (context.Context, pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(pprofile.Profiles), nil
}
