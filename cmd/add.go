package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Lets you add an object of a specified resource type",
}

func init() {
	rootCmd.AddCommand(addCmd)
}
