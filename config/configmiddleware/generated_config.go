package configmiddleware

import (
	"go.opentelemetry.io/collector/component"
)

type Config struct {
	ID component.ID `mapstructure:"id"`

	_ struct{}
}
