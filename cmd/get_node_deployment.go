package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// clustersCmd represents the clusters command
var getNodeDeploymentCmd = &cobra.Command{
	Use:     "nodedeployment [name]",
	Short:   "Lists all available datacenters",
	Example: "kkpctl get datacenter",
	Args:    cobra.MaximumNArgs(1),
	//ValidArgsFunction: getValidNodeDeploymentArgs,
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

		nodeDeployments := make([]models.NodeDeployment, 0)
		if len(args) == 0 {
			nodeDeployments, err = kkp.GetNodeDeployments(clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		} else {
			var nodeDeployment models.NodeDeployment
			nodeDeployment, err = kkp.GetNodeDeployment(args[0], clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
			nodeDeployments = append(nodeDeployments, nodeDeployment)
		}

		if err != nil {
			return errors.Wrap(err, "Error fetching node deployment")
		}

		parsed, err := output.ParseOutput(nodeDeployments, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "Error parsing datacenters")
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
