package cmd

import (
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/spf13/cobra"
)

var (
	osTenant         string
	osDomain         string
	osFloatingIPPool string
	osSecurityGroups string
	osNetwork        string
	osSubnetID       string
	username         string
	password         string
)

// configAddProviderCmd represents the add provider command
var configAddProviderCmd = &cobra.Command{
	Use:       "provider {openstack}",
	Short:     "Lets add a specific provider object",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"openstack"},
}

// adding an PpenStack provider
var configAddProviderOpenStackCmd = &cobra.Command{
	Use:     "openstack name",
	Short:   "Lets add a specific openstack provider",
	Args:    cobra.ExactArgs(1),
	Example: "kkpctl config add provider openstack optimist --username \"user@email.de\" --password \"my-super-secure-password\" --tenant \"internal-openstack-tenant\"",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := Config.Provider.AddProviderConfig(args[0], models.OpenstackCloudSpec{
			Username:       username,
			Password:       password,
			Domain:         osDomain,
			Tenant:         osTenant,
			FloatingIPPool: osFloatingIPPool,
			SecurityGroups: osSecurityGroups,
			Network:        osNetwork,
			SubnetID:       osSubnetID,
		})
		if err != nil {
			return err
		}

		return Config.Save()
	},
}

func init() {
	configAddCmd.AddCommand(configAddProviderCmd)

	// OpenStack
	configAddProviderCmd.AddCommand(configAddProviderOpenStackCmd)
	initOpenStackFlags()
}

func initOpenStackFlags() {
	configAddProviderOpenStackCmd.Flags().StringVar(&osTenant, "tenant", "", "The OpenStack tenant")
	configAddProviderOpenStackCmd.MarkFlagRequired("tenant")
	// TODO: configAddProviderOpenStackCmd.RegisterFlagCompletionFunc("tenant", ...)

	configAddProviderOpenStackCmd.Flags().StringVar(&osDomain, "domain", "default", "The OpenStack Domain")
	// TODO: configAddProviderOpenStackCmd.RegisterFlagCompletionFunc("domain", ...)

	configAddProviderOpenStackCmd.Flags().StringVar(&username, "username", "", "Your OpenStack Username")
	configAddProviderOpenStackCmd.MarkFlagRequired("username")

	configAddProviderOpenStackCmd.Flags().StringVar(&password, "password", "", "Your OpenStack Password")
	configAddProviderOpenStackCmd.MarkFlagRequired("password")

	configAddProviderOpenStackCmd.Flags().StringVar(&osFloatingIPPool, "floatingIpPool", "", "When specified, all worker nodes will receive a public ip from this floating ip pool")
	// TODO: configAddProviderOpenStackCmd.RegisterFlagCompletionFunc("floatingIpPool", ...)

	configAddProviderOpenStackCmd.Flags().StringVar(&osSecurityGroups, "securityGroup", "", "When specified, all worker nodes will be attached to this security group. If not specified, a security group will be created.")
	// TODO: configAddProviderOpenStackCmd.RegisterFlagCompletionFunc("securityGroup", ...)

	configAddProviderOpenStackCmd.Flags().StringVar(&osNetwork, "network", "", "When specified, all worker nodes will be attached to this network. If not specified, a network, subnet & router will be created")
	// TODO: configAddProviderOpenStackCmd.RegisterFlagCompletionFunc("network", ...)

	configAddProviderOpenStackCmd.Flags().StringVar(&osSubnetID, "subnet", "", "Please specify a SubnetID that exists in your network")
	// TODO: configAddProviderOpenStackCmd.RegisterFlagCompletionFunc("subnet", ...)
}
