package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/output"
	"github.com/kubermatic/go-kubermatic/models"
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
	ValidArgsFunction: getValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		var cluster models.Cluster
		if datacenter == "" && projectID != "" {
			cluster, err = kkp.GetClusterInProject(args[0], projectID)
		} else {
			cluster, err = kkp.GetClusterInProjectInDC(args[0], projectID, datacenter)
		}

		if err != nil {
			return errors.Wrap(err, "failed to retrieve cluster")
		}

		result, err := kkp.UpgradeCluster(toVersion, cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
		if err != nil {
			return errors.Wrap(err, "failed to upgrade cluster")
		}

		if !noUpgradeWorkers {
			err = kkp.UpgradeWorkerDeploymentVersion(toVersion, cluster.ID, projectID, cluster.Spec.Cloud.DatacenterName)
			if err != nil {
				return errors.Wrap(err, "failed to upgrade worker nodes")
			}
		}

		parsed, err := output.ParseOutput(result, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "Error parsing clusters")
		}
		fmt.Print(parsed)

		return nil
	},
}

func init() {
	editClusterCmd.AddCommand(editClusterUpgradeCmd)

	editClusterUpgradeCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project to list clusters for.")
	editClusterUpgradeCmd.MarkFlagRequired("project")
	editClusterUpgradeCmd.RegisterFlagCompletionFunc("project", getValidProjectArgs)

	editClusterUpgradeCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter to list clusters for.")
	editClusterUpgradeCmd.RegisterFlagCompletionFunc("datacenter", getValidDatacenterArgs)

	editClusterUpgradeCmd.Flags().StringVar(&toVersion, "to-version", "", "To which Version should the cluster be updated")
	editClusterUpgradeCmd.MarkFlagRequired("to-version")
	editClusterUpgradeCmd.RegisterFlagCompletionFunc("to-version", getValidToVersionArgs)

	editClusterUpgradeCmd.Flags().BoolVar(&noUpgradeWorkers, "no-upgrade-workers", false, "Do not automatically upgrade the workers kubelet version")
}
