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
		clusters, err := c.ListClustersInProject(project.ID)
		if err != nil {
			return result, errors.Wrap(err, "failed to list all clusters")
		}

		result = append(result, clusters...)
	}

	return result, nil
}

// ListClustersInDC lists all clusters for a given Project in a given datacenter
//	dc the datacenter in which to search clusters in
//	all lists all clusters in all projects, if you have the permission to do so
func (c *Client) ListClustersInDC(dc string, all bool) ([]models.Cluster, error) {
	result := make([]models.Cluster, 0)

	if dc == "" {
		return result, fmt.Errorf("failed to list clusters in datacenter: no datacenter specified")
	}

	projects, err := c.ListProjects(all)
	if err != nil {
		return result, errors.Wrapf(err, "failed to list clusters in datacenter %s", dc)
	}

	for _, project := range projects {
		clusters, err := c.ListClustersInProjectInDC(project.ID, dc)
		if err != nil {
			return result, errors.Wrapf(err, "failed to list clusters in datacenter %s", dc)
		}

		result = append(result, clusters...)
	}

	return result, nil
}

// ListClustersInProject lists all clusters for a given Project (identified by ID)
//	projectID the projectID in which to search clusters in
func (c *Client) ListClustersInProject(projectID string) ([]models.Cluster, error) {
	result := make([]models.Cluster, 0)

	if projectID == "" {
		return result, fmt.Errorf("failed to list clusters in project: no projectID specified")
	}

	requestURL := fmt.Sprintf("%s/%s/%s", projectPath, projectID, clusterPath)
	_, err := c.Get(requestURL, &result)
	if err != nil {
		return result, errors.Wrapf(err, "failed to list clusters in project %s", projectID)
	}

	return result, nil
}

// ListClustersInProjectInDC lists all clusters for a given Project in a given datacenter
//	projectID the projectID in which to search clusters in
//	dc the datacenter in which to search clusters in
func (c *Client) ListClustersInProjectInDC(projectID string, dc string) ([]models.Cluster, error) {
	result := make([]models.Cluster, 0)

	if projectID == "" && dc == "" {
		return c.ListAllClusters(false)
	} else if projectID != "" && dc == "" {
		return c.ListClustersInProject(projectID)
	}

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath)
	_, err := c.Get(requestURL, &result)
	if err != nil {
		return result, errors.Wrapf(err, "failed to list clusters in project %s in datacenter %s", projectID, dc)
	}

	return result, nil
}

// GetCluster gets a clusters in a given Project
//	clusterID the clusterID to lookup
func (c *Client) GetCluster(clusterID string, listAll bool) (models.Cluster, error) {
	var result models.Cluster

	if clusterID == "" {
		return result, fmt.Errorf("failed to get cluster: no clusterID specified")
	}

	clusters, err := c.ListAllClusters(listAll)
	if err != nil {
		return result, errors.Wrapf(err, "failed to get cluster %s", clusterID)
	}

	found := false
	for _, cluster := range clusters {
		if cluster.ID == clusterID {
			result = cluster
			found = true
			break
		}
	}

	if !found {
		return result, errors.Wrapf(err, "failed to get cluster %s: not found", clusterID)
	}

	return result, nil
}

// GetClusterInDC gets a clusters in a given Project
//	clusterID the clusterID to lookup
//	dc the datacenter in which to search for the cluster in
func (c *Client) GetClusterInDC(clusterID string, dc string, listAll bool) (models.Cluster, error) {
	var result models.Cluster

	if clusterID == "" {
		return result, fmt.Errorf("failed to get cluster in datacenter: no clusterID specified")
	}

	if dc == "" {
		return result, fmt.Errorf("failed to get cluster in datacenter: no datacenter specified")
	}

	clusters, err := c.ListClustersInDC(dc, listAll)
	if err != nil {
		return result, errors.Wrapf(err, "failed to get cluster %s in datacenter %s", clusterID, dc)
	}

	found := false
	for _, cluster := range clusters {
		if cluster.ID == clusterID {
			result = cluster
			found = true
			break
		}
	}

	if !found {
		return result, fmt.Errorf("failed to get cluster %s in datacenter %s: not found", clusterID, dc)
	}

	return result, nil
}

