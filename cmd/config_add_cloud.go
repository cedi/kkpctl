package cmd

import (
	"github.com/cedi/kkpctl/pkg/config"
	"github.com/spf13/cobra"
)

var cloudURL string

// configAddProviderCmd represents the add provider command
var configAddCloudCmd = &cobra.Command{
	Use:     "cloud name",
	Short:   "Lets add a specific cloud object",
	Example: "kkpctl config add cloud imke --url https://imke.cloud",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if Config.Cloud == nil {
			Config.Cloud = config.NewCloudConfig()
		}

		Config.Cloud.Set(name, config.Cloud{URL: cloudURL})

		return Config.Save()
	},
}

func init() {
	configAddCmd.AddCommand(configAddCloudCmd)
	configAddCloudCmd.Flags().StringVar(&cloudURL, "url", "", "The URL to your KKP installation")
	configAddCloudCmd.MarkFlagRequired("url")
}
