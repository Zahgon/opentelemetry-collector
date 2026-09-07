package xscraper

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/scraper"
)

type Profiles interface {
	component.Component

	ScrapeProfiles(context.Context) (pprofile.Profiles, error)
}

type ScrapeProfilesFunc scraper.ScrapeFunc[pprofile.Profiles]

func (sf ScrapeProfilesFunc) ScrapeProfiles(ctx context.Context) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}

type profiles struct {
	baseScraper
	ScrapeProfilesFunc
}

func NewProfiles(scrape ScrapeProfilesFunc, options ...Option) (Profiles, error) {
	_ = "STUB: not implemented"
	return *new(Profiles), nil
}
