package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ctxCmd = &cobra.Command{
	Use:   "ctx",
	Short: "Manipulate the current context kkpctl works with",
	Long:  "kkpctl uses this context to save the cloud which you want to connect to",
}

var ctxSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set context",
}

var ctxGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get context",
}

var ctxSetCloudCmd = &cobra.Command{
	Use:               "cloud name",
	Short:             "Set the cloud which you want to connect to",
	Args:              cobra.ExactArgs(1),
	Example:           "kkpctl ctx set cloud imke",
	ValidArgsFunction: getValidCloudContextArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cloudName := args[0]

		_, err := Config.Cloud.Get(cloudName)
		if err != nil {
			return fmt.Errorf("failed to set the cloud %s to the current context", cloudName)
		}

		Config.Context.CloudName = cloudName
		return Config.Save()
	},
}

var ctxGetCloudCmd = &cobra.Command{
	Use:               "cloud",
	Short:             "Get the cloud which you want to connect to",
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
