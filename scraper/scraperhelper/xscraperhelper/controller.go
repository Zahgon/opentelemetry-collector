package xscraperhelper

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/xreceiver"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
	"go.opentelemetry.io/collector/scraper/scraperhelper/internal/controller"
	"go.opentelemetry.io/collector/scraper/xscraper"
)

const (
	scraperKey  = "scraper"
	spanNameSep = "/"

	receiverKey = "receiver"

	formatKey = "format"
)

type factoryWithConfig struct {
	f   xscraper.Factory
	cfg component.Config
}

type controllerOptions struct {
	tickerCh            <-chan time.Time
	factoriesWithConfig []factoryWithConfig
}

type ControllerOption interface {
	apply(*controllerOptions)
}

type optionFunc func(*controllerOptions)

func (of optionFunc) apply(e *controllerOptions) { _ = "STUB: not implemented"; return }

func AddProfilesScraper(t component.Type, sc xscraper.Profiles) ControllerOption {
	_ = "STUB: not implemented"
	return *new(ControllerOption)
}

func AddFactoryWithConfig(f xscraper.Factory, cfg component.Config) ControllerOption {
	_ = "STUB: not implemented"
	return *new(ControllerOption)
}

func WithTickerChannel(tickerCh <-chan time.Time) ControllerOption {
	_ = "STUB: not implemented"
	return *new(ControllerOption)
}

func NewProfilesController(cfg *scraperhelper.ControllerConfig,
	rSet receiver.Settings,
	nextConsumer xconsumer.Profiles,
	options ...ControllerOption,
) (xreceiver.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xreceiver.Profiles), nil
}

func getOptions(options []ControllerOption) controllerOptions {
	_ = "STUB: not implemented"
	return *new(controllerOptions)
}

func scrapeProfiles(ctx context.Context, c *controller.Controller[xscraper.Profiles], nextConsumer xconsumer.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}
