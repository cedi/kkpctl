package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bwplotka/oidc/login"
	disk "github.com/bwplotka/oidc/login/diskcache"
	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/cedi/kkpctl/pkg/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var oidcLoginCmd = &cobra.Command{
	Use:               "oidc-login [cloudname]",
	Short:             "Login to your KKP Cloud",
	Args:              cobra.MaximumNArgs(1),
	Example:           "kkpctl oidc-login imke",
	ValidArgsFunction: completion.GetValidCloudContextArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var kkpCloud *config.Cloud
		var err error

		if len(args) == 1 {
			kkpCloud, err = Config.Cloud.Get(args[0])
		} else {
			kkpCloud, err = Config.Cloud.Get(Config.Context.CloudName)
		}

		if err != nil {
			return err
		}

		providerURL := fmt.Sprintf("%s/dex", kkpCloud.URL)
		oidcConfig := login.OIDCConfig{
			ClientID:     kkpCloud.OIDC.ClientID,
			ClientSecret: kkpCloud.OIDC.ClientSecret,
			Provider:     providerURL,
			Scopes:       []string{"openid"},
		}

		sourceConfig := login.Config{
			NonceCheck: true,
		}

		cache := disk.NewCache(".super_cache", oidcConfig)

		source, _, err := login.NewOIDCTokenSource(context.Background(), log.New(os.Stdout, "", 0), sourceConfig, cache, nil)
		if err != nil {
			return errors.Wrap(err, "failed to obtain OIDC token source")
		}

		token, err := source.OIDCToken(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed to obtain OIDC token")
		}

		kkpCloud.Bearer = token.AccessToken
		return Config.Save()
	},
}

func init() {
	rootCmd.AddCommand(oidcLoginCmd)
}
