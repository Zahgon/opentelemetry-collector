package featuregate

import (
	"flag"
)

const (
	featureGatesFlag            = "feature-gates"
	featureGatesFlagDescription = "Comma-delimited list of feature gate identifiers. Prefix with '-' to disable the feature. '+' or no prefix will enable the feature."
)

type RegisterFlagsOption interface {
	private()
}

func (r *Registry) RegisterFlags(flagSet *flag.FlagSet, _ ...RegisterFlagsOption) {
	_ = "STUB: not implemented"
	return
}

type flagValue struct {
	reg *Registry
}

func (f *flagValue) String() string { _ = "STUB: not implemented"; return "" }

func (f *flagValue) Set(s string) error { _ = "STUB: not implemented"; return nil }
