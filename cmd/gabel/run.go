package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hlts2/gabel"
	"github.com/hlts2/gabel/helpers"
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

//Output file Config for the result
const (
	DirForResult   = "GabelResult"
	OutputFileName = "labeld.csv"
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

	rf, err := helpers.OpenFile(l.Path, os.O_RDONLY)
	if err != nil {
		return err
	}

	if err := helpers.Mkdir(DirForResult); err != nil {
		return err
	}

	name := filepath.Join(DirForResult, OutputFileName)
	wf, err := helpers.CreateFile(name, os.O_RDWR)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(wf)
	reader := csv.NewReader(rf)
	reader.LazyQuotes = true

	g := gabel.Gabel{
		LabelingInfo: l,
		Stdin:        os.Stdin,
	}

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
