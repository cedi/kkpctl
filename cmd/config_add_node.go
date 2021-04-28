package cmd

import (
	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	flavor                    string
	image                     string
	useFloatingIP             bool
	instanceReadyCheckPeriod  string
	instanceReadyCheckTimeout string
)

var configAddNodeSpecCmd = &cobra.Command{
	Use:   "node",
	Short: "Lets add a cloud node spec",
	Args:  cobra.ExactArgs(1),
}

var configAddOSNodeSpecCmd = &cobra.Command{
	Use:     "openstack name",
	Short:   "Lets add a cloud node spec for openstack",
	Args:    cobra.ExactArgs(1),
	Example: "kkpctl config add node openstack --flavor \"m1.micro\" --image \"Flatcar_Production 2020 - Latest\" flatcar-m1micro",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		err := Config.NodeSpec.AddCloudNodeSpec(name, &models.OpenstackNodeSpec{
			Flavor:                    &flavor,
			Image:                     &image,
			UseFloatingIP:             useFloatingIP,
			InstanceReadyCheckPeriod:  instanceReadyCheckPeriod,
			InstanceReadyCheckTimeout: instanceReadyCheckTimeout,
		})

		if err != nil {
			return errors.Wrapf(err, "failed to add openstack cloud %s", name)
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
	configAddOSNodeSpecCmd.RegisterFlagCompletionFunc("flavor", completion.GetValidFlavorArgs)

	configAddOSNodeSpecCmd.Flags().StringVar(&image, "image", "", "The OS image to use")
	configAddOSNodeSpecCmd.MarkFlagRequired("image")

	configAddOSNodeSpecCmd.Flags().BoolVar(&useFloatingIP, "floating_ip", false, "Allocate a floating IP")

	configAddOSNodeSpecCmd.Flags().StringVar(&instanceReadyCheckPeriod, "instance_ready_check_period", "10s", "interval in which the instance readyness is checked")
	configAddOSNodeSpecCmd.Flags().StringVar(&instanceReadyCheckTimeout, "instance_ready_check_timeout", "5m", "timeout how long to wait for a instance")
}
