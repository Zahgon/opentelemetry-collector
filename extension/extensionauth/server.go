package extensionauth

import (
	"context"
)

type Server interface {
	Authenticate(ctx context.Context, sources map[string][]string) (context.Context, error)
}

type ServerAuthenticateFunc func(ctx context.Context, sources map[string][]string) (context.Context, error)

func (f ServerAuthenticateFunc) Authenticate(ctx context.Context, sources map[string][]string) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}
