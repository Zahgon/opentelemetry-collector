package component

import (
	"context"
)

type Component interface {
	Start(ctx context.Context, host Host) error

	Shutdown(ctx context.Context) error
}

type StartFunc func(context.Context, Host) error

func (f StartFunc) Start(ctx context.Context, host Host) error {
	_ = "STUB: not implemented"
	return nil
}

type ShutdownFunc func(context.Context) error

func (f ShutdownFunc) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type Kind struct {
	name string
}

var (
	KindReceiver  = Kind{name: "Receiver"}
	KindProcessor = Kind{name: "Processor"}
	KindExporter  = Kind{name: "Exporter"}
	KindExtension = Kind{name: "Extension"}
	KindConnector = Kind{name: "Connector"}
)

func (k Kind) String() string { _ = "STUB: not implemented"; return "" }

type StabilityLevel int

const (
	StabilityLevelUndefined StabilityLevel = iota
	StabilityLevelUnmaintained
	StabilityLevelDeprecated
	StabilityLevelDevelopment
	StabilityLevelAlpha
	StabilityLevelBeta
	StabilityLevelStable
)

func (sl *StabilityLevel) UnmarshalText(in []byte) error { _ = "STUB: not implemented"; return nil }

func (sl StabilityLevel) String() string { _ = "STUB: not implemented"; return "" }

func (sl StabilityLevel) LogMessage() string { _ = "STUB: not implemented"; return "" }

type Factory interface {
	Type() Type

	CreateDefaultConfig() Config
}

type CreateDefaultConfigFunc func() Config

func (f CreateDefaultConfigFunc) CreateDefaultConfig() Config {
	_ = "STUB: not implemented"
	return *new(Config)
}
