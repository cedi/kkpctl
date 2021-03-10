package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	projectID  string
	datacenter string
)

// clustersCmd represents the clusters command
var getClustersCmd = &cobra.Command{
	Use:               "cluster [clusterid]",
	Short:             "Lists clusters for a given project (and optional seed datacenter) or fetch a named cluster.",
	Long:              `If no clusterid is specified, all clusters of an project are listed. If a clusterid is specified only this cluster is shown`,
	Example:           "kkpctl get clusterr --project dw2s9jk28z",
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: getValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.Wrap(err, "Could not initialize Kubermatic API client")
		}

		var result interface{}
		if len(args) == 0 {
			if datacenter == "" {
				result, err = kkp.ListClusters(projectID)
			} else {
				result, err = kkp.ListClustersInDC(projectID, datacenter)
			}
		} else {
			if datacenter == "" {
				result, err = kkp.GetClusterProject(args[0], projectID)
			} else {
				result, err = kkp.GetCluster(args[0], projectID, datacenter)
			}
		}

		if err != nil {
			return errors.Wrap(err, "Error fetching clusters")
		}

		parsed, err := output.ParseOutput(result, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "Error parsing clusters")
		}
		fmt.Print(parsed)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getClustersCmd)
	getClustersCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project.")
	getClustersCmd.MarkFlagRequired("project")

	getClustersCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter.")
}
