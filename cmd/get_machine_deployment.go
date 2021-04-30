package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// clustersCmd represents the clusters command
var getMachineDeploymentCmd = &cobra.Command{
	Use:               "machinedeployment [name]",
	Short:             "List machinedeployments for a cluster",
	Example:           "kkpctl describe machinedeployment my_machinedeployment",
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: completion.GetValidMachineDeploymentArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		machineDeploymentName := ""
		if len(args) == 1 {
			machineDeploymentName = args[0]
		}

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		cluster, err := kkp.GetClusterInProjectInDC(clusterID, projectID, datacenter)
		if err != nil {
			return errors.Wrapf(err, "failed to get cluster %s in project %s", clusterID, projectID)
		}

		var machineDeployments interface{}
		if len(args) == 0 {
			machineDeployments, err = kkp.GetMachineDeployments(clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		} else {
			machineDeployments, err = kkp.GetMachineDeployment(machineDeploymentName, clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		}

		if err != nil {
			return errors.Wrapf(err, "failed to get machine deployment %s for cluster %s in project %s", machineDeploymentName, clusterID, projectID)
		}

		parsed, err := output.ParseOutput(machineDeployments, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "failed to parse machine deployment")
		}

		fmt.Print(parsed)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getMachineDeploymentCmd)

	AddProjectFlag(getMachineDeploymentCmd)
	AddDatacenterFlag(getMachineDeploymentCmd, false)
	AddClusterFlag(getMachineDeploymentCmd)
}
