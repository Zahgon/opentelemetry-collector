//go:generate mdatagen metadata.yaml

package processorhelper

import (
	"errors"

	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
)

var ErrSkipProcessingData = errors.New("sentinel error to skip processing data from the remainder of the pipeline")

type Option interface {
	apply(*baseSettings)
}

type optionFunc func(*baseSettings)

func (of optionFunc) apply(e *baseSettings) { _ = "STUB: not implemented"; return }

func WithStart(start component.StartFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithShutdown(shutdown component.ShutdownFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCapabilities(capabilities consumer.Capabilities) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type baseSettings struct {
	component.StartFunc
	component.ShutdownFunc
	consumerOptions []consumer.Option
}

func fromOptions(options []Option) *baseSettings { _ = "STUB: not implemented"; return nil }

func spanAttributes(id component.ID) trace.EventOption {
	_ = "STUB: not implemented"
	return *new(trace.EventOption)
}
