package cmd

import (
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an object of a specified resource type",
}

func init() {
	rootCmd.AddCommand(editCmd)
}
