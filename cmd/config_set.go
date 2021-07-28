package cmd

import (
	"github.com/cedi/kkpctl/pkg/config"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	disableAutoUpdate bool
	distUpgradeOnBoot bool
)

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set values in your configuration",
}

var configSetOperatingSystemCmd = &cobra.Command{
	Use:       "operatingsystem {flatcar|ubuntu}",
	Short:     "Configure operation system behaviour",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"flatcar", "ubuntu"},
}

var configSetOSFlatcarCmd = &cobra.Command{
	Use:     "flatcar [--disable-auto-update]",
	Short:   "Configure Flatcar system behaviour",
	Args:    cobra.ExactArgs(0),
	Example: "kkpctl config set operatingsystem flatcar --disable-auto-update",
	RunE: func(cmd *cobra.Command, args []string) error {

		err := Config.OSSpec.SetOperatingSystemSpec(config.Flatcar, &models.FlatcarSpec{
			DisableAutoUpdate: disableAutoUpdate,
		})

		if err != nil {
			return errors.Wrapf(err, "failed to set flatcar configuration")
		}

		return Config.Save()
	},
}

var configSetOSUbuntuCmd = &cobra.Command{
	Use:     "ubuntu",
	Short:   "Configure Ubuntu system behaviour",
	Args:    cobra.ExactArgs(0),
	Example: "kkpctl config set operatingsystem ubuntu --distupgrade-on-boot",
	RunE: func(cmd *cobra.Command, args []string) error {

		err := Config.OSSpec.SetOperatingSystemSpec(config.Ubuntu, &models.UbuntuSpec{
			DistUpgradeOnBoot: distUpgradeOnBoot,
		})

		if err != nil {
			return errors.Wrapf(err, "failed to set ubuntu configuration")
		}

		return Config.Save()
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)

	configSetCmd.AddCommand(configSetOperatingSystemCmd)

	// Flatcar
	configSetOperatingSystemCmd.AddCommand(configSetOSFlatcarCmd)
	configSetOSFlatcarCmd.Flags().BoolVar(&disableAutoUpdate, "disable-auto-update", false, "Disable Auto Update")

	// Ubuntu
	configSetOperatingSystemCmd.AddCommand(configSetOSUbuntuCmd)
	configSetOSUbuntuCmd.Flags().BoolVar(&distUpgradeOnBoot, "distupgrade-on-boot", false, "Upgrade System on first boot")
}
