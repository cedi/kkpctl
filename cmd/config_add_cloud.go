package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/config"
	"github.com/spf13/cobra"
)

var (
	cloudURL     string
	clientID     string
	clientSecret string
	authToken    string
)

// configAddProviderCmd represents the add provider command
var configAddCloudCmd = &cobra.Command{
	Use:     "cloud name",
	Short:   "Lets add a specific cloud object",
	Example: "kkpctl config add cloud imke --url https://imke.cloud --client_id kubermatic --client_secret akdfjhklqwerhli2uh=",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if clientID == "" && clientSecret != "" {
			return fmt.Errorf("--client_secret was set, but no --client_id")
		} else if clientID != "" && clientSecret == "" {
			return fmt.Errorf("--client_id was set, but no --client_secret")
		} else if clientID == "" && clientSecret == "" && authToken == "" {
			return fmt.Errorf("OIDC login is not configured (no --client_id and --client_secret), but also no static --auth_token was specified")
		}

		if Config.Cloud == nil {
			Config.Cloud = config.NewCloudConfig()
		}

		Config.Cloud.Set(name, config.NewCloud(cloudURL, clientID, clientSecret, authToken))

		return Config.Save()
	},
}

func init() {
	configAddCmd.AddCommand(configAddCloudCmd)
	configAddCloudCmd.Flags().StringVar(&cloudURL, "url", "", "The URL to your KKP installation")
	configAddCloudCmd.MarkFlagRequired("url")

	configAddCloudCmd.Flags().StringVar(&clientID, "client_id", "kubermatic", "The ClientID to use for OIDC-Login")
	configAddCloudCmd.Flags().StringVar(&clientSecret, "client_secret", "", "The ClientSecret to use for OIDC-Login")
	configAddCloudCmd.Flags().StringVar(&authToken, "auth_token", "", "If you're not using OIDC Login, you can specify a static API Token here")
}
