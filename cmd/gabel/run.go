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
		return runError()
	}

	var l gabel.LabelingInfo
	err := gabel.LoadLabelingInfoWithGivenConfigPath(configPath, &l)
	if err != nil {
		return err
	}

	gio, err := gabel.NewGabelio(l.Path)
	if err != nil {
		return err
	}

	defer func() {
		gio.FilesClose()
	}()

	g := gabel.Gabel{
		LabelingInfo: l,
		Gabelio:      gio,
	}

	writer := csv.NewWriter(g.WFile)
	reader := csv.NewReader(g.RFile)
	reader.LazyQuotes = true

	return g.Run(reader, writer)
}

func runError() error {
	return errors.New(`Error: config file does not exist
Usage:
   gabel run <option>

Available Options:
   -s, --set   Set config file path for gabel
               $ gabel run -s config.yaml
	`)
}
