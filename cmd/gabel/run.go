package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/hlts2/gabel"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run labeling tool",
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(args); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var (
	configPath string
)

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().StringVarP(&configPath, "set", "s", "", "set config file")
}

func run(args []string) error {
	if configPath == "" {
		return errors.New(errorTmplOfConfigFileDoesNotExist)
	}

	var l gabel.LabelingInfo
	err := gabel.LoadLabelingInfoWithGivenConfigPath(configPath, &l)
	if err != nil {
		return err
	}

	gio, err := gabel.NewGIO(l.Path)
	if err != nil {
		return err
	}

	defer gio.Closes()

	g := gabel.Gabel{
		LabelingInfo: l,
		GIO:          gio,
	}

	writer := csv.NewWriter(gio.WFile)
	reader := csv.NewReader(gio.RFile)
	reader.LazyQuotes = true

	return g.Run(reader, writer)
}

var errorTmplOfConfigFileDoesNotExist = `Error: config file does not exist
Usage:
   gabel run [flags]

Flags:
   -h, --help  Help for Run Command
               $ gabel run --help
   -s, --set   Set config file path for gabel
               $ gabel run -s config.yaml
	`
