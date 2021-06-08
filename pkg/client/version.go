package client

import (
	"github.com/cedi/kkpctl/pkg/model"
)

// ListClusterVersions gets a clusters in a given Project in a given datacenter
func (c *Client) ListClusterVersions() (model.VersionList, error) {
	result := make([]model.Version, 0)

	_, err := c.Get("/upgrades/cluster", &result, V1API)
	if err != nil {
		return result, err
	}

	return result, nil
}
