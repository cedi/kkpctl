package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/describe"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// clustersCmd represents the clusters command
var describeNodeDeploymentCmd = &cobra.Command{
	Use:               "nodedeployment name",
	Short:             "Describes a node deployment",
	Example:           "kkpctl describe nodedeployment my_nodedeployment",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: getValidNodeDeploymentArgs,
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

		nodeDeployment, err := kkp.GetNodeDeployment(args[0], cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrap(err, "Error fetching node deployment")
		}

		nodeDeploymentNodes, err := kkp.GetNodeDeploymentNodes(nodeDeployment.ID, cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrap(err, "Error fetching node deployment nodes")
		}

		nodeDeploymentEvents, err := kkp.GetNodeDeploymentEvents(nodeDeployment.ID, cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrap(err, "Error fetching node deployment events")
		}

		meta := &describe.NodeDeploymentDescribeMeta{
			NodeDeployment: &nodeDeployment,
			Nodes:          nodeDeploymentNodes,
			NodeEvents:     nodeDeploymentEvents,
		}

		parsed, err := describe.Object(meta)
		if err != nil {
			return errors.Wrap(err, "Error parsing datacenters")
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
