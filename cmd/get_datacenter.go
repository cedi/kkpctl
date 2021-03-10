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
	Short:             "Lists clusters for a given project (and optional seed datacenter) or fetch a named cluster.",
	Long:              `If no clusterid is specified, all clusters of an project are listed. If a clusterid is specified only this cluster is shown`,
	Example:           "kkpctl get datacenter ix1",
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: getValidClusterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
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

		fmt.Println(parsed)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getDatacenterCmd)
}
