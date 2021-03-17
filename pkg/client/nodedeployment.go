package client

import (
	"fmt"

	"github.com/kubermatic/go-kubermatic/models"
)

const (
	nodeDeploymentPath string = "nodedeployments"
	nodePath           string = "nodes"
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

// GetNodeDeployment lists all node deployments for a cluster
func (c *Client) GetNodeDeployment(nodeDeploymentID string, clusterID string, projectID string, dc string) (models.NodeDeployment, error) {
	result := models.NodeDeployment{}

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		datacenterPath,
		dc,
		clusterPath,
		clusterID,
		nodeDeploymentPath,
		nodeDeploymentID,
	)
	_, err := c.Get(requestURL, &result)
	return result, err
}

func (c *Client) CreateNodeDeployment(newNodeDeployment *models.NodeDeployment, clusterID string, projectID string, dc string) (models.NodeDeployment, error) {
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

func (c *Client) DeleteNodeDeployment(nodeDeploymentID string, clusterID string, projectID string, dc string) error {
	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		datacenterPath,
		dc,
		clusterPath,
		clusterID,
		nodeDeploymentPath,
		nodeDeploymentID,
	)

	_, err := c.Delete(requestURL)
	return err
}

// GetNodeDeploymentNodes gets all nodes in a node deployment
func (c *Client) GetNodeDeploymentNodes(nodeDeploymentID string, clusterID string, projectID string, dc string) ([]models.Node, error) {
	result := make([]models.Node, 0)

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		datacenterPath,
		dc,
		clusterPath,
		clusterID,
		nodeDeploymentPath,
		nodeDeploymentID,
		nodePath,
	)
	_, err := c.Get(requestURL, &result)
	return result, err
}

// GetNodeDeploymentNodes gets all nodes in a node deployment
func (c *Client) GetNodeDeploymentEvents(nodeDeploymentID string, clusterID string, projectID string, dc string) ([]models.Event, error) {
	result := make([]models.Event, 0)

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		datacenterPath,
		dc,
		clusterPath,
		clusterID,
		nodeDeploymentPath,
		nodeDeploymentID,
		nodePath,
		eventsPath,
	)
	_, err := c.Get(requestURL, &result)
	return result, err
}

// GetNodeDeploymentNodes gets all nodes in a node deployment
func (c *Client) UpgradeWorkerDeploymentVersion(toVersion string, clusterID string, projectID string, dc string) error {

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		datacenterPath,
		dc,
		clusterPath,
		clusterID,
		nodePath,
		upgradesPath,
	)

	updateVersion := models.MasterVersion{
		Default:                    true,
		RestrictedByKubeletVersion: true,
		Version:                    toVersion,
	}

	_, err := c.Put(requestURL, contentTypeJSON, updateVersion, nil)
	return err
}
