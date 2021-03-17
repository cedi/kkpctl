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
		if Config.Cloud == nil {
			Config.Cloud = make(map[string]config.Cloud)
		}
		Config.Cloud[args[0]] = config.Cloud{
			URL: cloudURL,
		}

		return Config.Save()
	},
}

func init() {
	configAddCmd.AddCommand(configAddCloudCmd)
	configAddCloudCmd.Flags().StringVar(&cloudURL, "url", "", "The KKP Cloud URL to connect to")
	configAddCloudCmd.MarkFlagRequired("url")
}
