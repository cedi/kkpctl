package client

import (
	"fmt"
	"io/ioutil"

	"github.com/cedi/kkpctl/pkg/model"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/pkg/errors"
)

const (
	clusterPath       string = "clusters"
	datacenterPath    string = "dc"
	kubeconfigPath    string = "kubeconfig"
	clusterHealthPath string = "health"
	upgradesPath      string = "upgrades"
)

// ListAllClusters lists all clusters
//	all lists all clusters in all projects, if you have the permission to do so
func (c *Client) ListAllClusters(all bool) ([]models.Cluster, error) {
	result := make([]models.Cluster, 0)

	projects, err := c.ListProjects(all)
	if err != nil {
		return result, errors.Wrap(err, "failed to list all clusters: failed to list all projects")
	}

	for _, project := range projects {
		clusters, err := c.ListClusters(project.ID)
		if err != nil {
			return result, errors.Wrap(err, "failed to list all clusters")
		}

		result = append(result, clusters...)
	}

	return result, nil
}

// ListClusters lists all clusters for a given Project (identified by ID)
//	projectID the projectID in which to search clusters in
func (c *Client) ListClusters(projectID string) ([]models.Cluster, error) {
	result := make([]models.Cluster, 0)

	if projectID == "" {
		return result, fmt.Errorf("failed to list clusters in project: no projectID specified")
	}

	requestURL := fmt.Sprintf("%s/%s/%s", projectPath, projectID, clusterPath)
	_, err := c.Get(requestURL, &result, V2API)
	if err != nil {
		return result, errors.Wrapf(err, "failed to list clusters in project %s", projectID)
	}

	return result, nil
}

// GetClusterByID gets a clusters from any project
//	clusterID the clusterID to lookup
func (c *Client) GetClusterByID(clusterID string) (*models.Cluster, error) {
	result := &models.Cluster{}

	if clusterID == "" {
		return result, fmt.Errorf("failed to get cluster: no clusterID specified")
	}

	clusters, err := c.ListAllClusters(true)
	if err != nil {
		return result, errors.Wrapf(err, "failed to get cluster %s", clusterID)
	}

	found := false
	for _, cluster := range clusters {
		if cluster.ID == clusterID {
			result = &cluster
			found = true
			break
		}
	}

	if found {
		err = nil
	} else {
		err = fmt.Errorf("failed to get cluster %s: not found", clusterID)
	}

	return result, err
}

// GetCluster gets a clusters in a given Project
//	clusterID the clusterID to lookup
//	projectID the projectID in which to search the cluster in
func (c *Client) GetCluster(id string, projectID string) (*models.Cluster, error) {
	result := &models.Cluster{}

	if projectID == "" {
		return result, fmt.Errorf("failed to list clusters in project: no projectID specified")
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s", projectPath, projectID, clusterPath, id)
	_, err := c.Get(requestURL, &result, V2API)
	if err != nil {
		return result, errors.Wrapf(err, "failed to get cluster in project %s", projectID)
	}

	return result, nil
}

// CreateCluster creates a new cluster
func (c *Client) CreateCluster(newCluster *models.CreateClusterSpec, projectID string) (models.Cluster, error) {
	result := models.Cluster{}

	if newCluster == nil {
		return result, fmt.Errorf("failed to create cluster: cluster object is nil")
	}

	if projectID == "" {
		return result, fmt.Errorf("failed to create cluster: no projectID given")
	}

	requestURL := fmt.Sprintf("%s/%s/%s", projectPath, projectID, clusterPath)
	_, err := c.Post(requestURL, contentTypeJSON, newCluster, &result, V2API)
	if err != nil {
		return result, errors.Wrapf(err, "failed to create cluster %s in project %s", newCluster.Cluster.Name, projectID)
	}

	return result, nil
}

// DeleteCluster deletes a cluster identified by id
func (c *Client) DeleteCluster(clusterID string, projectID string, deleteVolumes bool, deleteLoadBalancers bool) error {
	if clusterID == "" {
		return fmt.Errorf("failed to delete cluster: no clusterID given")
	}

	if projectID == "" {
		return fmt.Errorf("failed to delete cluster: no projectID given")
	}

	headers := make(map[string]string)
	if deleteVolumes {
		headers["DeleteVolumes"] = "true"
	}

	if deleteLoadBalancers {
		headers["DeleteLoadBalancers"] = "true"
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s", projectPath, projectID, clusterPath, clusterID)
	_, err := c.DeleteWithHeader(requestURL, headers, V2API)
	if err != nil {
		return errors.Wrapf(err, "failed to delete cluster %s project %s", clusterID, projectID)
	}

	return nil
}

// GetKubeConfig gets a clusters Kubeconfig
func (c *Client) GetKubeConfig(clusterID string, projectID string) (string, error) {
	if clusterID == "" {
		return "", fmt.Errorf("failed to get kubeconfig: no clusterID given")
	}

	if projectID == "" {
		return "", fmt.Errorf("failed to get kubeconfig: no projectID given")
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s/%s", projectPath, projectID, clusterPath, clusterID, kubeconfigPath)
	resp, err := c.Get(requestURL, nil, V2API)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get kubeconfig for cluster %s project %s", clusterID, projectID)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get kubeconfig for cluster %s project %s", clusterID, projectID)
	}

	return string(body), nil
}

// GetClusterHealth returns the health status of a cluster
func (c *Client) GetClusterHealth(clusterID string, projectID string) (*models.ClusterHealth, error) {
	result := &models.ClusterHealth{}

	if clusterID == "" {
		return result, fmt.Errorf("failed to get cluster health: no clusterID given")
	}

	if projectID == "" {
		return result, fmt.Errorf("failed to get cluster health: no projectID given")
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		clusterPath,
		clusterID,
		clusterHealthPath,
	)
	_, err := c.Get(requestURL, &result, V2API)
	return result, errors.Wrapf(err, "failed to get cluster health for cluster %s in project %s", clusterID, projectID)
}

// GetClusterUpgradeVersions upgrades a cluster to a specified version
func (c *Client) GetClusterUpgradeVersions(clusterID string, projectID string) (model.VersionList, error) {
	result := make(model.VersionList, 0)

	if clusterID == "" {
		return result, fmt.Errorf("failed to get upgradeable versions: no clusterID given")
	}

	if projectID == "" {
		return result, fmt.Errorf("failed to get upgradeable versions: no projectID given")
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		clusterPath,
		clusterID,
		upgradesPath,
	)
	_, err := c.Get(requestURL, &result, V2API)
	return result, errors.Wrapf(err, "failed to get upgradeable versions for cluster %s project %s", clusterID, projectID)
}

// UpgradeCluster upgrades a cluster to a specified version
func (c *Client) UpgradeCluster(toVersion string, clusterID string, projectID string) (*models.Cluster, error) {
	result := &models.Cluster{}

	if clusterID == "" {
		return result, fmt.Errorf("failed to upgrade cluster: no clusterID given")
	}

	if projectID == "" {
		return result, fmt.Errorf("failed to upgrade cluster: no projectID given")
	}

	cluster, err := c.GetCluster(clusterID, projectID)
	if err != nil {
		return result, errors.Wrapf(err, "failed to upgrade cluster %s in project %s", clusterID, projectID)
	}

	cluster.Spec.Version = toVersion

	requestURL := fmt.Sprintf("%s/%s/%s/%s", projectPath, projectID, clusterPath, clusterID)
	_, err = c.Patch(requestURL, contentTypeJSON, cluster, result, V2API)
	return result, errors.Wrapf(err, "failed to upgrade cluster %s in project %s", clusterID, projectID)
}
