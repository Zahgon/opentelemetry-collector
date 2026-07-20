package otelcol

//go:generate mdatagen metadata.yaml

import (
	"flag"

	"github.com/spf13/cobra"
)

func NewCommand(set CollectorSettings) *cobra.Command { _ = "STUB: not implemented"; return nil }

func updateSettingsUsingFlags(set *CollectorSettings, flags *flag.FlagSet) error {
	_ = "STUB: not implemented"
	return nil
}

func newFeatureGateCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }
