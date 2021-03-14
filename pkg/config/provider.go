package config

import "github.com/kubermatic/go-kubermatic/models"

// ProviderConfig contains the various possible providers in KKP
type ProviderConfig struct {
	Alibaba      map[string]models.AlibabaCloudSpec      `yaml:"alibaba,omitempty"`
	Anexia       map[string]models.AnexiaCloudSpec       `yaml:"anexia,omitempty"`
	Aws          map[string]models.AWSCloudSpec          `yaml:"aws,omitempty"`
	Azure        map[string]models.AzureCloudSpec        `yaml:"azure,omitempty"`
	Bringyourown map[string]models.BringYourOwnCloudSpec `yaml:"bringyourown,omitempty"`
	Digitalocean map[string]models.DigitaloceanCloudSpec `yaml:"digitalocean,omitempty"`
	Fake         map[string]models.FakeCloudSpec         `yaml:"fake,omitempty"`
	Gcp          map[string]models.GCPCloudSpec          `yaml:"gcp,omitempty"`
	Hetzner      map[string]models.HetznerCloudSpec      `yaml:"hetzner,omitempty"`
	Kubevirt     map[string]models.KubevirtCloudSpec     `yaml:"kubevirt,omitempty"`
	Openstack    map[string]models.OpenstackCloudSpec    `yaml:"openstack,omitempty"`
	Packet       map[string]models.PacketCloudSpec       `yaml:"packet,omitempty"`
	Vsphere      map[string]models.VSphereCloudSpec      `yaml:"vsphere,omitempty"`
}

// NewProvider creates a new, empty, provider config
func NewProvider() ProviderConfig {
	return ProviderConfig{
		Alibaba:      make(map[string]models.AlibabaCloudSpec),
		Anexia:       make(map[string]models.AnexiaCloudSpec),
		Aws:          make(map[string]models.AWSCloudSpec),
		Azure:        make(map[string]models.AzureCloudSpec),
		Bringyourown: make(map[string]models.BringYourOwnCloudSpec),
		Digitalocean: make(map[string]models.DigitaloceanCloudSpec),
		Fake:         make(map[string]models.FakeCloudSpec),
		Gcp:          make(map[string]models.GCPCloudSpec),
		Hetzner:      make(map[string]models.HetznerCloudSpec),
		Kubevirt:     make(map[string]models.KubevirtCloudSpec),
		Openstack:    make(map[string]models.OpenstackCloudSpec),
		Packet:       make(map[string]models.PacketCloudSpec),
		Vsphere:      make(map[string]models.VSphereCloudSpec),
	}
}
