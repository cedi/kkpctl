package client

import (
	"fmt"

	"github.com/kubermatic/go-kubermatic/models"
)

const (
	nodeDeploymentPath string = "nodedeployments"
)

// GetNodeDeployments lists all node deployments for a cluster
func (c *Client) GetNodeDeployments(clusterID string, projectID string, dc string) ([]models.NodeDeployment, error) {
	result := make([]models.NodeDeployment, 0)

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s/%s",
		projectPath,
		projectID,
		datacenterPath,
		dc,
		clusterPath,
		clusterID,
		nodeDeploymentPath,
	)
	_, err := c.Get(requestURL, &result)
	return result, err
}

func (c *Client) CreateNodeDeployment(newNodeDeployment *models.NodeDeploymentSpec, clusterID string, projectID string, dc string) (models.NodeDeployment, error) {
	result := models.NodeDeployment{}

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s/%s",
		projectPath,
		projectID,
		datacenterPath,
		dc,
		clusterPath,
		clusterID,
		nodeDeploymentPath,
	)

	_, err := c.Post(requestURL, contentTypeJSON, newNodeDeployment, result)
	return result, err
}
