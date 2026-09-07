package scraper

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
)

type Logs interface {
	component.Component

	ScrapeLogs(context.Context) (plog.Logs, error)
}

type ScrapeLogsFunc ScrapeFunc[plog.Logs]

func (sf ScrapeLogsFunc) ScrapeLogs(ctx context.Context) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

type logs struct {
	baseScraper
	ScrapeLogsFunc
}

func NewLogs(scrape ScrapeLogsFunc, options ...Option) (Logs, error) {
	_ = "STUB: not implemented"
	return *new(Logs), nil
}
