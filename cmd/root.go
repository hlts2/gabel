package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gabel",
	Short: "gabel is a CLI tool for creating teacher data",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var helpText = `
Usage:
   gabel <option>

Flags:
   -h, --help  Help for Run Command
               $ gabel --help
   -s, --set   Set config file path
               $ gabel -s config.yaml
	`

// Execute executes the CLI application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
