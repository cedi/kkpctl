package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe lets you describe a KKP object in more detail.",
}

func init() {
	rootCmd.AddCommand(describeCmd)
}
