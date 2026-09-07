package telemetry

import (
	"context"

	otelconf "go.opentelemetry.io/contrib/otelconf/v0.3.0"
	"go.opentelemetry.io/otel/metric"
	noopmetric "go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"
	nooptrace "go.opentelemetry.io/otel/trace/noop"
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type LoggerSettings struct {
	Settings

	ZapOptions []zap.Option

	BuildZapLogger func(zap.Config, ...zap.Option) (*zap.Logger, error)
}

type MeterSettings struct {
	Settings

	Logger *zap.Logger

	DefaultViews func(configtelemetry.Level) []otelconf.View
}

type TracerSettings struct {
	Settings

	Logger *zap.Logger
}

type Settings struct {
	BuildInfo component.BuildInfo

	Resource *pcommon.Resource

	SchemaURL string
}

type Factory interface {
	CreateDefaultConfig() component.Config

	CreateResource(context.Context, Settings, component.Config) (pcommon.Resource, string, error)

	CreateLogger(context.Context, LoggerSettings, component.Config) (*zap.Logger, component.ShutdownFunc, error)

	CreateMeterProvider(context.Context, MeterSettings, component.Config) (MeterProvider, error)

	CreateTracerProvider(context.Context, TracerSettings, component.Config) (TracerProvider, error)

	unexportedFactoryFunc()
}

type MeterProvider interface {
	metric.MeterProvider
	Shutdown(context.Context) error
}

type TracerProvider interface {
	trace.TracerProvider
	Shutdown(context.Context) error
}

type FactoryOption interface {
	applyOption(*factory)
}

type factoryOptionFunc func(*factory)

func (f factoryOptionFunc) applyOption(o *factory) { _ = "STUB: not implemented"; return }

type factory struct {
	component.CreateDefaultConfigFunc
	createResourceFunc       CreateResourceFunc
	createLoggerFunc         CreateLoggerFunc
	createMeterProviderFunc  CreateMeterProviderFunc
	createTracerProviderFunc CreateTracerProviderFunc
}

func NewFactory(createDefaultConfig component.CreateDefaultConfigFunc, opts ...FactoryOption) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

func WithCreateResource(createResource CreateResourceFunc) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

type CreateResourceFunc func(context.Context, Settings, component.Config) (pcommon.Resource, string, error)

func WithCreateLogger(createLogger CreateLoggerFunc) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

type CreateLoggerFunc func(context.Context, LoggerSettings, component.Config) (*zap.Logger, component.ShutdownFunc, error)

func WithCreateMeterProvider(createMeterProvider CreateMeterProviderFunc) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

type CreateMeterProviderFunc func(context.Context, MeterSettings, component.Config) (MeterProvider, error)

func WithCreateTracerProvider(createTracerProvider CreateTracerProviderFunc) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

type CreateTracerProviderFunc func(context.Context, TracerSettings, component.Config) (TracerProvider, error)

func (*factory) unexportedFactoryFunc() { _ = "STUB: not implemented"; return }

func (f *factory) CreateResource(ctx context.Context, settings Settings, cfg component.Config) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

func (f *factory) CreateLogger(ctx context.Context, settings LoggerSettings, cfg component.Config) (*zap.Logger, component.ShutdownFunc, error) {
	_ = "STUB: not implemented"
	return nil, *new(component.ShutdownFunc), nil
}

func (f *factory) CreateMeterProvider(ctx context.Context, settings MeterSettings, cfg component.Config) (MeterProvider, error) {
	_ = "STUB: not implemented"
	return *new(MeterProvider), nil
}

func (f *factory) CreateTracerProvider(ctx context.Context, settings TracerSettings, cfg component.Config) (TracerProvider, error) {
	_ = "STUB: not implemented"
	return *new(TracerProvider), nil
}

type noopMeterProvider struct {
	noopmetric.MeterProvider
	component.ShutdownFunc
}

type noopTracerProvider struct {
	nooptrace.TracerProvider
	component.ShutdownFunc
}
