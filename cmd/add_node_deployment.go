package cmd

import (
	"fmt"

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

// projectCmd represents the project command
var createNodeDeploymentCmd = &cobra.Command{
	Use:     "nodedeployment name",
	Short:   "Lets you create a new node deployment",
	Args:    cobra.ExactArgs(1),
	Example: "kkpctl add nodedeployment --project 6tmbnhdl7h --cluster qvjdddt72t --nodespec flatcar-m1micro --operatingsystem flatcar --provider optimist --labels \"size=micro\" my_node_deployment --replica 3",
	RunE: func(cmd *cobra.Command, args []string) error {
		nodeDeploymentName := args[0]

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		cluster, err := kkp.GetClusterInProjectInDC(clusterID, projectID, datacenter)
		if err != nil {
			return errors.Wrapf(err, "failed to find cluster %s in project %s to add a node deployment", clusterID, projectID)
		}

		clusterVersion, ok := cluster.Spec.Version.(string)
		if !ok {
			return fmt.Errorf("fatal: cluster version does not appear to be a string")
		}

		newNodeDp := model.NewNodeDeployment(
			nodeDeploymentName,
			clusterVersion,
			nodeReplica,
			dynamicConfig,
			Config.NodeSpec.GetNodeCloudSpec(nodeSpecName),
			Config.OSSpec.GetOperatingSystemSpec(),
			utils.SplitLabelString(labels),
		)

		nodeDp, err := kkp.CreateNodeDeployment(newNodeDp, clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrapf(err, "unable to create node deployment %s for cluster %s in project %s", nodeDeploymentName, clusterID, projectID)
		}

		fmt.Printf("Successfully created node deployment '%s' for cluster %s in project %s\n", nodeDp.ID, clusterID, projectID)
		return nil
	},
}

func init() {
	addCmd.AddCommand(createNodeDeploymentCmd)

	createNodeDeploymentCmd.Flags().StringVarP(&clusterID, "cluster", "c", "", "ID of the cluster")
	createNodeDeploymentCmd.MarkFlagRequired("cluster")
	createNodeDeploymentCmd.RegisterFlagCompletionFunc("cluster", getValidClusterArgs)

	createNodeDeploymentCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project")
	createNodeDeploymentCmd.MarkFlagRequired("project")
	createNodeDeploymentCmd.RegisterFlagCompletionFunc("project", getValidProjectArgs)

	createNodeDeploymentCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter")
	createNodeDeploymentCmd.RegisterFlagCompletionFunc("datacenter", getValidDatacenterArgs)

	createNodeDeploymentCmd.Flags().StringVar(&providerName, "provider", "", "Which provider should be used")
	createNodeDeploymentCmd.RegisterFlagCompletionFunc("provider", getValidProvider)
	createNodeDeploymentCmd.MarkFlagRequired("provider")

	createNodeDeploymentCmd.Flags().StringVar(&operatingSystem, "operatingsystem", "", "Which operating system should be used")
	createNodeDeploymentCmd.RegisterFlagCompletionFunc("operatingsystem", getValidOperatingSystem)
	createNodeDeploymentCmd.MarkFlagRequired("operatingsystem")

	createNodeDeploymentCmd.Flags().StringVar(&nodeSpecName, "nodespec", "", "Which node spec should be used")
	createNodeDeploymentCmd.RegisterFlagCompletionFunc("nodespec", getValidNodeSpecArgs)
	createNodeDeploymentCmd.MarkFlagRequired("nodespec")

	createNodeDeploymentCmd.Flags().BoolVar(&dynamicConfig, "dynamic_config", false, "Dynamic kubelet config")
	createNodeDeploymentCmd.Flags().Int32Var(&nodeReplica, "replica", 1, "Number of node replicas")
	createNodeDeploymentCmd.Flags().StringVarP(&labels, "labels", "l", "", "A comma separated list of labels in the format key=value")

}
