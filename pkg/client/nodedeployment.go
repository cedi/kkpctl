package client

import (
	"fmt"

	"github.com/kubermatic/go-kubermatic/models"
)

const (
	nodeDeploymentPath string = "nodedeployments"
)

// ListNodeDeployments lists all node deployments for a cluster
func (c *Client) ListNodeDeployments(clusterID string, projectID string, dc string) ([]models.NodeDeployment, error) {
	var err error
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
	_, err = c.Get(requestURL, &result)
	return result, err
}
