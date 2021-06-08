package client

import (
	"fmt"

	"github.com/kubermatic/go-kubermatic/models"
)

const (
	machineDeploymentPath string = "machinedeployments"
	nodePath              string = "nodes"
)

// GetMachineDeployments lists all machine deployments for a cluster
func (c *Client) GetMachineDeployments(clusterID string, projectID string) ([]models.NodeDeployment, error) {
	result := make([]models.NodeDeployment, 0)

	requestURL := fmt.Sprintf("%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		clusterPath,
		clusterID,
		machineDeploymentPath,
	)
	_, err := c.Get(requestURL, &result)
	return result, err
}

// GetMachineDeployment lists all machine deployments for a cluster
func (c *Client) GetMachineDeployment(id string, clusterID string, projectID string) (*models.NodeDeployment, error) {
	result := &models.NodeDeployment{}

	requestURL := fmt.Sprintf("%s/%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		clusterPath,
		clusterID,
		machineDeploymentPath,
		id,
	)
	_, err := c.Get(requestURL, &result)
	return result, err
}

// CreateMachineDeployment creates a new machine deployment on a cluster
func (c *Client) CreateMachineDeployment(new *models.NodeDeployment, clusterID string, projectID string) (*models.NodeDeployment, error) {
	result := &models.NodeDeployment{}

	requestURL := fmt.Sprintf("%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		clusterPath,
		clusterID,
		machineDeploymentPath,
	)

	_, err := c.Post(requestURL, contentTypeJSON, new, result)
	return result, err
}

// DeleteMachineDeployment delets a machine deployment from a cluster
func (c *Client) DeleteMachineDeployment(id string, clusterID string, projectID string) error {
	requestURL := fmt.Sprintf("%s/%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		clusterPath,
		clusterID,
		machineDeploymentPath,
		id,
	)

	_, err := c.Delete(requestURL)
	return err
}

// GetMachineDeploymentNodes gets all nodes in a machine deployment
func (c *Client) GetMachineDeploymentNodes(id string, clusterID string, projectID string) ([]models.Node, error) {
	result := make([]models.Node, 0)

	requestURL := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		clusterPath,
		clusterID,
		machineDeploymentPath,
		id,
		nodePath,
	)

	_, err := c.Get(requestURL, &result)
	return result, err
}

// GetMachineDeploymentEvents gets all nodes in a machine deployment
func (c *Client) GetMachineDeploymentEvents(id string, clusterID string, projectID string) ([]models.Event, error) {
	result := make([]models.Event, 0)

	requestURL := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		clusterPath,
		clusterID,
		machineDeploymentPath,
		id,
		nodePath,
		eventsPath,
	)
	_, err := c.Get(requestURL, &result)
	return result, err
}

// UpgradeWorkerDeploymentVersion upgrades node deployments in a cluster
func (c *Client) UpgradeWorkerDeploymentVersion(toVersion string, clusterID string, projectID string) error {
	requestURL := fmt.Sprintf("%s/%s/%s/%s/%s/%s",
		projectPath,
		projectID,
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
func (c *Client) AreAllWorkerDeploymentsReady(clusterID string, projectID string) (bool, error) {
	machineDeployments, err := c.GetMachineDeployments(clusterID, projectID)
	if err != nil {
		return false, err
	}

	for _, machineDeployment := range machineDeployments {
		if machineDeployment.Status.UnavailableReplicas != 0 {
			return false, nil
		}
	}

	return true, nil
}

// IsWorkerDeploymentsReady gets the specified worker deployment and checks if the unavailable replica count is 0
func (c *Client) IsWorkerDeploymentsReady(id string, clusterID string, projectID string) (bool, error) {
	machineDeployment, err := c.GetMachineDeployment(id, clusterID, projectID)
	if err != nil {
		return false, err
	}

	return machineDeployment.Status.UnavailableReplicas == 0, nil
}
