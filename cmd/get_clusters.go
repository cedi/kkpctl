package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pID string
var seed string

// clustersCmd represents the clusters command
var getClustersCmd = &cobra.Command{
	Use:   "cluster [id]",
	Short: "Lists clusters for a given project (and optional seed datacenter) or fetch a named cluster.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not yet implemented")
		fmt.Println(baseURL)
	},
}

func init() {
	getCmd.AddCommand(getClustersCmd)
}
