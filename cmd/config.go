package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Lets you configure KKPCTL",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
