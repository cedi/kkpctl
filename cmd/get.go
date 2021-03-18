package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an object of a specified resource type",
}

func init() {
	rootCmd.AddCommand(getCmd)
}
