package client

import (
	"github.com/cedi/kkpctl/pkg/model"
)

// GetClusterVersions gets a clusters in a given Project in a given datacenter
func (c *Client) GetClusterVersions() (model.VersionList, error) {
	result := make([]model.Version, 0)

	_, err := c.Get("/api/v1/upgrades/cluster", &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
