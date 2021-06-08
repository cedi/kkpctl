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

// GetDatacenterForCluster gets the datacenter for a cluster
func (c *Client) GetDatacenterForCluster(clusterID string) (*models.Datacenter, error) {
	cluster, err := c.GetClusterByID(clusterID)
	if err != nil {
		return nil, err
	}

	datacenter, err := c.GetDatacenter(cluster.Spec.Cloud.DatacenterName)
	if err != nil {
		return nil, err
	}

	return datacenter, nil
}

// GetDatacenterForClusterInProject gets the datacenter for a cluster
func (c *Client) GetDatacenterForClusterInProject(clusterID string, projectID string) (*models.Datacenter, error) {
	cluster, err := c.GetCluster(clusterID, projectID)
	if err != nil {
		return nil, err
	}

	datacenter, err := c.GetDatacenter(cluster.Spec.Cloud.DatacenterName)
	if err != nil {
		return nil, err
	}

	return datacenter, nil
}
