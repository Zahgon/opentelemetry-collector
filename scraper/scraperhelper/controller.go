package scraperhelper

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"
	"go.opentelemetry.io/collector/scraper/scraperhelper/internal/controller"
)

type ControllerOption interface {
	apply(*controllerOptions)
}

type optionFunc func(*controllerOptions)

func (of optionFunc) apply(e *controllerOptions) { _ = "STUB: not implemented"; return }

func AddMetricsScraper(t component.Type, sc scraper.Metrics) ControllerOption {
	_ = "STUB: not implemented"
	return *new(ControllerOption)
}

func AddScraper(t component.Type, sc scraper.Metrics) ControllerOption {
	_ = "STUB: not implemented"
	return *new(ControllerOption)
}

func AddFactoryWithConfig(f scraper.Factory, cfg component.Config) ControllerOption {
	_ = "STUB: not implemented"
	return *new(ControllerOption)
}

func WithTickerChannel(tickerCh <-chan time.Time) ControllerOption {
	_ = "STUB: not implemented"
	return *new(ControllerOption)
}

type factoryWithConfig struct {
	f   scraper.Factory
	cfg component.Config
}

type controllerOptions struct {
	tickerCh            <-chan time.Time
	factoriesWithConfig []factoryWithConfig
}

func NewLogsController(cfg *ControllerConfig,
	rSet receiver.Settings,
	nextConsumer consumer.Logs,
	options ...ControllerOption,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

func NewMetricsController(cfg *ControllerConfig,
	rSet receiver.Settings,
	nextConsumer consumer.Metrics,
	options ...ControllerOption,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func scrapeLogs(ctx context.Context, c *controller.Controller[scraper.Logs], nextConsumer consumer.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func scrapeMetrics(ctx context.Context, c *controller.Controller[scraper.Metrics], nextConsumer consumer.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func getOptions(options []ControllerOption) controllerOptions {
	_ = "STUB: not implemented"
	return *new(controllerOptions)
}
