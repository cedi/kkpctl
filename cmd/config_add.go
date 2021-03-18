package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var configAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a specific configuration object",
}

func init() {
	configCmd.AddCommand(configAddCmd)
}
