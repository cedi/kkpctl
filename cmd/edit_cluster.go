package cmd

import (
	"github.com/spf13/cobra"
)

var editClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Edit a cluster",
}

func init() {
	editCmd.AddCommand(editClusterCmd)
}
