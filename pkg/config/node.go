package config

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/utils"
	"github.com/kubermatic/go-kubermatic/models"
)

// OperatingSystemType is a string, describing the operation systm type
type OperatingSystemType string

const (
	// CentOS represents the CentOS Operating system
	CentOS OperatingSystemType = "centos"

	// Flatcar represents the Flatcar operating system
	Flatcar OperatingSystemType = "flatcar"

	// Rhel represents the Rhel operating system
	Rhel OperatingSystemType = "rhel"

	// Sles represents the Sles operating system
	Sles OperatingSystemType = "sles"

	// Ubuntu represents the Ubuntu operating system
	Ubuntu OperatingSystemType = "ubuntu"
)

// CloudNodeConfig is used to identify the node spec for a cloud
type CloudNodeConfig struct {
	Alibaba      map[string]*models.AlibabaNodeSpec      `json:"alibaba,omitempty"`
	Anexia       map[string]*models.AnexiaNodeSpec       `json:"anexia,omitempty"`
	Aws          map[string]*models.AWSNodeSpec          `json:"aws,omitempty"`
	Azure        map[string]*models.AzureNodeSpec        `json:"azure,omitempty"`
	Digitalocean map[string]*models.DigitaloceanNodeSpec `json:"digitalocean,omitempty"`
	Gcp          map[string]*models.GCPNodeSpec          `json:"gcp,omitempty"`
	Hetzner      map[string]*models.HetznerNodeSpec      `json:"hetzner,omitempty"`
	Kubevirt     map[string]*models.KubevirtNodeSpec     `json:"kubevirt,omitempty"`
	Openstack    map[string]*models.OpenstackNodeSpec    `json:"openstack,omitempty"`
	Packet       map[string]*models.PacketNodeSpec       `json:"packet,omitempty"`
	Vsphere      map[string]*models.VSphereNodeSpec      `json:"vsphere,omitempty"`
}

// NewNodeSpecConfig generates a new, empty, NodeSpec config
func NewNodeSpecConfig() *CloudNodeConfig {
	nodeSpecConfig := &CloudNodeConfig{
		Alibaba:      map[string]*models.AlibabaNodeSpec{},
		Anexia:       map[string]*models.AnexiaNodeSpec{},
		Aws:          map[string]*models.AWSNodeSpec{},
		Azure:        map[string]*models.AzureNodeSpec{},
		Digitalocean: map[string]*models.DigitaloceanNodeSpec{},
		Gcp:          map[string]*models.GCPNodeSpec{},
		Hetzner:      map[string]*models.HetznerNodeSpec{},
		Kubevirt:     map[string]*models.KubevirtNodeSpec{},
		Openstack:    map[string]*models.OpenstackNodeSpec{},
		Packet:       map[string]*models.PacketNodeSpec{},
		Vsphere:      map[string]*models.VSphereNodeSpec{},
	}

	nodeSpecConfig.Alibaba["Alibaba"] = newAlibabaNodeSpec()
	nodeSpecConfig.Anexia["Anexia"] = newAnexiaNodeSpec()
	nodeSpecConfig.Aws["Aws"] = newAWSNodeSpec()
	nodeSpecConfig.Azure["Azure"] = newAzureNodeSpec()
	nodeSpecConfig.Digitalocean["Digitalocean"] = newDigitaloceanNodeSpec()
	nodeSpecConfig.Gcp["Gcp"] = newGCPNodeSpec()
	nodeSpecConfig.Hetzner["Hetzner"] = newHetznerNodeSpec()
	nodeSpecConfig.Kubevirt["Kubevirt"] = newKubevirtNodeSpec()
	nodeSpecConfig.Openstack["Openstack"] = newOpenstackNodeSpec()
	nodeSpecConfig.Packet["Packet"] = newPacketNodeSpec()
	nodeSpecConfig.Vsphere["Vsphere"] = newVSphereNodeSpec()

	return nodeSpecConfig
}

