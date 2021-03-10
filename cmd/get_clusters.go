package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var projectID string

// clustersCmd represents the clusters command
var getClustersCmd = &cobra.Command{
	Use:   "cluster [clusterid]",
	Short: "Lists clusters for a given project (and optional seed datacenter) or fetch a named cluster.",
	Long:  `If no clusterid is specified, all clusters of an project are listed. If a clusterid is specified only this cluster is shown`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		//fmt.Println("clusters called")
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.Wrap(err, "Could not initialize Kubermatic API client")
		}

		var result interface{}
		if len(args) == 0 {
			result, err = kkp.ListClusters(projectID)
		}

		if err != nil {
			return errors.Wrap(err, "Error fetching clusters")
		}

		parsed, err := output.ParseOutput(result, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "Error parsing clusters")
		}
		fmt.Println(parsed)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getClustersCmd)

	getClustersCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project to list clusters for.")
	getClustersCmd.MarkFlagRequired("project")
}
