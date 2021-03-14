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
	Short:             "Lists clusters for a given project or datacenter",
	Example:           "kkpctl get cluster --project dw2s9jk28z",
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: getValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, apiToken := Config.GetCloudFromContext()
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.Wrap(err, "Could not initialize Kubermatic API client")
		}

		var result interface{}
		if len(args) == 0 || listAll {
			if datacenter == "" && projectID == "" {
				result, err = kkp.ListClusters(listAll)
			} else if datacenter == "" && projectID != "" {
				result, err = kkp.ListClustersInProject(projectID)
			} else if datacenter != "" && projectID == "" {
				result, err = kkp.ListClustersInDC(datacenter, listAll)
			} else if datacenter != "" && projectID != "" {
				result, err = kkp.ListClustersInProjectInDC(projectID, datacenter)
			}
		} else {
			if datacenter == "" && projectID == "" {
				result, err = kkp.GetCluster(args[0], listAll)
			} else if datacenter == "" && projectID != "" {
				result, err = kkp.GetClusterInProject(args[0], projectID)
			} else if datacenter != "" && projectID == "" {
				result, err = kkp.GetClusterInDC(args[0], datacenter, listAll)
			} else if datacenter != "" && projectID != "" {
				result, err = kkp.GetClusterInProjectInDC(args[0], projectID, datacenter)
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
	getClustersCmd.RegisterFlagCompletionFunc("project", getValidProjectArgs)

	getClustersCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter.")
	getClustersCmd.RegisterFlagCompletionFunc("datacenter", getValidDatacenterArgs)

	getClustersCmd.Flags().BoolVarP(&listAll, "all", "a", false, "To list all clusters in all projects if the users is allowed to see.")
}
