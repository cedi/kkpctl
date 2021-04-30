package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/cedi/kkpctl/pkg/model"
	"github.com/cedi/kkpctl/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	dynamicConfig   bool
	clusterID       string
	nodeReplica     int32
	operatingSystem string
	nodeSpecName    string
)

var createMachineDeploymentCmd = &cobra.Command{
	Use:     "machinedeployment name",
	Short:   "Lets you create a new machine deployment",
	Args:    cobra.ExactArgs(1),
	Example: "kkpctl add machinedeployment --project 6tmbnhdl7h --cluster qvjdddt72t --nodespec flatcar-m1micro --operatingsystem flatcar --provider optimist --labels \"size=micro\" my_node_deployment --replica 3",
	RunE: func(cmd *cobra.Command, args []string) error {
		machineDeploymentName := args[0]

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		cluster, err := kkp.GetClusterInProjectInDC(clusterID, projectID, datacenter)
		if err != nil {
			return errors.Wrapf(err, "failed to find cluster %s in project %s to add a machine deployment", clusterID, projectID)
		}

		clusterVersion, ok := cluster.Spec.Version.(string)
		if !ok {
			return fmt.Errorf("fatal: cluster version does not appear to be a string")
		}

		newNodeDp := model.NewMachineDeployment(
			machineDeploymentName,
			clusterVersion,
			nodeReplica,
			dynamicConfig,
			Config.NodeSpec.GetNodeCloudSpec(nodeSpecName),
			Config.OSSpec.GetOperatingSystemSpec(),
			utils.SplitLabelString(labels),
		)

		nodeDp, err := kkp.CreateMachineDeployment(newNodeDp, clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrapf(err, "unable to create machine deployment %s for cluster %s in project %s", machineDeploymentName, clusterID, projectID)
		}

		fmt.Printf("Successfully created machine deployment '%s' for cluster %s in project %s\n", nodeDp.ID, clusterID, projectID)
		return nil
	},
}

func init() {
	addCmd.AddCommand(createMachineDeploymentCmd)

	AddProjectFlag(createMachineDeploymentCmd)
	AddDatacenterFlag(createMachineDeploymentCmd, false)
	AddLabelsFlag(createMachineDeploymentCmd)
	AddClusterFlag(createMachineDeploymentCmd)

	createMachineDeploymentCmd.Flags().StringVar(&providerName, "provider", "", "Which provider should be used")
	createMachineDeploymentCmd.RegisterFlagCompletionFunc("provider", completion.GetValidProvider)
	createMachineDeploymentCmd.MarkFlagRequired("provider")

	createMachineDeploymentCmd.Flags().StringVar(&operatingSystem, "operatingsystem", "", "Which operating system should be used")
	createMachineDeploymentCmd.RegisterFlagCompletionFunc("operatingsystem", completion.GetValidOperatingSystem)
	createMachineDeploymentCmd.MarkFlagRequired("operatingsystem")

	createMachineDeploymentCmd.Flags().StringVar(&nodeSpecName, "nodespec", "", "Which node spec should be used")
	createMachineDeploymentCmd.RegisterFlagCompletionFunc("nodespec", completion.GetValidNodeSpecArgs)
	createMachineDeploymentCmd.MarkFlagRequired("nodespec")

	createMachineDeploymentCmd.Flags().BoolVar(&dynamicConfig, "dynamic_config", false, "Dynamic kubelet config")
	createMachineDeploymentCmd.Flags().Int32Var(&nodeReplica, "replica", 1, "Number of node replicas")
}
