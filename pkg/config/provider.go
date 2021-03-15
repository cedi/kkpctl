package config

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/utils"
	"github.com/kubermatic/go-kubermatic/models"
)

// CloudSpecType is used to identify which cloud to use
type CloudSpecType int

const (
	// Alibaba is the identifier for a Alibaba clous spec
	Alibaba CloudSpecType = iota
	// Anexia is the identifier for a Anexia clous spec
	Anexia
	// Aws is the identifier for a Aws clous spec
	Aws
	// Azure is the identifier for a Azure clous spec
	Azure
	// Bringyourown is the identifier for a Bringyourown clous spec
	Bringyourown
	// Digitalocean is the identifier for a Digitalocean clous spec
	Digitalocean
	// Fake is the identifier for a Fake clous spec
	Fake
	// Gcp is the identifier for a Gcp clous spec
	Gcp
	// Hetzner is the identifier for a Hetzner clous spec
	Hetzner
	// Kubevirt is the identifier for a Kubevirt clous spec
	Kubevirt
	// Openstack is the identifier for a Openstack clous spec
	Openstack
	// Packet is the identifier for a Packet clous spec
	Packet
	// Vsphere is the identifier for a Vsphere clous spec
	Vsphere
)

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

// GetProviderCloudSpec returns a models.CloudSpec object with the correct CloudProvider filled
func (p *ProviderConfig) GetProviderCloudSpec(providerName string, datacenter string) *models.CloudSpec {
	cs := &models.CloudSpec{
		DatacenterName: datacenter,
	}

	providerAlibaba, ok := p.Alibaba[providerName]
	if ok {
		cs.Alibaba = &providerAlibaba
	}
	providerAnexia, ok := p.Anexia[providerName]
	if ok {
		cs.Anexia = &providerAnexia
	}
	providerAws, ok := p.Aws[providerName]
	if ok {
		cs.Aws = &providerAws
	}
	providerAzure, ok := p.Azure[providerName]
	if ok {
		cs.Azure = &providerAzure
	}
	providerBringyourown, ok := p.Bringyourown[providerName]
	if ok {
		cs.Bringyourown = &providerBringyourown
	}
	providerDigitalocean, ok := p.Digitalocean[providerName]
	if ok {
		cs.Digitalocean = &providerDigitalocean
	}
	providerFake, ok := p.Fake[providerName]
	if ok {
		cs.Fake = &providerFake
	}
	providerGcp, ok := p.Gcp[providerName]
	if ok {
		cs.Gcp = &providerGcp
	}
	providerHetzner, ok := p.Hetzner[providerName]
	if ok {
		cs.Hetzner = &providerHetzner
	}
	providerKubevirt, ok := p.Kubevirt[providerName]
	if ok {
		cs.Kubevirt = &providerKubevirt
	}
	providerOpenstack, ok := p.Openstack[providerName]
	if ok {
		cs.Openstack = &providerOpenstack
	}
	providerPacket, ok := p.Packet[providerName]
	if ok {
		cs.Packet = &providerPacket
	}
	providerVsphere, ok := p.Vsphere[providerName]
	if ok {
		cs.Vsphere = &providerVsphere
	}

	return cs
}

// GetAllProviderNames returns a list of all provider names in use by the config
func (p *ProviderConfig) GetAllProviderNames() []string {
	providerNames := make([]string, 0)

	for name := range p.Alibaba {
		providerNames = append(providerNames, name)
	}
	for name := range p.Anexia {
		providerNames = append(providerNames, name)
	}
	for name := range p.Aws {
		providerNames = append(providerNames, name)
	}
	for name := range p.Azure {
		providerNames = append(providerNames, name)
	}
	for name := range p.Bringyourown {
		providerNames = append(providerNames, name)
	}
	for name := range p.Digitalocean {
		providerNames = append(providerNames, name)
	}
	for name := range p.Fake {
		providerNames = append(providerNames, name)
	}
	for name := range p.Gcp {
		providerNames = append(providerNames, name)
	}
	for name := range p.Hetzner {
		providerNames = append(providerNames, name)
	}
	for name := range p.Kubevirt {
		providerNames = append(providerNames, name)
	}
	for name := range p.Openstack {
		providerNames = append(providerNames, name)
	}
	for name := range p.Packet {
		providerNames = append(providerNames, name)
	}
	for name := range p.Vsphere {
		providerNames = append(providerNames, name)
	}

	return providerNames
}

