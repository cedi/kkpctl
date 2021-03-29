package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/describe"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// clustersCmd represents the clusters command
var describeNodeDeploymentCmd = &cobra.Command{
	Use:               "nodedeployment name",
	Short:             "Describes a node deployment",
	Example:           "kkpctl describe nodedeployment --project 6tmbnhdl7h --cluster qvjdddt72t hallowelt",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: getValidNodeDeploymentArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		nodeDeploymentName := args[0]

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		cluster, err := kkp.GetClusterInProjectInDC(clusterID, projectID, datacenter)

		if err != nil {
			return errors.Wrapf(err, "failed to get cluster %s in project %s", clusterID, projectID)
		}

		nodeDeployment, err := kkp.GetNodeDeployment(nodeDeploymentName, cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrapf(err, "failed to get node deployment %s for cluster %s in project %s", nodeDeploymentName, clusterID, projectID)
		}

		nodeDeploymentNodes, err := kkp.GetNodeDeploymentNodes(nodeDeployment.ID, cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrapf(err, "failed to get node deployment %s's nodes for cluster %s in project %s", nodeDeploymentName, clusterID, projectID)
		}

		nodeDeploymentEvents, err := kkp.GetNodeDeploymentEvents(nodeDeployment.ID, cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrapf(err, "failed to get node deployment %s's events for cluster %s in project %s", nodeDeploymentName, clusterID, projectID)
		}

		meta := describe.NodeDeploymentDescribeMeta{
			NodeDeployment: nodeDeployment,
			Nodes:          nodeDeploymentNodes,
			NodeEvents:     nodeDeploymentEvents,
		}

		parsed, err := describe.Object(&meta)
		if err != nil {
			return errors.Wrapf(err, "failed to describe node deployment %s", nodeDeploymentName)
		}

		fmt.Println(parsed)
		return nil
	},
}

func init() {
	describeCmd.AddCommand(describeNodeDeploymentCmd)

	describeNodeDeploymentCmd.Flags().StringVarP(&clusterID, "cluster", "c", "", "ID of the cluster")
	describeNodeDeploymentCmd.MarkFlagRequired("cluster")
	describeNodeDeploymentCmd.RegisterFlagCompletionFunc("cluster", getValidClusterArgs)

	describeNodeDeploymentCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project")
	describeNodeDeploymentCmd.MarkFlagRequired("project")
	describeNodeDeploymentCmd.RegisterFlagCompletionFunc("project", getValidProjectArgs)

	describeNodeDeploymentCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter")
	describeNodeDeploymentCmd.RegisterFlagCompletionFunc("datacenter", getValidDatacenterArgs)
}
