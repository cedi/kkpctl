package cmd

import (
	"fmt"
	"os"

	"github.com/cedi/kkpctl/pkg/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var write bool

var configGenerateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate a empty configuration",
	Example: "kkpctl config generate",
	Args:    cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if write {
			_, err := os.Stat(configPath)
			if err == nil {
				err = os.Remove(configPath)
				if err != nil {
					return errors.Wrap(err, "Unable to delete old configuration")
				}
			}

			return config.EnsureConfig()
		}

		config := config.NewConfig()

		yamlByte, err := yaml.Marshal(config)
		if err != nil {
			return errors.Wrap(err, "Unable to serialize empty configuration")
		}

		fmt.Println(string(yamlByte))

		return nil
	},
}

func init() {
	configCmd.AddCommand(configGenerateCmd)
	configGenerateCmd.Flags().BoolVarP(&write, "write", "w", false, "Write this configuratino to your in --config specified configuration file. Attention: This might overwrite your current configuration")
}
