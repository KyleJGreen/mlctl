package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "mlctl",
	Short: "mlctl - a machine learning toolkit",
	Long:  `mlctl provides starter code for common ML and data engineering utilities`,
	Run: func(cmd *cobra.Command, args []string) {
		// This is the default action when no subcommands are specified.
		fmt.Println("Please specify a command!")
	},
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
