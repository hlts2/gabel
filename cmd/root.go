package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//RootCmd is
var rootCmd = &cobra.Command{
	Use:   "gabel",
	Short: "gabel is a tool for creating teacher data",
}

//Execute Command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
