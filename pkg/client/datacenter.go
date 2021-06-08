package client

import (
	"fmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListDatacenter lists all datacenters available in KKP
func (c *Client) ListDatacenter() ([]models.Datacenter, error) {
	var err error
	result := make([]models.Datacenter, 0)

	requestURL := fmt.Sprintf("/%s", datacenterPath)
	_, err = c.Get(requestURL, &result, V1API)
	return result, err
}

// GetDatacenter gets a specific datacenter
func (c *Client) GetDatacenter(name string) (*models.Datacenter, error) {
	var err error
	result := &models.Datacenter{}

	requestURL := fmt.Sprintf("/%s/%s", datacenterPath, name)
	_, err = c.Get(requestURL, result, V1API)
	return result, err
}
