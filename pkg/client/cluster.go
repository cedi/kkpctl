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
func (c *Client) ListAllClusters(all bool) ([]model.ProjectCluster, error) {
	projects, err := c.ListProjects(all)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list all clusters: failed to list all projects")
	}

	projectCluster := make([]model.ProjectCluster, 0)
	for _, project := range projects {
		if project.ClustersNumber == 0 {
			continue
		}

		clusters, err := c.ListClusters(project.ID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to list all clusters")
		}

		projectCluster = append(projectCluster, clusters...)
	}

	return projectCluster, nil
}

// ListClusters lists all clusters for a given Project (identified by ID)
//	projectID the projectID in which to search clusters in
func (c *Client) ListClusters(projectID string) ([]model.ProjectCluster, error) {
	if projectID == "" {
		return nil, fmt.Errorf("failed to list clusters in project: no projectID specified")
	}

	requestURL := fmt.Sprintf("%s/%s/%s", projectPath, projectID, clusterPath)
	clusters := make([]models.Cluster, 0)
	_, err := c.Get(requestURL, &clusters, V2API)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list clusters in project %s", projectID)
	}

	projectCluster := make([]model.ProjectCluster, 0)
	for _, cluster := range clusters {
		projectCluster = append(projectCluster, *model.NewProjectCluster(projectID, cluster))
	}

	return projectCluster, nil
}

// GetClusterByID gets a clusters from any project
//	clusterID the clusterID to lookup
func (c *Client) GetClusterByID(clusterID string, all bool) (*model.ProjectCluster, error) {
	if clusterID == "" {
		return nil, fmt.Errorf("failed to get cluster: no clusterID specified")
	}

	projects, err := c.ListProjects(all)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list all clusters: failed to list all projects")
	}

	for _, project := range projects {
		if project.ClustersNumber == 0 {
			continue
		}

		clusters, err := c.ListClusters(project.ID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to list all clusters")
		}

		for _, cluster := range clusters {
			if cluster.Cluster.ID == clusterID {
				return &cluster, nil
			}
		}
	}

	return nil, fmt.Errorf("failed to get cluster %s: not found", clusterID)
}

// GetCluster gets a clusters in a given Project
//	clusterID the clusterID to lookup
//	projectID the projectID in which to search the cluster in
func (c *Client) GetCluster(id string, projectID string) (*model.ProjectCluster, error) {
	if projectID == "" {
		return nil, fmt.Errorf("failed to list clusters in project: no projectID specified")
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s", projectPath, projectID, clusterPath, id)
	cluster := &models.Cluster{}
	_, err := c.Get(requestURL, &cluster, V2API)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get cluster in project %s", projectID)
	}

	return model.NewProjectCluster(projectID, *cluster), nil
}

// CreateCluster creates a new cluster
func (c *Client) CreateCluster(newCluster *models.CreateClusterSpec, projectID string) (*models.Cluster, error) {
	if newCluster == nil {
		return nil, fmt.Errorf("failed to create cluster: cluster object is nil")
	}

	if projectID == "" {
		return nil, fmt.Errorf("failed to create cluster: no projectID given")
	}

	requestURL := fmt.Sprintf("%s/%s/%s", projectPath, projectID, clusterPath)
	cluster := &models.Cluster{}
	_, err := c.Post(requestURL, contentTypeJSON, newCluster, cluster, V2API)
	if err != nil {
		return cluster, errors.Wrapf(err, "failed to create cluster %s in project %s", newCluster.Cluster.Name, projectID)
	}

	return cluster, nil
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
	if clusterID == "" {
		return nil, fmt.Errorf("failed to get cluster health: no clusterID given")
	}

	if projectID == "" {
		return nil, fmt.Errorf("failed to get cluster health: no projectID given")
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		clusterPath,
		clusterID,
		clusterHealthPath,
	)

	clusterHealdh := &models.ClusterHealth{}
	_, err := c.Get(requestURL, &clusterHealdh, V2API)
	return clusterHealdh, errors.Wrapf(err, "failed to get cluster health for cluster %s in project %s", clusterID, projectID)
}

// GetClusterUpgradeVersions upgrades a cluster to a specified version
func (c *Client) GetClusterUpgradeVersions(clusterID string, projectID string) (model.VersionList, error) {
	if clusterID == "" {
		return nil, fmt.Errorf("failed to get upgradeable versions: no clusterID given")
	}

	if projectID == "" {
		return nil, fmt.Errorf("failed to get upgradeable versions: no projectID given")
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		clusterPath,
		clusterID,
		upgradesPath,
	)

	versionList := make(model.VersionList, 0)
	_, err := c.Get(requestURL, &versionList, V2API)
	return versionList, errors.Wrapf(err, "failed to get upgradeable versions for cluster %s project %s", clusterID, projectID)
}

// UpgradeCluster upgrades a cluster to a specified version
func (c *Client) UpgradeCluster(toVersion string, clusterID string, projectID string) (*model.ProjectCluster, error) {
	if clusterID == "" {
		return nil, fmt.Errorf("failed to upgrade cluster: no clusterID given")
	}

	if projectID == "" {
		return nil, fmt.Errorf("failed to upgrade cluster: no projectID given")
	}

	clusterProject, err := c.GetCluster(clusterID, projectID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to upgrade cluster %s in project %s", clusterID, projectID)
	}

	clusterProject.Cluster.Spec.Version = models.Semver(toVersion)

	requestURL := fmt.Sprintf("%s/%s/%s/%s", projectPath, projectID, clusterPath, clusterID)
	result := &models.Cluster{}
	_, err = c.Patch(requestURL, contentTypeJSON, clusterProject.Cluster, result, V2API)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to upgrade cluster %s in project %s", clusterID, projectID)
	}

	return model.NewProjectCluster(projectID, *result), nil
}
