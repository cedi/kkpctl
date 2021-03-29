package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var deleteNodeDeploymentCmd = &cobra.Command{
	Use:               "nodedeployment name",
	Short:             "Delete a node deployment from a cluster",
	Args:              cobra.ExactArgs(1),
	Example:           "kkpctl delete nodedeployment --project 6tmbnhdl7h --cluster qvjdddt72t my_first_nodedeployment",
	ValidArgsFunction: completion.GetValidNodeDeploymentArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		nodeDeploymentName := args[0]
		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		cluster, err := kkp.GetClusterInProjectInDC(clusterID, projectID, datacenter)
		if err != nil {
			return errors.Wrapf(err, "failed to find cluster %s in project %s", clusterID, projectID)
		}

		nodeDeployment, err := kkp.GetNodeDeployment(nodeDeploymentName, clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrapf(err, "failed to get node deployment %s in cluster %s", nodeDeploymentName, clusterID)
		}

		err = kkp.DeleteNodeDeployment(nodeDeployment.ID, clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrapf(err, "failed to delete node deployment %s (%s) in cluster %s", nodeDeployment.Name, nodeDeployment.ID, clusterID)
		}

		fmt.Printf("Successfully deleted node deployment %s in cluster %s\n", nodeDeployment.Name, clusterID)
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteNodeDeploymentCmd)

	deleteNodeDeploymentCmd.Flags().StringVarP(&clusterID, "cluster", "c", "", "ID of the cluster")
	deleteNodeDeploymentCmd.MarkFlagRequired("cluster")
	deleteNodeDeploymentCmd.RegisterFlagCompletionFunc("cluster", completion.GetValidClusterArgs)

	deleteNodeDeploymentCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project")
	deleteNodeDeploymentCmd.MarkFlagRequired("project")
	deleteNodeDeploymentCmd.RegisterFlagCompletionFunc("project", completion.GetValidProjectArgs)

	deleteNodeDeploymentCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter")
	deleteNodeDeploymentCmd.RegisterFlagCompletionFunc("datacenter", completion.GetValidDatacenterArgs)
}
