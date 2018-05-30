package cmd

import (
	"fmt"
	"os"

	"github.com/hlts2/gabel/pkg/gabel"
	"github.com/spf13/cobra"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "gabel",
	Short: "gabel is a CLI tool for creating teacher data",
	RunE:  Run,
}

// Run executes gabel command
func Run(cmd *cobra.Command, args []string) error {
	var config *gabel.Config

	err := gabel.LoadConfig(config, configPath)
	if err != nil {
		return err
	}

	csv, err := gabel.NewCSV(configPath)
	if err != nil {
		return err
	}

	g, err := gabel.NewGabel(os.Stdin, os.Stdout, config, csv, func() string {
		return ""
	})
	if err != nil {
		return err
	}

	return g.Run(0, len(csv.Records))
}

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
