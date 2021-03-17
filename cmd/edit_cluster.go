package cmd

import (
	"github.com/spf13/cobra"
)

var editClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Lets you edit a cluster",
}

func init() {
	editCmd.AddCommand(editClusterCmd)
}
