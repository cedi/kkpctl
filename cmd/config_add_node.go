package cmd

import (
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/spf13/cobra"
)

var (
	flavor        string
	image         string
	useFloatingIP bool
)

var configAddNodeSpecCmd = &cobra.Command{
	Use:       "node",
	Short:     "Lets add a cloud node spec",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"openstack"},
}

var configAddOSNodeSpecCmd = &cobra.Command{
	Use:     "openstack",
	Short:   "Lets add a cloud node spec for openstack",
	Args:    cobra.ExactArgs(1),
	Example: "",
	RunE: func(cmd *cobra.Command, args []string) error {

		err := Config.NodeSpec.AddCloudNodeSpec(args[0], models.OpenstackNodeSpec{
			Flavor:        &flavor,
			Image:         &image,
			UseFloatingIP: useFloatingIP,
		})

		if err != nil {
			return err
		}

		return Config.Save()
	},
}

func init() {
	configAddCmd.AddCommand(configAddNodeSpecCmd)

	// OpenStack
	configAddNodeSpecCmd.AddCommand(configAddOSNodeSpecCmd)

	configAddOSNodeSpecCmd.Flags().StringVar(&flavor, "flavor", "", "The OS Flavor to use")
	configAddOSNodeSpecCmd.MarkFlagRequired("flavor")
	configAddOSNodeSpecCmd.RegisterFlagCompletionFunc("flavor", getValidFlavorArgs)

	configAddOSNodeSpecCmd.Flags().StringVar(&image, "image", "", "The OS image to use")
	configAddOSNodeSpecCmd.MarkFlagRequired("image")

	configAddOSNodeSpecCmd.Flags().BoolVar(&useFloatingIP, "floating_ip", false, "Allocate a floating IP")
	configAddOSNodeSpecCmd.MarkFlagRequired("image")
}
