package xprocessorhelper

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/xprocessor"
)

type ProcessProfilesFunc func(context.Context, pprofile.Profiles) (pprofile.Profiles, error)

type profiles struct {
	component.StartFunc
	component.ShutdownFunc
	xconsumer.Profiles
}

func NewProfiles(
	_ context.Context,
	_ processor.Settings,
	_ component.Config,
	nextConsumer xconsumer.Profiles,
	profilesFunc ProcessProfilesFunc,
	options ...Option,
) (xprocessor.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xprocessor.Profiles), nil
}
