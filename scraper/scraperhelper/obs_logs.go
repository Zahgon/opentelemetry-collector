package scraperhelper

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/scraper"
)

const (
	scrapedLogRecordsKey = "scraped_log_records"

	erroredLogRecordsKey = "errored_log_records"
)

func wrapObsLogs(sc scraper.Logs, receiverID, scraperID component.ID, set component.TelemetrySettings) (scraper.Logs, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Logs), nil
}
