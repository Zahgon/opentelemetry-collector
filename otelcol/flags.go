package otelcol

import (
	"flag"

	"go.opentelemetry.io/collector/featuregate"
)

const (
	configFlag = "config"
)

type configFlagValue struct {
	values []string
	sets   []string
}

func (s *configFlagValue) Set(val string) error { _ = "STUB: not implemented"; return nil }

func (s *configFlagValue) String() string { _ = "STUB: not implemented"; return "" }

func flags(reg *featuregate.Registry) *flag.FlagSet { _ = "STUB: not implemented"; return nil }

func getConfigFlag(flagSet *flag.FlagSet) []string { _ = "STUB: not implemented"; return nil }
