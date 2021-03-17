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

var ctxGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Lets you get the context of an specific type",
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
			return fmt.Errorf("cloud '%s' is not configured and therefore cannot be set as context", cloudName)
		}

		Config.Context.CloudName = cloudName
		return Config.Save()
	},
}

var ctxGetCloudCmd = &cobra.Command{
	Use:               "cloud",
	Short:             "Lets you get the context of which cloud to use",
	Args:              cobra.ExactArgs(0),
	Example:           "kkpctl ctx get cloud",
	ValidArgsFunction: getValidCloudContextArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Config.Context.CloudName)
	},
}

func init() {
	rootCmd.AddCommand(ctxCmd)
	ctxCmd.AddCommand(ctxSetCmd)
	ctxCmd.AddCommand(ctxGetCmd)

	// Set
	ctxSetCmd.AddCommand(ctxSetCloudCmd)

	// Get
	ctxGetCmd.AddCommand(ctxGetCloudCmd)
}
