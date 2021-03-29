package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/cmd/completion"
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
	ValidArgsFunction: completion.GetValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		clusterID := args[0]

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		cluster, err := kkp.GetClusterInProjectInDC(clusterID, projectID, datacenter)
		if err != nil || cluster.Spec == nil || cluster.Spec.Cloud == nil {
			return errors.Wrapf(err, "failed to get cluster %s in project %s", clusterID, projectID)
		}

		nodeDeployments, err := kkp.GetNodeDeployments(cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			// If we couldn't fetch the NodeDeployments that shouldn't bother us, just use a empty array instead
			nodeDeployments = make([]models.NodeDeployment, 0)
		}

		clusterHealth, err := kkp.GetClusterHealth(cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrapf(err, "failed to get health status for cluster %s in project %s", clusterID, projectID)
		}

		clusterEvents, err := kkp.GetEvents(cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			// If we couldn't fetch the Events that shouldn't bother us, just use a empty array instead
			clusterEvents = make([]models.Event, 0)
		}

		// this meta object contains all the information about one specific cluster
		meta := describe.ClusterDescribeMeta{
			Cluster:         cluster,
			NodeDeployments: nodeDeployments,
			ClusterHealth:   clusterHealth,
			ClusterEvents:   clusterEvents,
		}

		parsed, err := describe.Object(&meta)
		if err != nil {
			return errors.Wrapf(err, "failed to describe cluster %s in project %s", clusterID, projectID)
		}
		fmt.Println(parsed)

		return nil
	},
}

func init() {
	describeCmd.AddCommand(describeClusterCmd)

	describeClusterCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project to list clusters for.")
	describeClusterCmd.MarkFlagRequired("project")
	describeClusterCmd.RegisterFlagCompletionFunc("project", completion.GetValidProjectArgs)

	describeClusterCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter to list clusters for.")
	describeClusterCmd.RegisterFlagCompletionFunc("datacenter", completion.GetValidDatacenterArgs)

	describeClusterCmd.Flags().BoolVarP(&listAll, "all", "a", false, "To list all clusters in all projects if the users is allowed to see.")
}
