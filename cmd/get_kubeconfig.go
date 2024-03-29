package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/cedi/kkpctl/pkg/utils"
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

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		cluster, err := kkp.GetCluster(clusterID, projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to get cluster %s in project %s", clusterID, projectID)
		}

		result, err := kkp.GetKubeConfig(cluster.ID, projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to get kubeconfig for cluster %s", clusterID)
		}

		if !writeConfig {
			fmt.Print(result)
			return nil
		}

		fileName := fmt.Sprintf("kubeconfig-admin-%s", clusterID)

		currentDir, err := os.Getwd()
		if err != nil {
			return errors.Wrapf(err, "failed to retrieve current location")
		}

		filePath := fmt.Sprintf("%s/%s", currentDir, fileName)

		err = ioutil.WriteFile(filePath, []byte(result), 0644)
		if err != nil {
			return errors.Wrapf(err, "failed to write kubeconfig to current location")
		}

		size, err := utils.GetFileSize(filePath)
		if err != nil {
			return errors.Wrapf(err, "failed to retrieve filesize")
		}

		fmt.Printf("kubeconfig written: %s (%d bytes)\n", filePath, size)
		fmt.Printf("\nConfigure kubectl using\nexport KUBECONFIG=%s\n", filePath)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getKubeconfigCmd)

	AddProjectFlag(getKubeconfigCmd)

	getKubeconfigCmd.Flags().BoolVarP(&writeConfig, "write", "w", false, "write the kubeconfig to the local directory")
}
