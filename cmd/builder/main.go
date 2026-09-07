package main

import (
	"os"

	"github.com/spf13/cobra"

	"go.opentelemetry.io/collector/cmd/builder/internal"
)

func main() {
	cmd, err := internal.Command()
	cobra.CheckErr(err)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
