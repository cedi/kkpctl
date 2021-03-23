package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/errors"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/spf13/cobra"
)

// clustersCmd represents the clusters command
var getDatacenterCmd = &cobra.Command{
	Use:               "datacenter [name]",
	Short:             "Lists available datacenters",
	Example:           "kkpctl get datacenter",
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: getValidDatacenterArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		dc := ""
		if len(args) == 1 {
			dc = args[1]
		}

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		var result interface{}
		if dc == "" {
			result, err = kkp.ListDatacenter()
		} else {
			result, err = kkp.GetDatacenter(args[0])
		}

		if err != nil {
			return errors.Wrap(err, "failed to get datacenter")
		}

		parsed, err := output.ParseOutput(result, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "failed to parse datacenter")
		}

		fmt.Print(parsed)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getDatacenterCmd)
}
