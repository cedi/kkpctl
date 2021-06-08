package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/describe"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// clustersCmd represents the clusters command
var describeMachineDeploymentCmd = &cobra.Command{
	Use:               "machinedeployment name",
	Short:             "Describes a machine deployment",
	Example:           "kkpctl describe machinedeployment --project 6tmbnhdl7h --cluster qvjdddt72t hallowelt",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completion.GetValidMachineDeploymentArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		machineDeploymentName := args[0]

		kkp, err := Config.GetKKPClient(client.V2API)
		if err != nil {
			return err
		}

		machineDeployment, err := kkp.GetMachineDeployment(machineDeploymentName, clusterID, projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to get machine deployment %s for cluster %s in project %s", machineDeploymentName, clusterID, projectID)
		}

		machineDeploymentNodes, err := kkp.GetMachineDeploymentNodes(machineDeployment.ID, clusterID, projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to get machine deployment %s's nodes for cluster %s in project %s", machineDeploymentName, clusterID, projectID)
		}

		machineDeploymentEvents, err := kkp.GetMachineDeploymentEvents(machineDeployment.ID, clusterID, projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to get machine deployment %s's events for cluster %s in project %s", machineDeploymentName, clusterID, projectID)
		}

		meta := describe.MachineDeploymentDescribeMeta{
			MachineDeployment: machineDeployment,
			Nodes:             machineDeploymentNodes,
			NodeEvents:        machineDeploymentEvents,
		}

		parsed, err := describe.Object(&meta)
		if err != nil {
			return errors.Wrapf(err, "failed to describe machine deployment %s", machineDeploymentName)
		}

		fmt.Println(parsed)
		return nil
	},
}

func init() {
	describeCmd.AddCommand(describeMachineDeploymentCmd)

	AddClusterFlag(describeMachineDeploymentCmd)
	AddProjectFlag(describeMachineDeploymentCmd)
	AddDatacenterFlag(describeMachineDeploymentCmd, false)
}
