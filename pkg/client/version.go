package client

import (
	"github.com/cedi/kkpctl/pkg/model"
)

// GetKKPVersion gets a clusters in a given Project in a given datacenter
func (c *Client) GetKKPVersion() (map[string]string, error) {
	result := make(map[string]string)

	_, err := c.Get("/version", &result, V1API)
	if err != nil {
		return result, err
	}

	return result, nil
}

// ListClusterVersions gets all K8s versions available in this KKP installation
func (c *Client) ListClusterVersions() (model.VersionList, error) {
	result := make([]model.Version, 0)

	_, err := c.Get("/upgrades/cluster", &result, V1API)
	if err != nil {
		return result, err
	}

	return result, nil
}
