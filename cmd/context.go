package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ctxCmd = &cobra.Command{
	Use:   "ctx",
	Short: "Lets you work with the current context in which kkpctl commands are executed",
}

var ctxSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Lets you set the context of an specific type",
}

var ctxSetCloudCmd = &cobra.Command{
	Use:               "cloud name",
	Short:             "Lets you set the context of which cloud to use",
	Args:              cobra.ExactArgs(1),
	Example:           "kkpctl ctx set cloud optimist",
	ValidArgsFunction: getValidCloudContextArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cloudName := args[0]

		_, ok := Config.Cloud[cloudName]
		if !ok {
			return fmt.Errorf("Cloud '%s' is not configured and therefore cannot be set as context", cloudName)
		}

		Config.Context.CloudName = cloudName
		return Config.Save(ConfigPath)
	},
}

var ctxSetBearerCmd = &cobra.Command{
	Use:     "bearer token",
	Short:   "Lets you set the context of which bearer token to use",
	Args:    cobra.ExactArgs(1),
	Example: "kkpctl ctx set bearer sdfhjsldkfjsdklfhj...",
	RunE: func(cmd *cobra.Command, args []string) error {
		Config.Context.Bearer = args[0]
		return Config.Save(ConfigPath)
	},
}

func init() {
	rootCmd.AddCommand(ctxCmd)
	ctxCmd.AddCommand(ctxSetCmd)
	ctxSetCmd.AddCommand(ctxSetCloudCmd)
	ctxSetCmd.AddCommand(ctxSetBearerCmd)
}
