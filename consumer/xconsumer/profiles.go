package xconsumer

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/internal"
	"go.opentelemetry.io/collector/pdata/pprofile"
)

var errNilFunc = errors.New("nil consumer func")

type Profiles interface {
	internal.BaseConsumer

	ConsumeProfiles(ctx context.Context, td pprofile.Profiles) error
}

type ConsumeProfilesFunc func(ctx context.Context, td pprofile.Profiles) error

func (f ConsumeProfilesFunc) ConsumeProfiles(ctx context.Context, td pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

type baseProfiles struct {
	*internal.BaseImpl
	ConsumeProfilesFunc
}

func NewProfiles(consume ConsumeProfilesFunc, options ...consumer.Option) (Profiles, error) {
	_ = "STUB: not implemented"
	return *new(Profiles), nil
}
