package cmd

import (
	"github.com/spf13/cobra"
)

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Lets get a specific configuration object",
}

func init() {
	configCmd.AddCommand(configGetCmd)
}
