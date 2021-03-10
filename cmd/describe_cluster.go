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
	Short:             "Describe a cluster.",
	Example:           "kkpctl describe cluster rbw47nm2h8 --project dw2s9jk28z --datacenter ix1",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: getValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.New("Could not initialize Kubermatic API client")
		}

		var cluster models.Cluster
		if datacenter == "" {
			cluster, err = kkp.GetClusterProject(args[0], projectID)
		} else {
			cluster, err = kkp.GetCluster(args[0], projectID, datacenter)
		}

		if err != nil {
			return errors.Wrap(err, "Error fetching cluster")
		}

		nodeDeployments, err := kkp.ListNodeDeployments(cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrap(err, "Error fetching node deployments for cluster")
		}

		meta := describe.ClusterDescribeMeta{
			Cluster:         &cluster,
			NodeDeployments: nodeDeployments,
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

	describeClusterCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter to list clusters for.")
}
