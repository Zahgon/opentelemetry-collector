package extensions

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensioncapabilities"
	"go.opentelemetry.io/collector/service/internal/builders"
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
