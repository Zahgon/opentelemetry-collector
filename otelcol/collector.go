package otelcol

import (
	"context"
	"os"
	"sync"
	"sync/atomic"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service"
)

type State int

const (
	StateStarting State = iota
	StateRunning
	StateClosing
	StateClosed
)

func (s State) String() string { _ = "STUB: not implemented"; return "" }

type CollectorSettings struct {
	Factories func() (Factories, error)

	BuildInfo component.BuildInfo

	DisableGracefulShutdown bool

	ConfigProviderSettings ConfigProviderSettings

	ProviderModules map[string]string

	ConverterModules []string

	LoggingOptions []zap.Option

	SkipSettingGRPCLogger bool
}

type Collector struct {
	set            CollectorSettings
	buildZapLogger func(zap.Config, ...zap.Option) (*zap.Logger, error)

	configProvider *ConfigProvider

	service *service.Service
	state   *atomic.Int64

	shutdownChan chan struct{}
	shutdownOnce sync.Once

	wg sync.WaitGroup

	signalsChannel chan os.Signal

	asyncErrorChannel          chan error
	bc                         *bufferedCore
	updateConfigProviderLogger func(core zapcore.Core)

	currentFingerprint *configFingerprint
}

func NewCollector(set CollectorSettings) (*Collector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Collector) GetState() State { _ = "STUB: not implemented"; return *new(State) }

func (col *Collector) Shutdown() { _ = "STUB: not implemented"; return }

func buildModuleInfo(m map[component.Type]string) map[component.Type]service.ModuleInfo {
	_ = "STUB: not implemented"
	return nil
}

func (col *Collector) setupConfigurationComponents(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Collector) reloadConfiguration(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Collector) tryPartialReceiverReload(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (col *Collector) DryRun(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func newFallbackLogger(options []zap.Option) (*zap.Logger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Collector) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

//nolint:contextcheck

func (col *Collector) shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (col *Collector) setCollectorState(state State) { _ = "STUB: not implemented"; return }
