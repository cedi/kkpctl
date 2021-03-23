package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/errors"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/spf13/cobra"
)

// clustersCmd represents the clusters command
var getNodeDeploymentCmd = &cobra.Command{
	Use:               "nodedeployment [name]",
	Short:             "List nodedeployments for a cluster",
	Example:           "kkpctl describe nodedeployment my_nodedeployment",
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: getValidNodeDeploymentArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		nodeDeploymentName := ""
		if len(args) == 1 {
			nodeDeploymentName = args[0]
		}

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		cluster, err := kkp.GetClusterInProjectInDC(clusterID, projectID, datacenter)
		if err != nil {
			return errors.Wrapf(err, "failed to get cluster %s in project", clusterID, projectID)
		}

		var nodeDeployments interface{}
		if len(args) == 0 {
			nodeDeployments, err = kkp.GetNodeDeployments(clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		} else {
			nodeDeployments, err = kkp.GetNodeDeployment(nodeDeploymentName, clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		}

		if err != nil {
			return errors.Wrapf(err, "failed to get node deployment %s for cluster %s in project %s", nodeDeploymentName, clusterID, projectID)
		}

		parsed, err := output.ParseOutput(nodeDeployments, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "failed to parse node deployment")
		}

		fmt.Print(parsed)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getNodeDeploymentCmd)

	getNodeDeploymentCmd.Flags().StringVarP(&clusterID, "cluster", "c", "", "ID of the cluster")
	getNodeDeploymentCmd.MarkFlagRequired("cluster")
	getNodeDeploymentCmd.RegisterFlagCompletionFunc("cluster", getValidClusterArgs)

	getNodeDeploymentCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project")
	getNodeDeploymentCmd.MarkFlagRequired("project")
	getNodeDeploymentCmd.RegisterFlagCompletionFunc("project", getValidProjectArgs)

	getNodeDeploymentCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter")
	getNodeDeploymentCmd.RegisterFlagCompletionFunc("datacenter", getValidDatacenterArgs)
}
