package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// clustersCmd represents the clusters command
var getClusterVersionCmd = &cobra.Command{
	Use:               "version",
	Short:             "Lists all available Kubernetes Versions in your KKP installation",
	Example:           "kkpctl get version",
	Args:              cobra.ExactArgs(0),
	ValidArgsFunction: completion.GetValidDatacenterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		kkp, err := Config.GetKKPClient(client.V1API)
		if err != nil {
			return err
		}

		result, err := kkp.ListClusterVersions()

		if err != nil {
			return errors.Wrap(err, "failed to get versions")
		}

		parsed, err := output.ParseOutput(result, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "failed parsing versions")
		}

		fmt.Print(parsed)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getClusterVersionCmd)
}