// AddProviderConfig adds a new provider to the configuration
//	Note: Name must be unique
//	Returns an error, if the name is already in use to avoid ambigous naming.
func (p *ProviderConfig) AddProviderConfig(name string, cloudSpecType CloudSpecType, provider interface{}) error {
	if utils.IsOneOf(name, p.GetAllProviderNames()) {
		return fmt.Errorf("The provider name '%s' is already used", name)
	}

	switch cloudSpecType {

	case Alibaba:
		alibabaCloudSpec, ok := provider.(models.AlibabaCloudSpec)
		if !ok {
			return fmt.Errorf("Failed to cast provider into models.AlibabaCloudSpec")
		}
		if p.Alibaba == nil {
			p.Alibaba = make(map[string]models.AlibabaCloudSpec)
		}
		p.Alibaba[name] = alibabaCloudSpec

	case Anexia:
		anexiaCloudSpec, ok := provider.(models.AnexiaCloudSpec)
		if !ok {
			return fmt.Errorf("Failed to cast provider into models.AnexiaCloudSpec")
		}
		if p.Anexia == nil {
			p.Anexia = make(map[string]models.AnexiaCloudSpec)
		}
		p.Anexia[name] = anexiaCloudSpec

	case Aws:
		awsCloudSpec, ok := provider.(models.AWSCloudSpec)
		if !ok {
			return fmt.Errorf("Failed to cast provider into models.AwsCloudSpec")
		}
		if p.Aws == nil {
			p.Aws = make(map[string]models.AWSCloudSpec)
		}
		p.Aws[name] = awsCloudSpec

	case Azure:
		azureCloudSpec, ok := provider.(models.AzureCloudSpec)
		if !ok {
			return fmt.Errorf("Failed to cast provider into models.AzureCloudSpec")
		}
		if p.Azure == nil {
			p.Azure = make(map[string]models.AzureCloudSpec)
		}
		p.Azure[name] = azureCloudSpec

	case Bringyourown:
		bringyourownCloudSpec, ok := provider.(models.BringYourOwnCloudSpec)
		if !ok {
			return fmt.Errorf("Failed to cast provider into models.BringyourownCloudSpec")
		}
		if p.Bringyourown == nil {
			p.Bringyourown = make(map[string]models.BringYourOwnCloudSpec)
		}
		p.Bringyourown[name] = bringyourownCloudSpec

	case Digitalocean:
		digitaloceanCloudSpec, ok := provider.(models.DigitaloceanCloudSpec)
		if !ok {
			return fmt.Errorf("Failed to cast provider into models.DigitaloceanCloudSpec")
		}
		if p.Digitalocean == nil {
			p.Digitalocean = make(map[string]models.DigitaloceanCloudSpec)
		}
		p.Digitalocean[name] = digitaloceanCloudSpec

	case Fake:
		fakeCloudSpec, ok := provider.(models.FakeCloudSpec)
		if !ok {
			return fmt.Errorf("Failed to cast provider into models.FakeCloudSpec")
		}
		if p.Fake == nil {
			p.Fake = make(map[string]models.FakeCloudSpec)
		}
		p.Fake[name] = fakeCloudSpec

	case Gcp:
		gcpCloudSpec, ok := provider.(models.GCPCloudSpec)
		if !ok {
			return fmt.Errorf("Failed to cast provider into models.GcpCloudSpec")
		}
		if p.Gcp == nil {
			p.Gcp = make(map[string]models.GCPCloudSpec)
		}
		p.Gcp[name] = gcpCloudSpec

	case Hetzner:
		hetznerCloudSpec, ok := provider.(models.HetznerCloudSpec)
		if !ok {
			return fmt.Errorf("Failed to cast provider into models.HetznerCloudSpec")
		}
		if p.Hetzner == nil {
			p.Hetzner = make(map[string]models.HetznerCloudSpec)
		}
		p.Hetzner[name] = hetznerCloudSpec

	case Kubevirt:
		kubevirtCloudSpec, ok := provider.(models.KubevirtCloudSpec)
		if !ok {
			return fmt.Errorf("Failed to cast provider into models.KubevirtCloudSpec")
		}
		if p.Kubevirt == nil {
			p.Kubevirt = make(map[string]models.KubevirtCloudSpec)
		}
		p.Kubevirt[name] = kubevirtCloudSpec

	case Openstack:
		openstackCloudSpec, ok := provider.(models.OpenstackCloudSpec)
		if !ok {
			return fmt.Errorf("Failed to cast provider into models.OpenstackCloudSpec")
		}
		if p.Openstack == nil {
			p.Openstack = make(map[string]models.OpenstackCloudSpec)
		}
		p.Openstack[name] = openstackCloudSpec

	case Packet:
		packetCloudSpec, ok := provider.(models.PacketCloudSpec)
		if !ok {
			return fmt.Errorf("Failed to cast provider into models.PacketCloudSpec")
		}
		if p.Packet == nil {
			p.Packet = make(map[string]models.PacketCloudSpec)
		}
		p.Packet[name] = packetCloudSpec

	case Vsphere:
		vsphereCloudSpec, ok := provider.(models.VSphereCloudSpec)
		if !ok {
			return fmt.Errorf("Failed to cast provider into models.VsphereCloudSpec")
		}
		if p.Vsphere == nil {
			p.Vsphere = make(map[string]models.VSphereCloudSpec)
		}
		p.Vsphere[name] = vsphereCloudSpec
	}

	return nil
}
