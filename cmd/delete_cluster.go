package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// delProjectsCmd represents the projects command
var delClusterCmd = &cobra.Command{
	Use:               "cluster clusterid",
	Short:             "Delete a cluster.",
	Example:           "kkpctl delete cluster rbw47nm2h8 --project dw2s9jk28z --datacenter ix1",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: getValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.New("Could not initialize Kubermatic API client")
		}

		cluster, err := kkp.GetCluster(args[0], projectID, datacenter)
		if err != nil {
			return errors.Wrap(err, "Error finding cluster")
		}

		err = kkp.DeleteCluster(args[0], projectID, datacenter)
		if err != nil {
			return errors.Wrap(err, "Error deleting cluster")
		}

		fmt.Printf("Successfully deleted Cluster %s (ClusterId: %s, ProjectId %s, Datacenter %s)",
			cluster.Name,
			args[0],
			projectID,
			datacenter,
		)

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(delClusterCmd)

	delClusterCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project to list clusters for.")
	delClusterCmd.MarkFlagRequired("project")

	delClusterCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter to list clusters for.")
	delClusterCmd.MarkFlagRequired("datacenter")
}
