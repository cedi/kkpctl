package model

import (
	"github.com/kubermatic/go-kubermatic/models"
)

// NewCluster creates a new Cluster Object
func NewCluster(name string, datacenter string, version string, labels map[string]string) *models.Cluster {
	return &models.Cluster{
		Credential: "",
		Labels:     labels,
		Name:       name,
		Spec: &models.ClusterSpec{
			AdmissionPlugins:      []string{},
			EnableUserSSHKeyAgent: false,
			MachineNetworks: []*models.MachineNetworkingConfig{
				{
					CIDR:       "",
					DNSServers: []string{},
					Gateway:    "",
				},
			},
			PodNodeSelectorAdmissionPluginConfig: map[string]string{},
			UsePodNodeSelectorAdmissionPlugin:    false,
			UsePodSecurityPolicyAdmissionPlugin:  false,
			AuditLogging:                         &models.AuditLoggingSettings{},
			Cloud: &models.CloudSpec{
				DatacenterName: datacenter,
				Alibaba:        &models.AlibabaCloudSpec{},
				Anexia:         &models.AnexiaCloudSpec{},
				Aws:            &models.AWSCloudSpec{},
				Azure:          &models.AzureCloudSpec{},
				Bringyourown:   nil,
				Digitalocean:   &models.DigitaloceanCloudSpec{},
				Fake:           &models.FakeCloudSpec{},
				Gcp:            &models.GCPCloudSpec{},
				Hetzner:        &models.HetznerCloudSpec{},
				Kubevirt:       &models.KubevirtCloudSpec{},
				Openstack:      &models.OpenstackCloudSpec{},
				Packet:         &models.PacketCloudSpec{},
				Vsphere:        &models.VSphereCloudSpec{},
			},
			Version: version,
		},
	}
}

// NewCreateClusterSpec creates a new models.CreateClusterSpec object
func NewCreateClusterSpec(clusterName string, clusterType string, k8sVersion string, cloudSpec *models.CloudSpec, labels map[string]string, usePodNodeSelectorAdmissionPlugin bool, usePodSecurityPolicyAdmissionPlugin bool, enableAuditLogging bool) *models.CreateClusterSpec {
	return &models.CreateClusterSpec{
		Cluster: &models.Cluster{
			Labels: labels,
			Name:   clusterName,
			Type:   clusterType,
			Spec: &models.ClusterSpec{
				UsePodNodeSelectorAdmissionPlugin:   usePodNodeSelectorAdmissionPlugin,
				UsePodSecurityPolicyAdmissionPlugin: usePodSecurityPolicyAdmissionPlugin,
				AuditLogging: &models.AuditLoggingSettings{
					Enabled: enableAuditLogging,
				},
				Cloud:   cloudSpec,
				Version: k8sVersion,
			},
		},
	}
}
