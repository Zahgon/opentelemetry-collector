package extensioncapabilities

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/extension"
)

type Dependent interface {
	extension.Extension
	Dependencies() []component.ID
}

type PipelineWatcher interface {
	Ready() error

	NotReady() error
}

type ConfigWatcher interface {
	NotifyConfig(ctx context.Context, conf *confmap.Conf) error
}

type ConfigSnapshot interface {
	Effective() *confmap.Conf

	Unexpanded() *confmap.Conf

	unexportedConfigSnapshot()
}

type configSnapshot struct {
	effective  *confmap.Conf
	unexpanded *confmap.Conf
}

func NewConfigSnapshot(effective, unexpanded *confmap.Conf) ConfigSnapshot {
	_ = "STUB: not implemented"
	return *new(ConfigSnapshot)
}

func (cs configSnapshot) unexportedConfigSnapshot() { _ = "STUB: not implemented"; return }

func (cs configSnapshot) Effective() *confmap.Conf { _ = "STUB: not implemented"; return nil }

func (cs configSnapshot) Unexpanded() *confmap.Conf { _ = "STUB: not implemented"; return nil }

func cloneConf(conf *confmap.Conf) *confmap.Conf { _ = "STUB: not implemented"; return nil }

type ConfigSnapshotWatcher interface {
	NotifyConfigSnapshot(ctx context.Context, configSnapshot ConfigSnapshot) error
}
