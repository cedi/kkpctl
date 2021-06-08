package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/cedi/kkpctl/pkg/client"
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
		kkp, err := Config.GetKKPClient(client.V2API)
		if err != nil {
			return err
		}

		machineDeployment, err := kkp.GetMachineDeployment(machineDeploymentName, clusterID, projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to get machine deployment %s in cluster %s", machineDeploymentName, clusterID)
		}

		err = kkp.DeleteMachineDeployment(machineDeployment.ID, clusterID, projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to delete machine deployment %s (%s) in cluster %s", machineDeployment.Name, machineDeployment.ID, clusterID)
		}

		fmt.Printf("Successfully deleted machine deployment %s in cluster %s\n", machineDeployment.Name, clusterID)
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteMachineDeploymentCmd)

	AddClusterFlag(deleteMachineDeploymentCmd)
	AddProjectFlag(deleteMachineDeploymentCmd)
	AddDatacenterFlag(deleteMachineDeploymentCmd, false)
}
