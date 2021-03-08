package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get lets you fetch a list or a specific named object.",
}

func init() {
	rootCmd.AddCommand(getCmd)
}
