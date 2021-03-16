package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// clustersCmd represents the clusters command
var getClusterVersionCmd = &cobra.Command{
	Use:               "version",
	Short:             "Lists all available Kubernetes Versions",
	Example:           "kkpctl get version",
	Args:              cobra.ExactArgs(0),
	ValidArgsFunction: getValidDatacenterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, apiToken := Config.GetCloudFromContext()
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.Wrap(err, "Could not initialize Kubermatic API client")
		}

		result, err := kkp.ListClusterVersions()

		if err != nil {
			return errors.Wrap(err, "Error fetching versions")
		}

		parsed, err := output.ParseOutput(result, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "Error parsing versions")
		}

		fmt.Print(parsed)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getClusterVersionCmd)
}
