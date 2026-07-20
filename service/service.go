//go:generate mdatagen metadata.yaml

package service

import (
	"context"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensioncapabilities"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/service/extensions"
	"go.opentelemetry.io/collector/service/internal/builders"
	"go.opentelemetry.io/collector/service/internal/graph"
	"go.opentelemetry.io/collector/service/internal/moduleinfo"
	"go.opentelemetry.io/collector/service/internal/proctelemetry"
	"go.opentelemetry.io/collector/service/pipelines"
	"go.opentelemetry.io/collector/service/telemetry"
)

type ModuleInfo = moduleinfo.ModuleInfo

type ModuleInfos = moduleinfo.ModuleInfos

type Settings struct {
	BuildInfo component.BuildInfo

	ConfigSnapshot extensioncapabilities.ConfigSnapshot

	CollectorConf *confmap.Conf

	ReceiversConfigs   map[component.ID]component.Config
	ReceiversFactories map[component.Type]receiver.Factory

	ProcessorsConfigs   map[component.ID]component.Config
	ProcessorsFactories map[component.Type]processor.Factory

	ExportersConfigs   map[component.ID]component.Config
	ExportersFactories map[component.Type]exporter.Factory

	ConnectorsConfigs   map[component.ID]component.Config
	ConnectorsFactories map[component.Type]connector.Factory

	Extensions builders.Extension

	ExtensionsConfigs   map[component.ID]component.Config
	ExtensionsFactories map[component.Type]extension.Factory

	ModuleInfos ModuleInfos

	AsyncErrorChannel chan error

	LoggingOptions []zap.Option

	BuildZapLogger func(zap.Config, ...zap.Option) (*zap.Logger, error)

	TelemetryFactory telemetry.Factory
}

type Service struct {
	buildInfo          component.BuildInfo
	telemetrySettings  component.TelemetrySettings
	host               *graph.Host
	configSnapshot     extensioncapabilities.ConfigSnapshot
	loggerShutdownFunc component.ShutdownFunc
	meterProvider      telemetry.MeterProvider
	tracerProvider     telemetry.TracerProvider
	graphSettings      graph.Settings
}

func New(ctx context.Context, set Settings, cfg Config) (_ *Service, resultErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (srv *Service) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func ReceiverPartialReloadEnabled() bool { _ = "STUB: not implemented"; return false }

func (srv *Service) UpdateReceivers(ctx context.Context,
	changedReceivers map[component.ID]bool,
	newReceiverConfigs map[component.ID]component.Config,
	receiverFactories map[component.Type]receiver.Factory,
	pipelineConfigs pipelines.Config,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (srv *Service) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (srv *Service) initExtensions(ctx context.Context, cfg extensions.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (srv *Service) initGraph(ctx context.Context, cfg Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (srv *Service) Logger() *zap.Logger { _ = "STUB: not implemented"; return nil }

func Validate(ctx context.Context, set Settings, cfg Config) error {
	_ = "STUB: not implemented"
	return nil
}

func registerProcessMetrics(
	srv *Service,
	goos string,
	register func(component.TelemetrySettings, ...proctelemetry.RegisterOption) error,
) error {
	_ = "STUB: not implemented"
	return nil
}
