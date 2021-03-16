package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var deleteNodeDeploymentCmd = &cobra.Command{
	Use:     "nodedeployment name",
	Short:   "Lets you create a new node deployment",
	Args:    cobra.ExactArgs(1),
	Example: "kkpctl delete nodedeployment --project 6tmbnhdl7h --cluster qvjdddt72t",
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, apiToken := Config.GetCloudFromContext()
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.New("Could not initialize Kubermatic API client")
		}

		var cluster models.Cluster
		if datacenter == "" {
			cluster, err = kkp.GetClusterInProject(clusterID, projectID)
		} else if datacenter != "" && projectID != "" {
			cluster, err = kkp.GetClusterInProjectInDC(clusterID, projectID, datacenter)
		}

		if err != nil {
			return errors.Wrap(err, "could not fetch cluster")
		}

		nodeDeployment, err := kkp.GetNodeDeployment(args[0], clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrap(err, "Error fetching node deployment")
		}

		err = kkp.DeleteNodeDeployment(nodeDeployment.ID, clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrap(err, "Error deleting node deployment")
		}

		fmt.Printf("Successfully deleted node deployment '%s'\n", nodeDeployment.Name)
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteNodeDeploymentCmd)

	deleteNodeDeploymentCmd.Flags().StringVarP(&clusterID, "cluster", "c", "", "ID of the cluster")
	deleteNodeDeploymentCmd.MarkFlagRequired("cluster")
	deleteNodeDeploymentCmd.RegisterFlagCompletionFunc("cluster", getValidClusterArgs)

	deleteNodeDeploymentCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project")
	deleteNodeDeploymentCmd.MarkFlagRequired("project")
	deleteNodeDeploymentCmd.RegisterFlagCompletionFunc("project", getValidProjectArgs)

	deleteNodeDeploymentCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter")
	deleteNodeDeploymentCmd.RegisterFlagCompletionFunc("datacenter", getValidDatacenterArgs)
}
