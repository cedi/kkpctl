package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	noDeleteVolumes       bool
	noDeleteLoadBalancers bool
)

// delProjectsCmd represents the projects command
var delClusterCmd = &cobra.Command{
	Use:               "cluster clusterid",
	Short:             "Delete a cluster",
	Example:           "kkpctl delete cluster rbw47nm2h8 --project dw2s9jk28z",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: getValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		clusterID := args[0]

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		cluster, err := kkp.GetClusterInProject(clusterID, projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to find cluster %s which should be deleted", clusterID)
		}

		err = kkp.DeleteClusterInDC(clusterID, projectID, datacenter, !noDeleteVolumes, !noDeleteLoadBalancers)
		if err != nil {
			return errors.Wrapf(err, "failed to delete cluster %s in project %s", clusterID, projectID)
		}

		fmt.Printf("Successfully deleted Cluster %s (%s) in ProjectId %s in Datacenter %s\n",
			cluster.Name,
			cluster.ID,
			projectID,
			cluster.Spec.Cloud.DatacenterName,
		)

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(delClusterCmd)

	delClusterCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project to list clusters for")
	delClusterCmd.MarkFlagRequired("project")
	delClusterCmd.RegisterFlagCompletionFunc("project", getValidProjectArgs)

	delClusterCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter to delete the cluster in")
	delClusterCmd.RegisterFlagCompletionFunc("datacenter", getValidDatacenterArgs)

	delClusterCmd.Flags().BoolVar(&noDeleteVolumes, "no-delete-volumes", false, "Do not cleanup connected volumes (PVs and PCVs)")
	delClusterCmd.Flags().BoolVar(&noDeleteLoadBalancers, "no-delete-loadbalancers", false, "Do not cleanup connected Load Balancers")
}
