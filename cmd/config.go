package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure kkpctl",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
