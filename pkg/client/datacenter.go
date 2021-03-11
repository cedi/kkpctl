package client

import (
	"fmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListDatacenter lists all node deployments for a cluster
func (c *Client) ListDatacenter() ([]models.Datacenter, error) {
	var err error
	result := make([]models.Datacenter, 0)

	requestURL := fmt.Sprintf("/%s", datacenterPath)
	_, err = c.Get(requestURL, &result)
	return result, err
}

// GetDatacenter lists all node deployments for a cluster
func (c *Client) GetDatacenter(name string) (models.Datacenter, error) {
	var err error
	result := models.Datacenter{}

	requestURL := fmt.Sprintf("/%s/%s", datacenterPath, name)
	_, err = c.Get(requestURL, &result)
	return result, err
}

// GetDatacenterForCluster gets the datacenter for a cluster
func (c *Client) GetDatacenterForCluster(clusterID string) (models.Datacenter, error) {
	cluster, err := c.GetCluster(clusterID, false)
	if err != nil {
		return models.Datacenter{}, err
	}

	datacenter, err := c.GetDatacenter(cluster.Spec.Cloud.DatacenterName)
	if err != nil {
		return models.Datacenter{}, err
	}

	return datacenter, nil
}

// GetDatacenterForClusterInProject gets the datacenter for a cluster
func (c *Client) GetDatacenterForClusterInProject(clusterID string, projectID string) (models.Datacenter, error) {
	cluster, err := c.GetClusterInProject(clusterID, projectID)
	if err != nil {
		return models.Datacenter{}, err
	}

	datacenter, err := c.GetDatacenter(cluster.Spec.Cloud.DatacenterName)
	if err != nil {
		return models.Datacenter{}, err
	}

	return datacenter, nil
}
