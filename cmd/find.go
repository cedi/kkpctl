package cmd

import (
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find specific things",
}

func init() {
	rootCmd.AddCommand(findCmd)
}
