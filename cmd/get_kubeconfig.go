package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/cedi/kkpctl/pkg/client"
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
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completion.GetValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		clusterID := args[0]

		kkp, err := Config.GetKKPClient(client.V1API)
		if err != nil {
			return err
		}

		cluster, err := kkp.GetClusterInProjectInDC(clusterID, projectID, datacenter)
		if err != nil {
			return errors.Wrapf(err, "failed to get cluster %s in project %s", clusterID, projectID)
		}

		result, err := kkp.GetKubeConfig(cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrapf(err, "failed to get kubeconfig for cluster %s", clusterID)
		}

		if !writeConfig {
			fmt.Print(result)
		} else {
			err := ioutil.WriteFile(fmt.Sprintf("kubeconfig-admin-%s", args[0]), []byte(result), 0644)
			if err != nil {
				return errors.Wrapf(err, "failed to write kubeconfig to current location")
			}
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(getKubeconfigCmd)

	AddProjectFlag(getKubeconfigCmd)
	AddDatacenterFlag(getKubeconfigCmd, false)

	getKubeconfigCmd.Flags().BoolVarP(&writeConfig, "write", "w", false, "write the kubeconfig to the local directory")
}
