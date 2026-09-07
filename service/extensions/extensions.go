package extensions

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensioncapabilities"
	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/service/hostcapabilities"
	"go.opentelemetry.io/collector/service/internal/builders"
	"go.opentelemetry.io/collector/service/internal/moduleinfo"
	"go.opentelemetry.io/collector/service/internal/status"
)

const zExtensionName = "zextensionname"

type Extensions struct {
	telemetry    component.TelemetrySettings
	extMap       map[component.ID]extension.Extension
	instanceIDs  map[component.ID]*componentstatus.InstanceID
	extensionIDs []component.ID
	reporter     status.Reporter
}

func (bes *Extensions) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (bes *Extensions) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (bes *Extensions) NotifyPipelineReady() error { _ = "STUB: not implemented"; return nil }

func (bes *Extensions) NotifyPipelineNotReady() error { _ = "STUB: not implemented"; return nil }

func (bes *Extensions) NotifyConfig(ctx context.Context, conf *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

func (bes *Extensions) NotifyConfigSnapshot(ctx context.Context, configSnapshot extensioncapabilities.ConfigSnapshot) error {
	_ = "STUB: not implemented"
	return nil
}

func (bes *Extensions) NotifyComponentStatusChange(source *componentstatus.InstanceID, event *componentstatus.Event) {
	_ = "STUB: not implemented"
	return
}

func (bes *Extensions) GetExtensions() map[component.ID]component.Component {
	_ = "STUB: not implemented"
	return nil
}

func (bes *Extensions) HandleZPages(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type Settings struct {
	Telemetry component.TelemetrySettings
	BuildInfo component.BuildInfo

	Extensions builders.Extension
}

type Option interface {
	apply(*Extensions)
}

type optionFunc func(*Extensions)

func (of optionFunc) apply(e *Extensions) { _ = "STUB: not implemented"; return }

func WithReporter(reporter status.Reporter) Option { _ = "STUB: not implemented"; return *new(Option) }

func New(ctx context.Context, set Settings, cfg Config, options ...Option) (*Extensions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	_ componentstatus.Reporter          = (*hostWrapper)(nil)
	_ component.Host                    = (*hostWrapper)(nil)
	_ hostcapabilities.ModuleInfo       = (*hostWrapper)(nil)
	_ hostcapabilities.ExposeExporters  = (*hostWrapper)(nil) //nolint:staticcheck // SA1019
	_ hostcapabilities.ComponentFactory = (*hostWrapper)(nil)
)

type hostWrapper struct {
	component.Host
	reporter   status.Reporter
	instanceID *componentstatus.InstanceID
}

func (host *hostWrapper) Report(event *componentstatus.Event) { _ = "STUB: not implemented"; return }

func (host *hostWrapper) RegisterZPages(mux *http.ServeMux, pathPrefix string) {
	_ = "STUB: not implemented"
	return
}

func (host *hostWrapper) GetModuleInfos() moduleinfo.ModuleInfos {
	_ = "STUB: not implemented"
	return *new(moduleinfo.ModuleInfos)
}

//nolint:staticcheck // SA1019 forwards the deprecated hostcapabilities.ExposeExporters capability.
func (host *hostWrapper) GetExporters() map[pipeline.Signal]map[component.ID]component.Component {
	_ = "STUB: not implemented"
	return nil
}

func (host *hostWrapper) GetFactory(kind component.Kind, componentType component.Type) component.Factory {
	_ = "STUB: not implemented"
	return *new(component.Factory)
}
