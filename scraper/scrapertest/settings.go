package scrapertest

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/scraper"
)

var NopType = component.MustNewType("nop")

func NewNopSettings(typ component.Type) scraper.Settings {
	_ = "STUB: not implemented"
	return *new(scraper.Settings)
}
