package scraperhelper

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/scraper"
)

const (
	scraperKey = "scraper"

	scrapedMetricPointsKey = "scraped_metric_points"

	erroredMetricPointsKey = "errored_metric_points"

	spanNameSep = "/"

	receiverKey = "receiver"

	formatKey = "format"
)

func wrapObsMetrics(sc scraper.Metrics, receiverID, scraperID component.ID, set component.TelemetrySettings) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}