// GetClusterInProject gets a clusters in a given Project
//	clusterID the clusterID to lookup
//	projectID the projectID in which to search the cluster in
func (c *Client) GetClusterInProject(clusterID string, projectID string) (models.Cluster, error) {
	result := models.Cluster{}

	if clusterID == "" {
		return result, fmt.Errorf("failed to get cluster in project: no clusterID specified")
	}

	if projectID == "" {
		return result, fmt.Errorf("failed to get cluster in project: no projectID specified")
	}

	clusters, err := c.ListClustersInProject(projectID)
	if err != nil {
		return result, errors.Wrapf(err, "failed to get cluster %s in project %s", clusterID, projectID)
	}

	found := false
	for _, cluster := range clusters {
		if cluster.ID == clusterID {
			result = cluster
			found = true
			break
		}
	}

	if !found {
		return result, fmt.Errorf("failed to get cluster %s in project %s: not found", clusterID, projectID)
	}

	return result, nil
}

// GetClusterInProjectInDC gets a clusters in a given Project in a given datacenter
//	clusterID the clusterID to lookup
//	projectID the projectID in which to search the cluster in
//	dc the datacenter in which to search for the cluster in
func (c *Client) GetClusterInProjectInDC(clusterID string, projectID string, dc string) (models.Cluster, error) {
	result := models.Cluster{}

	if clusterID == "" {
		return result, fmt.Errorf("failed to get cluster in project in datacenter: no clusterID specified")
	}

	if projectID == "" && dc == "" {
		return c.GetCluster(clusterID, false)
	} else if projectID != "" && dc == "" {
		return c.GetClusterInProject(clusterID, projectID)
	} else if projectID == "" && dc != "" {
		return c.GetClusterInDC(clusterID, dc, false)
	}

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath, clusterID)
	_, err := c.Get(requestURL, &result)
	if err != nil {
		return result, errors.Wrapf(err, "failed to get cluster %s in project %s in datacenter %s", clusterID, projectID, dc)
	}

	return result, nil
}

// CreateCluster creates a new cluster
func (c *Client) CreateCluster(newCluster *models.CreateClusterSpec, projectID string, dc string) (models.Cluster, error) {
	result := models.Cluster{}

	if newCluster == nil {
		return result, fmt.Errorf("failed to create cluster: cluster object is nil")
	}

	if projectID == "" {
		return result, fmt.Errorf("failed to create cluster: no projectID given")
	}

	if dc == "" {
		return result, fmt.Errorf("failed to create cluster: no datacenter given")
	}

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath)
	_, err := c.Post(requestURL, contentTypeJSON, newCluster, result)
	if err != nil {
		return result, errors.Wrapf(err, "failed to create cluster %s in project %s in datacenter %s", newCluster.Cluster.Name, projectID, dc)
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

	clusters, err := c.ListClustersInProject(projectID)
	if err != nil {
		return errors.Wrapf(err, "failed to delete cluster %s in project %s", clusterID, projectID)
	}

	datacenter := ""
	for _, cluster := range clusters {
		if cluster.ID == clusterID {
			datacenter = cluster.Spec.Cloud.DatacenterName
		}
	}

	if datacenter == "" {
		return errors.Wrapf(err, "failed to delete cluster %s in project %s: not found", clusterID, projectID)
	}

	return c.DeleteClusterInDC(clusterID, projectID, datacenter, deleteVolumes, deleteLoadBalancers)
}

// DeleteClusterInDC deletes a cluster identified by id
func (c *Client) DeleteClusterInDC(clusterID string, projectID string, dc string, deleteVolumes bool, deleteLoadBalancers bool) error {
	if clusterID == "" {
		return fmt.Errorf("failed to delete cluster: no clusterID given")
	}

	if projectID == "" {
		return fmt.Errorf("failed to delete cluster: no projectID given")
	}

	if dc == "" {
		return c.DeleteCluster(clusterID, projectID, deleteVolumes, deleteLoadBalancers)
	}

	headers := make(map[string]string)
	if deleteVolumes {
		headers["DeleteVolumes"] = "true"
	}

	if deleteLoadBalancers {
		headers["DeleteLoadBalancers"] = "true"
	}

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath, clusterID)
	_, err := c.DeleteWithHeader(requestURL, headers)
	if err != nil {
		return errors.Wrapf(err, "failed to delete cluster %s project %s datacenter %s", clusterID, projectID, dc)
	}

	return nil
}

