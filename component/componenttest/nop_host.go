package componenttest

import (
	"go.opentelemetry.io/collector/component"
)

var _ component.Host = (*nopHost)(nil)

type nopHost struct{}

func NewNopHost() component.Host { _ = "STUB: not implemented"; return *new(component.Host) }

func (nh *nopHost) GetExtensions() map[component.ID]component.Component {
	_ = "STUB: not implemented"
	return nil
}
