package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Lets you describe an object of a specified resource type in greater detail",
}

func init() {
	rootCmd.AddCommand(describeCmd)
}
