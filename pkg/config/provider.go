package config

import "github.com/kubermatic/go-kubermatic/models"

// ProviderConfig contains the various possible providers in KKP
type ProviderConfig struct {
	// DatacenterName where the users 'cloud' lives in.
	DatacenterName string `json:"dc,omitempty"`

	// alibaba
	Alibaba *models.AlibabaCloudSpec `json:"alibaba,omitempty"`

	// anexia
	Anexia *models.AnexiaCloudSpec `json:"anexia,omitempty"`

	// aws
	Aws *models.AWSCloudSpec `json:"aws,omitempty"`

	// azure
	Azure *models.AzureCloudSpec `json:"azure,omitempty"`

	// bringyourown
	Bringyourown models.BringYourOwnCloudSpec `json:"bringyourown,omitempty"`

	// digitalocean
	Digitalocean *models.DigitaloceanCloudSpec `json:"digitalocean,omitempty"`

	// fake
	Fake *models.FakeCloudSpec `json:"fake,omitempty"`

	// gcp
	Gcp *models.GCPCloudSpec `json:"gcp,omitempty"`

	// hetzner
	Hetzner *models.HetznerCloudSpec `json:"hetzner,omitempty"`

	// kubevirt
	Kubevirt *models.KubevirtCloudSpec `json:"kubevirt,omitempty"`

	// openstack
	Openstack *models.OpenstackCloudSpec `json:"openstack,omitempty"`

	// packet
	Packet *models.PacketCloudSpec `json:"packet,omitempty"`

	// vsphere
	Vsphere *models.VSphereCloudSpec `json:"vsphere,omitempty"`
}
