package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/describe"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var describeClusterCmd = &cobra.Command{
	Use:               "cluster clusterid",
	Short:             "Describe a cluster",
	Example:           "kkpctl describe cluster rbw47nm2h8 --project dw2s9jk28z",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: getValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, apiToken := Config.GetCloudFromContext()
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.New("Could not initialize Kubermatic API client")
		}

		var cluster models.Cluster
		if datacenter == "" && projectID == "" {
			cluster, err = kkp.GetCluster(args[0], listAll)
		} else if datacenter == "" && projectID != "" {
			cluster, err = kkp.GetClusterInProject(args[0], projectID)
		} else if datacenter != "" && projectID == "" {
			cluster, err = kkp.GetClusterInDC(args[0], datacenter, listAll)
		} else if datacenter != "" && projectID != "" {
			cluster, err = kkp.GetClusterInProjectInDC(args[0], projectID, datacenter)
		}

		if err != nil {
			return errors.Wrap(err, "Error fetching cluster")
		}

		nodeDeployments, err := kkp.GetNodeDeployments(cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			// If we couldn't fetch the NodeDeployments, that doesn't bother me, just use a empty array
			nodeDeployments = make([]models.NodeDeployment, 0)
		}

		clusterHealth, err := kkp.GetClusterHealth(cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrap(err, "Error fetching cluster health status")
		}

		meta := describe.ClusterDescribeMeta{
			Cluster:         &cluster,
			NodeDeployments: nodeDeployments,
			ClusterHealth:   clusterHealth,
		}

		parsed, err := describe.Object(&meta)
		if err != nil {
			return errors.Wrap(err, "Error describing cluster")
		}
		fmt.Print(parsed)

		return nil
	},
}

func init() {
	describeCmd.AddCommand(describeClusterCmd)

	describeClusterCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project to list clusters for.")
	describeClusterCmd.MarkFlagRequired("project")
	describeClusterCmd.RegisterFlagCompletionFunc("project", getValidProjectArgs)

	describeClusterCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter to list clusters for.")
	describeClusterCmd.RegisterFlagCompletionFunc("datacenter", getValidDatacenterArgs)

	describeClusterCmd.Flags().BoolVarP(&listAll, "all", "a", false, "To list all clusters in all projects if the users is allowed to see.")
}
