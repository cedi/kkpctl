package cmd

import (
	"fmt"
	"strings"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/kubermatic/go-kubermatic/models"
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

// projectCmd represents the project command
var createClusterCmd = &cobra.Command{
	Use:   "cluster name",
	Short: "Lets you create a new cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		clusterName := args[0]

		baseURL, apiToken := Config.GetCloudFromContext()
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.Wrap(err, "Could not initialize Kubermatic API client")
		}

		mapLabels := make(map[string]string)
		if labels != "" {
			slicedLabels := strings.Split(labels, ",")
			for _, slicedLabel := range slicedLabels {
				splitLabel := strings.Split(slicedLabel, "=")
				mapLabels[splitLabel[0]] = splitLabel[1]
			}
		}

		newCluster := models.CreateClusterSpec{
			Cluster: &models.Cluster{
				Labels: mapLabels,
				Name:   clusterName,
				Type:   clusterType,
				Spec: &models.ClusterSpec{
					UsePodNodeSelectorAdmissionPlugin:   usePodNodeSelectorAdmissionPlugin,
					UsePodSecurityPolicyAdmissionPlugin: usePodSecurityPolicyAdmissionPlugin,
					AuditLogging: &models.AuditLoggingSettings{
						Enabled: enableAuditLogging,
					},
					Cloud:   Config.Provider.GetProviderCloudSpec(providerName, datacenter),
					Version: k8sVersion,
				},
			},
		}

		_, err = kkp.CreateCluster(&newCluster, projectID, datacenter)
		if err != nil {
			return errors.Wrap(err, "Error fetching projects")
		}

		fmt.Printf("Successfully created cluster '%s'\n", clusterName)
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
