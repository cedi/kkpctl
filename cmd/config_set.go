package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/config"
	"github.com/spf13/cobra"
)

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Lets you set the config of an specific type",
}

var configSetCloudCmd = &cobra.Command{
	Use:   "cloud",
	Short: "Lets you set the config of an specific type",
}

var configSetBearerCmd = &cobra.Command{
	Use:               "bearer token",
	Short:             "Lets you set the bearer token in the config",
	Args:              cobra.ExactArgs(2),
	Example:           "kkpctl config set cloud bearer imke_prod sdfhjsldkfjsdklfhj...",
	ValidArgsFunction: getValidCloudContextArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cloudName := args[0]
		bearer := args[1]

		cloud, ok := Config.Cloud[cloudName]
		if !ok {
			return fmt.Errorf("unable to find cloud with name '%s'", cloudName)
		}

		Config.Cloud[cloudName] = config.Cloud{
			URL:    cloud.URL,
			Bearer: bearer,
		}

		return Config.Save()
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
	configSetCmd.AddCommand(configSetCloudCmd)
	configSetCloudCmd.AddCommand(configSetBearerCmd)
}
