package model

import "github.com/kubermatic/go-kubermatic/models"

// GetProviderNameFromCloudSpec returns the name of the provider from a given models.CloudSpec object
//gocyclo:ignore
func GetProviderNameFromCloudSpec(cloudSpec *models.CloudSpec) string {
	if cloudSpec.Alibaba != nil {
		return "Alibaba"
	}
	if cloudSpec.Anexia != nil {
		return "Anexia"
	}
	if cloudSpec.Aws != nil {
		return "AWS"
	}
	if cloudSpec.Azure != nil {
		return "Azure"
	}
	if cloudSpec.Bringyourown != nil {
		return "BringYourOwn"
	}
	if cloudSpec.Digitalocean != nil {
		return "DigitalOcean"
	}
	if cloudSpec.Fake != nil {
		return "Fake"
	}
	if cloudSpec.Gcp != nil {
		return "GCP"
	}
	if cloudSpec.Hetzner != nil {
		return "Hetzner"
	}
	if cloudSpec.Kubevirt != nil {
		return "Kubevirt"
	}
	if cloudSpec.Openstack != nil {
		return "OpenStack"
	}
	if cloudSpec.Packet != nil {
		return "Packet"
	}
	if cloudSpec.Vsphere != nil {
		return "Vsphere"
	}

	return ""
}
