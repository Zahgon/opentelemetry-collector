package hostcapabilities

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/service/internal/moduleinfo"
)

type ModuleInfo interface {
	GetModuleInfos() moduleinfo.ModuleInfos
}

type ExposeExporters interface {
	GetExporters() map[pipeline.Signal]map[component.ID]component.Component
}

type ComponentFactory interface {
	GetFactory(kind component.Kind, componentType component.Type) component.Factory
}
