package cmd

import (
	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/cedi/kkpctl/pkg/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set values in your configuration",
}

var configSetCloudCmd = &cobra.Command{
	Use:   "cloud",
	Short: "Set values in your cloud configuration",
}

var configSetBearerCmd = &cobra.Command{
	Use:               "bearer cloudName bearerToken",
	Short:             "Set the bearer token for a cloud",
	Args:              cobra.ExactArgs(2),
	Example:           "kkpctl config set cloud bearer imke_prod sdfhjsldkfjsdklfhj...",
	ValidArgsFunction: completion.GetValidCloudContextArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cloudName := args[0]
		bearer := args[1]

		cloud, err := Config.Cloud.Get(cloudName)
		if err != nil {
			return errors.Wrapf(err, "failed to set berar token for cloud %s", cloudName)
		}

		Config.Cloud.Set(cloudName, config.NewCloud(cloud.URL, bearer))

		return Config.Save()
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
	configSetCmd.AddCommand(configSetCloudCmd)
	configSetCloudCmd.AddCommand(configSetBearerCmd)
}
