package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jr",
	Short: "jr is a cli journal.",
	Long:  "jr is a command-line journal.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing jr '%s'\n", err)
		os.Exit(1)
	}
}
