package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	projectID  string
	datacenter string
)

// clustersCmd represents the clusters command
var getClustersCmd = &cobra.Command{
	Use:               "cluster [clusterid]",
	Short:             "Lists clusters for a given project or datacenter",
	Example:           "kkpctl get cluster --project dw2s9jk28z",
	Args:              cobra.MaximumNArgs(1),
	Aliases:           []string{"clusters"},
	ValidArgsFunction: completion.GetValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		clusterID := ""
		if len(args) == 1 {
			clusterID = args[0]
		}

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		var result interface{}
		if clusterID == "" && projectID == "" {
			result, err = kkp.ListAllClusters(listAll)
		} else if clusterID == "" && projectID != "" {
			result, err = kkp.ListClusters(projectID)
		} else if clusterID != "" && projectID == "" {
			result, err = kkp.GetClusterByID(clusterID)
		} else {
			result, err = kkp.GetCluster(clusterID, projectID)
		}

		if err != nil {
			return errors.Wrap(err, "unable to get cluster")
		}

		parsed, err := output.ParseOutput(result, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "failed to parse cluster")
		}
		fmt.Print(parsed)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getClustersCmd)

	getClustersCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project.")
	getClustersCmd.RegisterFlagCompletionFunc("project", completion.GetValidProjectArgs)
}