// GetKubeConfig gets a clusters Kubeconfig
func (c *Client) GetKubeConfig(clusterID string, projectID string, dc string) (string, error) {
	if clusterID == "" {
		return "", fmt.Errorf("failed to get kubeconfig: no clusterID given")
	}

	if projectID == "" {
		return "", fmt.Errorf("failed to get kubeconfig: no projectID given")
	}

	if dc == "" {
		return "", fmt.Errorf("failed to get kubeconfig: no datacenter given")
	}

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath, clusterID, kubeconfigPath)
	resp, err := c.Get(requestURL, nil)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get kubeconfig for cluster %s project %s datacenter %s", clusterID, projectID, dc)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get kubeconfig for cluster %s project %s datacenter %s", clusterID, projectID, dc)
	}

	return string(body), nil
}

// GetClusterHealth returns the health status of a cluster
func (c *Client) GetClusterHealth(clusterID string, projectID string, dc string) (*models.ClusterHealth, error) {
	result := &models.ClusterHealth{}

	if clusterID == "" {
		return result, fmt.Errorf("failed to get cluster health: no clusterID given")
	}

	if projectID == "" {
		return result, fmt.Errorf("failed to get cluster health: no projectID given")
	}

	if dc == "" {
		return result, fmt.Errorf("failed to get cluster health: no datacenter given")
	}

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s/%s",
		projectPath,
		projectID,
		datacenterPath,
		dc,
		clusterPath,
		clusterID,
		clusterHealthPath,
	)
	_, err := c.Get(requestURL, &result)
	return result, errors.Wrapf(err, "failed to get cluster health for cluster %s in project %s in datacenter %s", clusterID, projectID, dc)
}

// GetClusterUpgradeVersions upgrades a cluster to a specified version
func (c *Client) GetClusterUpgradeVersions(clusterID string, projectID string, dc string) (model.VersionList, error) {
	result := make(model.VersionList, 0)

	if clusterID == "" {
		return result, fmt.Errorf("failed to get upgradeable versions: no clusterID given")
	}

	if projectID == "" {
		return result, fmt.Errorf("failed to get upgradeable versions: no projectID given")
	}

	if dc == "" {
		return result, fmt.Errorf("failed to get upgradeable versions: no datacenter given")
	}

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s/%s",
		projectPath,
		projectID,
		datacenterPath,
		dc,
		clusterPath,
		clusterID,
		upgradesPath,
	)
	_, err := c.Get(requestURL, &result)
	return result, errors.Wrapf(err, "failed to get upgradeable versions for cluster %s project %s datacenter %s", clusterID, projectID, dc)
}

// UpgradeCluster upgrades a cluster to a specified version
func (c *Client) UpgradeCluster(toVersion string, clusterID string, projectID string, dc string) (models.Cluster, error) {
	result := models.Cluster{}

	if clusterID == "" {
		return result, fmt.Errorf("failed to upgrade cluster: no clusterID given")
	}

	if projectID == "" {
		return result, fmt.Errorf("failed to upgrade cluster: no projectID given")
	}

	if dc == "" {
		return result, fmt.Errorf("failed to upgrade cluster: no datacenter given")
	}

	cluster, err := c.GetClusterInProjectInDC(clusterID, projectID, dc)
	if err != nil {
		return result, errors.Wrapf(err, "failed to upgrade cluster %s in project %s in datacenter %s", clusterID, projectID, dc)
	}

	cluster.Spec.Version = toVersion

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath, clusterID)
	_, err = c.Patch(requestURL, contentTypeJSON, cluster, &result)
	return result, errors.Wrapf(err, "failed to upgrade cluster %s in project %s in datacenter %s", clusterID, projectID, dc)
}
