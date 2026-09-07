package scraper

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
)

var errNilFunc = errors.New("nil scrape func")

type ScrapeFunc[T any] func(context.Context) (T, error)

type Option interface {
	apply(*baseScraper)
}

type scraperOptionFunc func(*baseScraper)

func (of scraperOptionFunc) apply(e *baseScraper) { _ = "STUB: not implemented"; return }

func WithStart(start component.StartFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithShutdown(shutdown component.ShutdownFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type baseScraper struct {
	component.StartFunc
	component.ShutdownFunc
}

func newBaseScraper(options []Option) baseScraper {
	_ = "STUB: not implemented"
	return *new(baseScraper)
}
