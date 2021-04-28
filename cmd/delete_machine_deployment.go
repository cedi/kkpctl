package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var deleteMachineDeploymentCmd = &cobra.Command{
	Use:               "machinedeployment name",
	Short:             "Delete a machine deployment from a cluster",
	Args:              cobra.ExactArgs(1),
	Example:           "kkpctl delete machinedeployment --project 6tmbnhdl7h --cluster qvjdddt72t my_first_machine_deployment",
	ValidArgsFunction: completion.GetValidMachineDeploymentArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		machineDeploymentName := args[0]
		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		cluster, err := kkp.GetClusterInProjectInDC(clusterID, projectID, datacenter)
		if err != nil {
			return errors.Wrapf(err, "failed to find cluster %s in project %s", clusterID, projectID)
		}

		machineDeployment, err := kkp.GetMachineDeployment(machineDeploymentName, clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrapf(err, "failed to get machine deployment %s in cluster %s", machineDeploymentName, clusterID)
		}

		err = kkp.DeleteMachineDeployment(machineDeployment.ID, clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrapf(err, "failed to delete machine deployment %s (%s) in cluster %s", machineDeployment.Name, machineDeployment.ID, clusterID)
		}

		fmt.Printf("Successfully deleted machine deployment %s in cluster %s\n", machineDeployment.Name, clusterID)
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteMachineDeploymentCmd)

	deleteMachineDeploymentCmd.Flags().StringVarP(&clusterID, "cluster", "c", "", "ID of the cluster")
	deleteMachineDeploymentCmd.MarkFlagRequired("cluster")
	deleteMachineDeploymentCmd.RegisterFlagCompletionFunc("cluster", completion.GetValidClusterArgs)

	deleteMachineDeploymentCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project")
	deleteMachineDeploymentCmd.MarkFlagRequired("project")
	deleteMachineDeploymentCmd.RegisterFlagCompletionFunc("project", completion.GetValidProjectArgs)

	deleteMachineDeploymentCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter")
	deleteMachineDeploymentCmd.RegisterFlagCompletionFunc("datacenter", completion.GetValidDatacenterArgs)
}
