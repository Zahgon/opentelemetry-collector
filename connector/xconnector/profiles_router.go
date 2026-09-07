package xconnector

import (
	"go.opentelemetry.io/collector/connector/internal"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/pipeline"
)

type ProfilesRouterAndConsumer interface {
	xconsumer.Profiles
	Consumer(...pipeline.ID) (xconsumer.Profiles, error)
	PipelineIDs() []pipeline.ID
	privateFunc()
}

type profilesRouter struct {
	xconsumer.Profiles
	internal.BaseRouter[xconsumer.Profiles]
}

func NewProfilesRouter(cm map[pipeline.ID]xconsumer.Profiles) ProfilesRouterAndConsumer {
	_ = "STUB: not implemented"
	return *new(ProfilesRouterAndConsumer)
}

func (r *profilesRouter) privateFunc() { _ = "STUB: not implemented"; return }
