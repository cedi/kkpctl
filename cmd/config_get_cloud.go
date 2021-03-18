package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var configGetCloudCmd = &cobra.Command{
	Use:     "cloud",
	Short:   "Get your configured clouds",
	Example: "kkpctl config get cloud",
	Args:    cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if Config.Cloud == nil || len(Config.Cloud) == 0 {
			fmt.Printf("no clouds configured")
			return nil
		}

		parsed, err := output.ParseOutput(Config.Cloud, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "failed to parse output")
		}

		fmt.Print(parsed)
		return nil
	},
}

func init() {
	configGetCmd.AddCommand(configGetCloudCmd)
}
