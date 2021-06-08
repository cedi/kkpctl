package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	toVersion        string
	noUpgradeWorkers bool
)

var editClusterUpgradeCmd = &cobra.Command{
	Use:               "upgrade clusterid",
	Short:             "Upgrades a clusters",
	Example:           "kkpctl edit cluster upgrade --project dw2s9jk28z x5zvx9bcx6 --to-version 1.18.13",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completion.GetValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		clusterID := args[0]

		kkp, err := Config.GetKKPClient(client.V1API)
		if err != nil {
			return err
		}

		kkp2, err := Config.GetKKPClient(client.V1API)
		if err != nil {
			return err
		}

		cluster, err := kkp.GetClusterInProjectInDC(clusterID, projectID, datacenter)
		if err != nil {
			return errors.Wrapf(err, "failed to get cluster %s in project %s", clusterID, projectID)
		}

		result, err := kkp.UpgradeCluster(toVersion, cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrapf(err, "failed to upgrade cluster %s to version %s", clusterID, toVersion)
		}

		if !noUpgradeWorkers {
			err = kkp2.UpgradeWorkerDeploymentVersion(toVersion, cluster.ID, projectID)
			if err != nil {
				return errors.Wrapf(err, "failed to upgrade worker-nodes in cluster %s to version %s", clusterID, toVersion)
			}
		}

		parsed, err := output.ParseOutput(result, outputType, sortBy)
		if err != nil {
			return errors.Wrapf(err, "failed to parse cluster %s", clusterID)
		}
		fmt.Print(parsed)

		return nil
	},
}

func init() {
	editClusterCmd.AddCommand(editClusterUpgradeCmd)

	AddProjectFlag(editClusterUpgradeCmd)
	AddDatacenterFlag(editClusterUpgradeCmd, false)

	editClusterUpgradeCmd.Flags().StringVar(&toVersion, "to-version", "", "To which Version should the cluster be updated")
	editClusterUpgradeCmd.MarkFlagRequired("to-version")
	editClusterUpgradeCmd.RegisterFlagCompletionFunc("to-version", completion.GetValidToVersionArgs)

	editClusterUpgradeCmd.Flags().BoolVar(&noUpgradeWorkers, "no-upgrade-workers", false, "Do not automatically upgrade the workers kubelet version")
}