// GetAllNodeSpecNames returns a list of all provider names in use by the config
//gocyclo:ignore
func (c *CloudNodeConfig) GetAllNodeSpecNames() []string {
	providerNames := make([]string, 0)

	for name := range c.Alibaba {
		providerNames = append(providerNames, name)
	}
	for name := range c.Anexia {
		providerNames = append(providerNames, name)
	}
	for name := range c.Aws {
		providerNames = append(providerNames, name)
	}
	for name := range c.Azure {
		providerNames = append(providerNames, name)
	}
	for name := range c.Digitalocean {
		providerNames = append(providerNames, name)
	}
	for name := range c.Gcp {
		providerNames = append(providerNames, name)
	}
	for name := range c.Hetzner {
		providerNames = append(providerNames, name)
	}
	for name := range c.Kubevirt {
		providerNames = append(providerNames, name)
	}
	for name := range c.Openstack {
		providerNames = append(providerNames, name)
	}
	for name := range c.Packet {
		providerNames = append(providerNames, name)
	}
	for name := range c.Vsphere {
		providerNames = append(providerNames, name)
	}

	return providerNames
}

// AddCloudNodeSpec adds a new CloudNodeSpec to the configuration
//gocyclo:ignore
func (c *CloudNodeConfig) AddCloudNodeSpec(name string, nodeSpec interface{}) error {
	if utils.IsOneOf(name, c.GetAllNodeSpecNames()) {
		return fmt.Errorf("the nodespec name '%s' is already used", name)
	}

	switch cloudNodeSpec := nodeSpec.(type) {
	case *models.AlibabaNodeSpec:
		if c.Alibaba == nil {
			c.Alibaba = make(map[string]*models.AlibabaNodeSpec)
		}
		c.Alibaba[name] = cloudNodeSpec
	case *models.AnexiaNodeSpec:
		if c.Anexia == nil {
			c.Anexia = make(map[string]*models.AnexiaNodeSpec)
		}
		c.Anexia[name] = cloudNodeSpec
	case *models.AWSNodeSpec:
		if c.Aws == nil {
			c.Aws = make(map[string]*models.AWSNodeSpec)
		}
		c.Aws[name] = cloudNodeSpec
	case *models.AzureNodeSpec:
		if c.Azure == nil {
			c.Azure = make(map[string]*models.AzureNodeSpec)
		}
		c.Azure[name] = cloudNodeSpec
	case *models.DigitaloceanNodeSpec:
		if c.Digitalocean == nil {
			c.Digitalocean = make(map[string]*models.DigitaloceanNodeSpec)
		}
		c.Digitalocean[name] = cloudNodeSpec
	case *models.GCPNodeSpec:
		if c.Gcp == nil {
			c.Gcp = make(map[string]*models.GCPNodeSpec)
		}
		c.Gcp[name] = cloudNodeSpec
	case *models.HetznerNodeSpec:
		if c.Hetzner == nil {
			c.Hetzner = make(map[string]*models.HetznerNodeSpec)
		}
		c.Hetzner[name] = cloudNodeSpec
	case *models.KubevirtNodeSpec:
		if c.Kubevirt == nil {
			c.Kubevirt = make(map[string]*models.KubevirtNodeSpec)
		}
		c.Kubevirt[name] = cloudNodeSpec
	case *models.OpenstackNodeSpec:
		if c.Openstack == nil {
			c.Openstack = make(map[string]*models.OpenstackNodeSpec)
		}
		c.Openstack[name] = cloudNodeSpec
	case *models.PacketNodeSpec:
		if c.Packet == nil {
			c.Packet = make(map[string]*models.PacketNodeSpec)
		}
		c.Packet[name] = cloudNodeSpec
	case *models.VSphereNodeSpec:
		if c.Vsphere == nil {
			c.Vsphere = make(map[string]*models.VSphereNodeSpec)
		}
		c.Vsphere[name] = cloudNodeSpec
	default:
		return fmt.Errorf("unable to use nodeSpec")
	}

	return nil
}

