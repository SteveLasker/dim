package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "dim [command]",
		SilenceUsage: true,
	}
	cmd.AddCommand(loginCmd(), pullCmd(), pushCmd())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
