package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hlts2/gabel"
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

//OutputFilePath is the output file for the result
const (
	OutputFilePath = "res/labeld.csv"
)

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().StringVarP(&configPath, "set", "s", "", "set config file")
}

func run(args []string) error {
	if configPath == "" {
		return errorUsage()
	}

	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	var l gabel.LabelingInfo
	if err = yaml.Unmarshal(b, &l); err != nil {
		return err
	}

	rFile, err := os.OpenFile(l.Path, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}

	wFile, err := os.OpenFile(l.Path, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		rFile.Close()
		return err
	}

	c := &gabel.Config{
		LabelingInfo: l,
		Stdin:        os.Stdin,
		RFile:        rFile,
		WFile:        wFile,
	}

	return c.Run()
}

func errorUsage() error {
	return errors.New(`Error: config file does not exist
Usage:
   gabel run <option>

Available Options:
   -s, --set   Set config file path for gabel
               $ gabel run -s config.yaml
	`)
}
