package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Lets you delete an object of a specified resource type",
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