// GetNodeCloudSpec gets the *models.NodeCloudSpec for the specified name from the CloudNodeConfig
//gocyclo:ignore
func (c *CloudNodeConfig) GetNodeCloudSpec(name string) *models.NodeCloudSpec {
	nodeCloudSpec := &models.NodeCloudSpec{}

	providerAlibaba, ok := c.Alibaba[name]
	if ok {
		nodeCloudSpec.Alibaba = providerAlibaba
	}
	providerAnexia, ok := c.Anexia[name]
	if ok {
		nodeCloudSpec.Anexia = providerAnexia
	}
	providerAws, ok := c.Aws[name]
	if ok {
		nodeCloudSpec.Aws = providerAws
	}
	providerAzure, ok := c.Azure[name]
	if ok {
		nodeCloudSpec.Azure = providerAzure
	}
	providerDigitalocean, ok := c.Digitalocean[name]
	if ok {
		nodeCloudSpec.Digitalocean = providerDigitalocean
	}
	providerGcp, ok := c.Gcp[name]
	if ok {
		nodeCloudSpec.Gcp = providerGcp
	}
	providerHetzner, ok := c.Hetzner[name]
	if ok {
		nodeCloudSpec.Hetzner = providerHetzner
	}
	providerKubevirt, ok := c.Kubevirt[name]
	if ok {
		nodeCloudSpec.Kubevirt = providerKubevirt
	}
	providerOpenstack, ok := c.Openstack[name]
	if ok {
		nodeCloudSpec.Openstack = providerOpenstack
	}
	providerPacket, ok := c.Packet[name]
	if ok {
		nodeCloudSpec.Packet = providerPacket
	}
	providerVsphere, ok := c.Vsphere[name]
	if ok {
		nodeCloudSpec.Vsphere = providerVsphere
	}

	return nodeCloudSpec
}

func newAlibabaNodeSpec() *models.AlibabaNodeSpec {
	return &models.AlibabaNodeSpec{
		DiskSize:                "",
		DiskType:                "",
		InstanceType:            "",
		InternetMaxBandwidthOut: "",
		Labels:                  map[string]string{},
		VSwitchID:               "",
		ZoneID:                  "",
	}
}
func newAnexiaNodeSpec() *models.AnexiaNodeSpec {
	return &models.AnexiaNodeSpec{
		CPUs:       new(int64),
		DiskSize:   new(int64),
		Memory:     new(int64),
		TemplateID: new(string),
		VlanID:     new(string),
	}
}
func newAWSNodeSpec() *models.AWSNodeSpec {
	return &models.AWSNodeSpec{
		AMI:              "",
		AssignPublicIP:   false,
		AvailabilityZone: "",
		InstanceType:     new(string),
		SubnetID:         "",
		Tags:             map[string]string{},
		VolumeSize:       new(int64),
		VolumeType:       new(string),
	}
}
func newAzureNodeSpec() *models.AzureNodeSpec {
	return &models.AzureNodeSpec{
		AssignPublicIP: false,
		DataDiskSize:   0,
		ImageID:        "",
		OSDiskSize:     0,
		Size:           new(string),
		Tags:           map[string]string{},
		Zones:          []string{},
	}
}
func newDigitaloceanNodeSpec() *models.DigitaloceanNodeSpec {
	return &models.DigitaloceanNodeSpec{
		Backups:    false,
		IPV6:       false,
		Monitoring: false,
		Size:       new(string),
		Tags:       []string{},
	}
}
func newGCPNodeSpec() *models.GCPNodeSpec {
	return &models.GCPNodeSpec{
		CustomImage: "",
		DiskSize:    0,
		DiskType:    "",
		Labels:      map[string]string{},
		MachineType: "",
		Preemptible: false,
		Tags:        []string{},
		Zone:        "",
	}
}
func newHetznerNodeSpec() *models.HetznerNodeSpec {
	return &models.HetznerNodeSpec{
		Network: "",
		Type:    new(string),
	}
}
func newKubevirtNodeSpec() *models.KubevirtNodeSpec {
	return &models.KubevirtNodeSpec{
		CPUs:             new(string),
		Memory:           new(string),
		Namespace:        new(string),
		PVCSize:          new(string),
		SourceURL:        new(string),
		StorageClassName: new(string),
	}
}
func newOpenstackNodeSpec() *models.OpenstackNodeSpec {
	return &models.OpenstackNodeSpec{
		AvailabilityZone:          "",
		Flavor:                    new(string),
		Image:                     new(string),
		InstanceReadyCheckPeriod:  "",
		InstanceReadyCheckTimeout: "",
		RootDiskSizeGB:            0,
		Tags:                      map[string]string{},
		UseFloatingIP:             false,
	}
}
func newPacketNodeSpec() *models.PacketNodeSpec {
	return &models.PacketNodeSpec{
		InstanceType: new(string),
		Tags:         []string{},
	}
}
func newVSphereNodeSpec() *models.VSphereNodeSpec {
	return &models.VSphereNodeSpec{
		CPUs:       0,
		DiskSizeGB: 0,
		Memory:     0,
		Template:   "",
	}
}
