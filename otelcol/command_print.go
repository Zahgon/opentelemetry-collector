package otelcol

import (
	"flag"
	"io"

	"github.com/spf13/cobra"
)

func newConfigPrintSubCommand(set CollectorSettings, flagSet *flag.FlagSet) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

type printContext struct {
	cmd          *cobra.Command
	stdout       io.Writer
	set          CollectorSettings
	outputFormat string
	validate     bool
}

func (pctx *printContext) configPrintSubCommand(flagSet *flag.FlagSet, mode string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pctx *printContext) printConfigData(data map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (pctx *printContext) getPrintableConfig() (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (pctx *printContext) printUnredactedConfig() error { _ = "STUB: not implemented"; return nil }

func (pctx *printContext) printRedactedConfig() error { _ = "STUB: not implemented"; return nil }
