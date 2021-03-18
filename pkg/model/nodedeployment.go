package model

import (
	"github.com/kubermatic/go-kubermatic/models"
)

func NewNodeDeployment(name string, version string, replicas int32, dynamicConfig bool, cloudSpec *models.NodeCloudSpec, osSpec *models.OperatingSystemSpec, labels map[string]string) *models.NodeDeployment {
	return &models.NodeDeployment{
		Name: name,
		Spec: &models.NodeDeploymentSpec{
			DynamicConfig: dynamicConfig,
			Replicas:      &replicas,
			Template: &models.NodeSpec{
				Labels:          labels,
				SSHUserName:     "",
				Taints:          []*models.TaintSpec{},
				Cloud:           cloudSpec,
				OperatingSystem: osSpec,
				Versions: &models.NodeVersionInfo{
					Kubelet: version,
				},
			},
		},
	}
}
