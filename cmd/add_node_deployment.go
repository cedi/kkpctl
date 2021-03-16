package cmd

import (
	"fmt"
	"strings"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/kubermatic/go-kubermatic/models"
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
	Example: "kkpctl add nodedeployment --project 6tmbnhdl7h --cluster qvjdddt72t --nodespec flatcar1 --operatingsystem flatcar --provider optimist first_node_deployment",
	RunE: func(cmd *cobra.Command, args []string) error {
		clusterName := args[0]

		baseURL, apiToken := Config.GetCloudFromContext()
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.Wrap(err, "could not initialize Kubermatic API client")
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

		mapLabels := make(map[string]string)
		if labels != "" {
			slicedLabels := strings.Split(labels, ",")
			for _, slicedLabel := range slicedLabels {
				splitLabel := strings.Split(slicedLabel, "=")
				mapLabels[splitLabel[0]] = splitLabel[1]
			}
		}
		clusterVersion, ok := cluster.Spec.Version.(string)
		if !ok {
			return fmt.Errorf("cluster version does not appear to be a string")
		}

		newNodeDp := models.NodeDeployment{
			Name: args[0],
			Spec: &models.NodeDeploymentSpec{
				DynamicConfig: dynamicConfig,
				Replicas:      &nodeReplica,
				Template: &models.NodeSpec{
					Labels:          mapLabels,
					SSHUserName:     "",
					Taints:          []*models.TaintSpec{},
					Cloud:           Config.NodeSpec.GetNodeCloudSpec(nodeSpecName),
					OperatingSystem: Config.OSSpec.GetOperatingSystemSpec(),
					Versions: &models.NodeVersionInfo{
						Kubelet: clusterVersion,
					},
				},
			},
		}

		_, err = kkp.CreateNodeDeployment(&newNodeDp, clusterID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrap(err, "Error fetching projects")
		}

		fmt.Printf("Successfully created cluster '%s'\n", clusterName)
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
