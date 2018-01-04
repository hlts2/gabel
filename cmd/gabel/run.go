package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hlts2/gabel"
	"github.com/hlts2/gabel/helpers"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
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

	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	var l gabel.LabelingInfo
	if err = yaml.Unmarshal(b, &l); err != nil {
		return err
	}

	if err := helpers.Mkdir(gabel.DirForResult); err != nil {
		return err
	}

	c := &gabel.Config{
		LabelingInfo: l,
		Stdin:        os.Stdin,
	}

	rF, err := helpers.OpenFile(c.Path, os.O_RDONLY)
	if err != nil {
		return err
	}

	name := filepath.Join(gabel.DirForResult, gabel.OutputFileName)
	wF, err := helpers.CreateFile(name, os.O_WRONLY)
	if err != nil {
		rF.Close()
		return err
	}

	defer func() {
		wF.Close()
		rF.Close()
	}()

	writer := csv.NewWriter(wF)
	reader := csv.NewReader(rF)
	reader.LazyQuotes = true

	return c.Run(reader, writer)
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
