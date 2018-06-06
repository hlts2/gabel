package cmd

import (
	"fmt"
	"os"

	"github.com/hlts2/gabel/pkg/gabel"
	"github.com/spf13/cobra"
)

var configPath string
var outputFileN string

var rootCmd = &cobra.Command{
	Use:   "gabel",
	Short: "gabel is a CLI tool for creating teacher data",
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(cmd, args); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

// Run executes gabel command
func run(cmd *cobra.Command, args []string) error {
	var config gabel.Config

	err := gabel.LoadConfig(&config, configPath)
	if err != nil {
		return err
	}

	csv, err := gabel.NewCSV(config.Path)
	if err != nil {
		return err
	}

	sw := gabel.NewScanWriter(os.Stdin, os.Stdout)

	g, err := gabel.NewGabel(sw, config, csv, func() string {
		return "\"{{index $ 0}}\"\n"
	})
	if err != nil {
		return err
	}

	err = g.Run(0, len(csv.Records))
	if err != nil {
		return err
	}

	err = g.WriteCSV(outputFileN)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "set", "s", "", "set config file path")
	rootCmd.PersistentFlags().StringVarP(&outputFileN, "output", "o", "output.csv", "set output file path")
}

// Execute executes the CLI application
func Execute() {
	rootCmd.Execute()
}
