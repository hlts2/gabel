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

	var c gabel.Config
	if err = yaml.Unmarshal(b, &c); err != nil {
		return err
	}

	fmt.Println(c)

	return nil
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
