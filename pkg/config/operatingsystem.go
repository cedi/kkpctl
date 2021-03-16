package config

import (
	"fmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// OperatingSystemConfig is used to identify which OperationSystem to use
type OperatingSystemConfig struct {
	CentOS  *models.CentOSSpec  `json:"centos,omitempty"`
	Flatcar *models.FlatcarSpec `json:"flatcar,omitempty"`
	Rhel    *models.RHELSpec    `json:"rhel,omitempty"`
	Sles    *models.SLESSpec    `json:"sles,omitempty"`
	Ubuntu  *models.UbuntuSpec  `json:"ubuntu,omitempty"`
}

// AddProviderConfig adds a new provider to the configuration
func (o *OperatingSystemConfig) AddOperatingSystemSpec(osType OperatingSystemType, osSpec interface{}) error {
	var ok bool

	switch osType {
	case CentOS:
		o.CentOS, ok = osSpec.(*models.CentOSSpec)
	case Flatcar:
		o.Flatcar, ok = osSpec.(*models.FlatcarSpec)
	case Rhel:
		o.Rhel, ok = osSpec.(*models.RHELSpec)
	case Sles:
		o.Sles, ok = osSpec.(*models.SLESSpec)
	case Ubuntu:
		o.Ubuntu, ok = osSpec.(*models.UbuntuSpec)
	}

	if !ok {
		return fmt.Errorf("unable to use operating system type %v", osType)
	}

	return nil
}

func (o *OperatingSystemConfig) GetOperatingSystemSpec() *models.OperatingSystemSpec {
	return &models.OperatingSystemSpec{
		Centos:  o.CentOS,
		Flatcar: o.Flatcar,
		Rhel:    o.Rhel,
		Sles:    o.Sles,
		Ubuntu:  o.Ubuntu,
	}
}

func (o *OperatingSystemConfig) GetValidOSSpecNames() []OperatingSystemType {
	result := make([]OperatingSystemType, 0)

	if o.CentOS != nil {
		result = append(result, CentOS)
	}
	if o.Flatcar != nil {
		result = append(result, Flatcar)
	}
	if o.Rhel != nil {
		result = append(result, Rhel)
	}
	if o.Sles != nil {
		result = append(result, Sles)
	}
	if o.Ubuntu != nil {
		result = append(result, Ubuntu)
	}

	return result
}
