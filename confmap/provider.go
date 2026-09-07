package confmap

import (
	"context"

	"go.uber.org/zap"
)

type ProviderSettings struct {
	Logger *zap.Logger

	_ struct{}
}

type ProviderFactory = moduleFactory[Provider, ProviderSettings]

type CreateProviderFunc = createConfmapFunc[Provider, ProviderSettings]

func NewProviderFactory(f CreateProviderFunc) ProviderFactory {
	_ = "STUB: not implemented"
	return *new(ProviderFactory)
}

type Provider interface {
	Retrieve(ctx context.Context, uri string, watcher WatcherFunc) (*Retrieved, error)

	Scheme() string

	Shutdown(ctx context.Context) error
}

type WatcherFunc func(*ChangeEvent)

type ChangeEvent struct {
	Error error

	_ struct{}
}

type Retrieved struct {
	rawConf   any
	errorHint error
	closeFunc CloseFunc

	stringRepresentation string
	isSetString          bool
}

type retrievedSettings struct {
	errorHint            error
	stringRepresentation string
	isSetString          bool
	closeFunc            CloseFunc
}

type RetrievedOption interface {
	apply(*retrievedSettings)
}

type retrievedOptionFunc func(*retrievedSettings)

func (of retrievedOptionFunc) apply(e *retrievedSettings) { _ = "STUB: not implemented"; return }

func WithRetrievedClose(closeFunc CloseFunc) RetrievedOption {
	_ = "STUB: not implemented"
	return *new(RetrievedOption)
}

func withStringRepresentation(stringRepresentation string) RetrievedOption {
	_ = "STUB: not implemented"
	return *new(RetrievedOption)
}

func withErrorHint(errorHint error) RetrievedOption {
	_ = "STUB: not implemented"
	return *new(RetrievedOption)
}

func NewRetrievedFromYAML(yamlBytes []byte, opts ...RetrievedOption) (*Retrieved, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewRetrieved(rawConf any, opts ...RetrievedOption) (*Retrieved, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Retrieved) AsConf() (*Conf, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Retrieved) AsRaw() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (r *Retrieved) AsString() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r *Retrieved) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type CloseFunc func(context.Context) error

func checkRawConfType(rawConf any) error { _ = "STUB: not implemented"; return nil }
