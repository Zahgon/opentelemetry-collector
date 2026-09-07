package xscraperhelper

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/scraper/xscraper"
)

const (
	scrapedProfileRecordsKey = "scraped_profile_records"

	erroredProfileRecordsKey = "errored_profile_records"
)

func wrapObsProfiles(sc xscraper.Profiles, receiverID, scraperID component.ID, set component.TelemetrySettings) (xscraper.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xscraper.Profiles), nil
}
