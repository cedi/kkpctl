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

// CreateNodeDeployment creates a new node deployment on a cluster
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

// DeleteNodeDeployment delets a node deployment from a cluster
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

// GetNodeDeploymentEvents gets all nodes in a node deployment
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

// UpgradeWorkerDeploymentVersion gets all versions which are valid to upgrade to from a node deployment
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

// AreAllWorkerDeploymentsReady gets all worker deployments and checks if their unavailable replica count is 0
//	returns true if all WorkerDeployments are ready
func (c *Client) AreAllWorkerDeploymentsReady(clusterID string, projectID string, dc string) (bool, error) {
	var nodeDeployments []models.NodeDeployment

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s/%s",
		projectPath,
		projectID,
		datacenterPath,
		dc,
		clusterPath,
		clusterID,
		nodeDeploymentPath,
	)

	_, err := c.Get(requestURL, &nodeDeployments)
	if err != nil {
		return false, err
	}

	for _, nodeDeployment := range nodeDeployments {
		if nodeDeployment.Status.UnavailableReplicas != 0 {
			return false, nil
		}
	}

	return true, nil
}

// IsWorkerDeploymentsReady gets the specified worker deployment and checks if the unavailable replica count is 0
func (c *Client) IsWorkerDeploymentsReady(nodeDeploymentID string, clusterID string, projectID string, dc string) (bool, error) {
	var nodeDeployment models.NodeDeployment

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

	_, err := c.Get(requestURL, &nodeDeployment)
	if err != nil {
		return false, err
	}

	return nodeDeployment.Status.UnavailableReplicas == 0, nil
}
