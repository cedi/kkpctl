package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe an object of a specified resource type in greater detail",
}

func init() {
	rootCmd.AddCommand(describeCmd)
}
