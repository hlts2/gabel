package cmd

import (
	"fmt"
	"os"

	"github.com/hlts2/gabel/pkg/gabel"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gabel",
	Short: "gabel is a CLI tool for creating teacher data",
	RunE: func(cmd *cobra.Command, args []string) error {

		_, err := gabel.NewCSV(configPath)
		if err != nil {
			return err
		}

		return nil
	},
}

var configPath string

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "set", "s", "", "set config file path")
}

// Execute executes the CLI application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
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
