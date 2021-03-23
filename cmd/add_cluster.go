package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/model"
	"github.com/cedi/kkpctl/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	clusterType                         string
	k8sVersion                          string
	providerName                        string
	enableAuditLogging                  bool
	usePodSecurityPolicyAdmissionPlugin bool
	usePodNodeSelectorAdmissionPlugin   bool
)

var createClusterCmd = &cobra.Command{
	Use:     "cluster clusterName",
	Short:   "Create a new cluster",
	Example: "kkpctl add cluster --project 6tmbnhdl7h --datacenter ix2 --provider optimist --version 1.18.13 --labels stage=dev kkpctltest",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		clusterName := args[0]

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		newCluster := model.NewCreateClusterSpec(
			clusterName,
			clusterType,
			k8sVersion,
			Config.Provider.GetProviderCloudSpec(providerName, datacenter),
			utils.SplitLabelString(labels),
			usePodNodeSelectorAdmissionPlugin,
			usePodSecurityPolicyAdmissionPlugin,
			enableAuditLogging,
		)

		result, err := kkp.CreateCluster(newCluster, projectID, datacenter)
		if err != nil {
			return errors.Wrapf(err, "failed to create %s cluster %s", clusterType, clusterName)
		}

		fmt.Printf("Successfully created cluster '%s' (%s)\n", clusterName, result.ID)
		return nil
	},
}

func init() {
	addCmd.AddCommand(createClusterCmd)

	createClusterCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project.")
	createClusterCmd.MarkFlagRequired("project")
	createClusterCmd.RegisterFlagCompletionFunc("project", getValidProjectArgs)

	createClusterCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter.")
	createClusterCmd.MarkFlagRequired("datacenter")
	createClusterCmd.RegisterFlagCompletionFunc("datacenter", getValidDatacenterArgs)

	createClusterCmd.Flags().StringVar(&clusterType, "type", "kubernetes", "Type of the cluster (kubernetes or openshift)")
	createClusterCmd.RegisterFlagCompletionFunc("type", getValidClusterTypes)

	createClusterCmd.Flags().StringVarP(&k8sVersion, "version", "v", "", "Name of the datacenter")
	createClusterCmd.MarkFlagRequired("version")
	createClusterCmd.RegisterFlagCompletionFunc("version", getValidKubernetesVersions)

	createClusterCmd.Flags().StringVar(&providerName, "provider", "", "Which provider should be used")
	createClusterCmd.RegisterFlagCompletionFunc("provider", getValidProvider)
	createClusterCmd.MarkFlagRequired("provider")

	createClusterCmd.Flags().BoolVar(&enableAuditLogging, "audit-logging", false, "Enable audit logging")
	createClusterCmd.Flags().BoolVar(&usePodSecurityPolicyAdmissionPlugin, "pod-security-policy", false, "Pod Security Policies allow detailed authorizatin of pod creation and updates")
	createClusterCmd.Flags().BoolVar(&usePodNodeSelectorAdmissionPlugin, "pod-node-selector", false, "Use the Pod Node Selector Admission Plugin")

	createClusterCmd.Flags().StringVarP(&labels, "labels", "l", "", "A comma separated list of labels in the format key=value")
}
