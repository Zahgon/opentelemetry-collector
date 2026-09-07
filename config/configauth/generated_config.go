package configauth

import (
	"go.opentelemetry.io/collector/component"
)

type Config struct {
	AuthenticatorID component.ID `mapstructure:"authenticator,omitempty"`

	_ struct{}
}
