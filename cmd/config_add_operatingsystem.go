package cmd

import (
	"github.com/cedi/kkpctl/pkg/config"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/spf13/cobra"
)

var (
	disableAutoUpdate bool
	distUpgradeOnBoot bool
)

var configAddOperatingSystemCmd = &cobra.Command{
	Use:       "operatingsystem {flatcar|ubuntu}",
	Short:     "Lets add a operation system spec",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"flatcar", "ubuntu"},
}

var configAddOSFlatcarCmd = &cobra.Command{
	Use:     "flatcar",
	Short:   "Lets add a flatcar operation system spec",
	Args:    cobra.ExactArgs(0),
	Example: "",
	RunE: func(cmd *cobra.Command, args []string) error {

		err := Config.OSSpec.AddOperatingSystemSpec(config.Flatcar, &models.FlatcarSpec{
			DisableAutoUpdate: disableAutoUpdate,
		})

		if err != nil {
			return err
		}

		return Config.Save()
	},
}

var configAddOSUbuntuCmd = &cobra.Command{
	Use:     "ubuntu",
	Short:   "Lets add a ubuntu operation system spec",
	Args:    cobra.ExactArgs(0),
	Example: "",
	RunE: func(cmd *cobra.Command, args []string) error {

		err := Config.OSSpec.AddOperatingSystemSpec(config.Ubuntu, &models.UbuntuSpec{
			DistUpgradeOnBoot: distUpgradeOnBoot,
		})

		if err != nil {
			return err
		}

		return Config.Save()
	},
}

func init() {
	configAddCmd.AddCommand(configAddOperatingSystemCmd)

	// Flatcar
	configAddOperatingSystemCmd.AddCommand(configAddOSFlatcarCmd)
	configAddOSFlatcarCmd.Flags().BoolVar(&disableAutoUpdate, "disable_auto_update", false, "Disable Auto Update")

	// Ubuntu
	configAddOperatingSystemCmd.AddCommand(configAddOSUbuntuCmd)
	configAddOSFlatcarCmd.Flags().BoolVar(&distUpgradeOnBoot, "dist_upgrade_on_boot", false, "Upgrade System on first boot")
}
