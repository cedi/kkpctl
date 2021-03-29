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

// NewOSSpecConfig creates a new, empty OperatingSystemConfig
func NewOSSpecConfig() *OperatingSystemConfig {
	return &OperatingSystemConfig{
		CentOS: &models.CentOSSpec{
			DistUpgradeOnBoot: false,
		},
		Flatcar: &models.FlatcarSpec{
			DisableAutoUpdate: false,
		},
		Rhel: &models.RHELSpec{
			DistUpgradeOnBoot:               false,
			RHELSubscriptionManagerPassword: "",
			RHELSubscriptionManagerUser:     "",
			RHSMOfflineToken:                "",
		},
		Sles: &models.SLESSpec{
			DistUpgradeOnBoot: false,
		},
		Ubuntu: &models.UbuntuSpec{
			DistUpgradeOnBoot: false,
		},
	}
}

// SetOperatingSystemSpec sets the OperatingSystemType type to the OperatingSystemConfig
func (o *OperatingSystemConfig) SetOperatingSystemSpec(osType OperatingSystemType, osSpec interface{}) error {
	ok := false

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

// GetOperatingSystemSpec returns a *models.OperatingSystemSpec object from the OperatingSystemConfig
func (o *OperatingSystemConfig) GetOperatingSystemSpec() *models.OperatingSystemSpec {
	return &models.OperatingSystemSpec{
		Centos:  o.CentOS,
		Flatcar: o.Flatcar,
		Rhel:    o.Rhel,
		Sles:    o.Sles,
		Ubuntu:  o.Ubuntu,
	}
}

// GetValidOSSpecNames returns an []OperatingSystemType from the OperatingSystemConfig
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
