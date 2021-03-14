package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// clustersCmd represents the clusters command
var getDatacenterCmd = &cobra.Command{
	Use:               "datacenter [name]",
	Short:             "Lists all available datacenters",
	Example:           "kkpctl get datacenter",
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: getValidDatacenterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, apiToken := Config.GetCloudFromContext()
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.Wrap(err, "Could not initialize Kubermatic API client")
		}

		var result interface{}
		if len(args) == 0 {
			result, err = kkp.ListDatacenter()
		} else {
			result, err = kkp.GetDatacenter(args[0])
		}

		if err != nil {
			return errors.Wrap(err, "Error fetching datacenters")
		}

		parsed, err := output.ParseOutput(result, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "Error parsing datacenters")
		}

		fmt.Print(parsed)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getDatacenterCmd)
}
