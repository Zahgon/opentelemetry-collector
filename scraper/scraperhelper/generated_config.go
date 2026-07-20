package scraperhelper

import (
	"go.opentelemetry.io/collector/scraper/scraperhelper/internal/controller"
)

type ControllerConfig = controller.ControllerConfig

func NewDefaultControllerConfig() ControllerConfig {
	_ = "STUB: not implemented"
	return *new(ControllerConfig)
}
