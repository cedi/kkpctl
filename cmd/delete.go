package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Get lets you fetch a list or a specific named object.",
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
