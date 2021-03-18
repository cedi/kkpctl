package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/kubermatic/go-kubermatic/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	writeConfig bool
)

// clustersCmd represents the clusters command
var getKubeconfigCmd = &cobra.Command{
	Use:               "kubeconfig clusterid",
	Short:             "Gets the kubeconfig of a cluster",
	Example:           "kkpctl get kubeconfig --project x5zvx9bcx6 -w",
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: getValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		var cluster models.Cluster
		if datacenter == "" {
			cluster, err = kkp.GetClusterInProject(args[0], projectID)
		} else {
			cluster, err = kkp.GetClusterInProjectInDC(args[0], projectID, datacenter)
		}
		if err != nil {
			return err
		}

		result, err := kkp.GetKubeConfig(cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrap(err, "Error fetching kubeconfig")
		}

		if !writeConfig {
			fmt.Print(result)
		} else {
			err := ioutil.WriteFile(fmt.Sprintf("kubeconfig-admin-%s", args[0]), []byte(result), 0644)
			if err != nil {
				return errors.Wrap(err, "Error writing kubeconfig")
			}
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(getKubeconfigCmd)
	getKubeconfigCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project.")
	describeClusterCmd.MarkFlagRequired("project")
	getClustersCmd.RegisterFlagCompletionFunc("project", getValidProjectArgs)

	getKubeconfigCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter.")
	getClustersCmd.RegisterFlagCompletionFunc("datacenter", getValidDatacenterArgs)

	getKubeconfigCmd.Flags().BoolVarP(&writeConfig, "write", "w", false, "write the kubeconfig to the local directory")
}
