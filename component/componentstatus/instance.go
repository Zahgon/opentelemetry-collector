package componentstatus

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pipeline"
)

const pipelineDelim = byte(0x20)

type InstanceID struct {
	componentID component.ID
	kind        component.Kind
	pipelineIDs string
}

func NewInstanceID(componentID component.ID, kind component.Kind, pipelineIDs ...pipeline.ID) *InstanceID {
	_ = "STUB: not implemented"
	return nil
}

func (id *InstanceID) ComponentID() component.ID {
	_ = "STUB: not implemented"
	return *new(component.ID)
}

func (id *InstanceID) Kind() component.Kind { _ = "STUB: not implemented"; return *new(component.Kind) }

func (id *InstanceID) AllPipelineIDs(f func(pipeline.ID) bool) { _ = "STUB: not implemented"; return }

func (id *InstanceID) WithPipelines(pipelineIDs ...pipeline.ID) *InstanceID {
	_ = "STUB: not implemented"
	return nil
}

func (id *InstanceID) addPipelines(pipelineIDs []pipeline.ID) { _ = "STUB: not implemented"; return }
