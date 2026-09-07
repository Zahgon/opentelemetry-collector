package confmap

import (
	"context"
	"regexp"
)

var driverLetterRegexp = regexp.MustCompile("^[A-z]:")

type Resolver struct {
	uris          []location
	providers     map[string]Provider
	defaultScheme string
	converters    []Converter

	closers []CloseFunc
	watcher chan error

	unexpandedConfMap map[string]any
}

type ResolverSettings struct {
	URIs []string

	ProviderFactories []ProviderFactory

	DefaultScheme string

	ProviderSettings ProviderSettings

	ConverterFactories []ConverterFactory

	ConverterSettings ConverterSettings

	_ struct{}
}

func NewResolver(set ResolverSettings) (*Resolver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *Resolver) Resolve(ctx context.Context) (*Conf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *Resolver) UnexpandedConf() *Conf { _ = "STUB: not implemented"; return nil }

func escapeDollarSigns(val any) any { _ = "STUB: not implemented"; return *new(any) }

func (mr *Resolver) Watch() <-chan error { _ = "STUB: not implemented"; return nil }

func (mr *Resolver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (mr *Resolver) onChange(event *ChangeEvent) { _ = "STUB: not implemented"; return }

func (mr *Resolver) closeIfNeeded(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (mr *Resolver) retrieveValue(ctx context.Context, uri location) (*Retrieved, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
