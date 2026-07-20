package scraper

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type Metrics interface {
	component.Component

	ScrapeMetrics(context.Context) (pmetric.Metrics, error)
}

type ScrapeMetricsFunc ScrapeFunc[pmetric.Metrics]

func (sf ScrapeMetricsFunc) ScrapeMetrics(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

type metrics struct {
	baseScraper
	ScrapeMetricsFunc
}

func NewMetrics(scrape ScrapeMetricsFunc, options ...Option) (Metrics, error) {
	_ = "STUB: not implemented"
	return *new(Metrics), nil
}
