package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/kubermatic/go-kubermatic/models"
	"github.com/pkg/errors"
)

const (
	clusterPath       string = "clusters"
	datacenterPath    string = "dc"
	kubeconfigPath    string = "kubeconfig"
	clusterHealthPath string = "health"
)

// ListClusters lists all clusters
//	all lists all clusters in all projects
func (c *Client) ListClusters(all bool) ([]models.Cluster, error) {
	result := make([]models.Cluster, 0)

	projects, err := c.ListProjects(all)
	if err != nil {
		return result, err
	}

	for _, project := range projects {
		clusters, err := c.ListClustersInProject(project.ID)
		if err != nil {
			return result, err
		}

		result = append(result, clusters...)
	}

	return result, nil
}

// ListClustersInDC lists all clusters for a given Project in a given datacenter
//	all lists all clusters in all projects
func (c *Client) ListClustersInDC(dc string, all bool) ([]models.Cluster, error) {
	result := make([]models.Cluster, 0)

	projects, err := c.ListProjects(all)
	if err != nil {
		return result, err
	}
	for _, project := range projects {
		clusters, err := c.ListClustersInProjectInDC(project.ID, dc)
		if err != nil {
			return result, err
		}

		result = append(result, clusters...)
	}

	return result, nil
}

// ListClustersInProject lists all clusters for a given Project (identified by ID)
func (c *Client) ListClustersInProject(projectID string) ([]models.Cluster, error) {
	var resp *http.Response
	var err error
	result := make([]models.Cluster, 0)

	requestURL := fmt.Sprintf("%s/%s/%s", projectPath, projectID, clusterPath)
	resp, err = c.Get(requestURL, &result)

	if err != nil {
		return result, err
	}

	// Something non-2xx - not good...
	if resp.StatusCode >= 299 {
		return nil, errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
	}

	return result, nil
}

// ListClustersInProjectInDC lists all clusters for a given Project in a given datacenter
func (c *Client) ListClustersInProjectInDC(projectID string, dc string) ([]models.Cluster, error) {
	result := make([]models.Cluster, 0)

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath)
	_, err := c.Get(requestURL, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetCluster gets a clusters in a given Project
func (c *Client) GetCluster(clusterID string, listAll bool) (models.Cluster, error) {
	var result *models.Cluster
	result = nil

	clusters, err := c.ListClusters(listAll)
	if err != nil {
		return *result, errors.Wrap(err, "Failed to determine the correct datacenter for a cluster in a given project")
	}

	for _, cluster := range clusters {
		if cluster.ID == clusterID {
			result = &cluster
			break
		}
	}

	if result == nil {
		return *result, errors.Wrap(err, "Failed to find a cluster with this clusterID")
	}

	return *result, nil
}

// GetClusterInDC gets a clusters in a given Project
func (c *Client) GetClusterInDC(clusterID string, datacenter string, listAll bool) (models.Cluster, error) {
	var result *models.Cluster
	result = nil

	clusters, err := c.ListClustersInDC(datacenter, listAll)
	if err != nil {
		return *result, errors.Wrap(err, "Failed to determine the correct datacenter for a cluster in a given project")
	}

	for _, cluster := range clusters {
		if cluster.ID == clusterID && cluster.Spec.Cloud.DatacenterName == datacenter {
			result = &cluster
			break
		}
	}

	if result == nil {
		return *result, errors.Wrap(err, "Failed to find a project with this clusterID")
	}

	return *result, nil
}

// GetClusterInProject gets a clusters in a given Project
func (c *Client) GetClusterInProject(clusterID string, projectID string) (models.Cluster, error) {
	result := models.Cluster{}

	clusters, err := c.ListClustersInProject(projectID)
	if err != nil {
		return result, errors.Wrap(err, "Failed to determine the correct datacenter for a cluster in a given project")
	}

	datacenter := ""
	for _, cluster := range clusters {
		if cluster.ID == clusterID {
			datacenter = cluster.Spec.Cloud.DatacenterName
		}
	}

	if datacenter == "" {
		return result, errors.Wrap(err, "Failed to determine the correct datacenter for a cluster in a given project")
	}

	return c.GetClusterInProjectInDC(clusterID, projectID, datacenter)
}

// GetClusterInProjectInDC gets a clusters in a given Project in a given datacenter
func (c *Client) GetClusterInProjectInDC(clusterID string, projectID string, dc string) (models.Cluster, error) {
	result := models.Cluster{}

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath, clusterID)
	_, err := c.Get(requestURL, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// CreateCluster creates a new cluster
func (c *Client) CreateCluster(newCluster *models.CreateClusterSpec, projectID string, dc string) (models.Cluster, error) {
	result := models.Cluster{}

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath)
	_, err := c.Post(requestURL, contentTypeJSON, newCluster, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// DeleteCluster deletes a cluster identified by id
func (c *Client) DeleteCluster(clusterID string, projectID string, deleteVolumes bool, deleteLoadBalancers bool) error {
	clusters, err := c.ListClustersInProject(projectID)
	if err != nil {
		return errors.Wrap(err, "Failed to determine the correct datacenter for a cluster in a given project")
	}

	datacenter := ""
	for _, cluster := range clusters {
		if cluster.ID == clusterID {
			datacenter = cluster.Spec.Cloud.DatacenterName
		}
	}

	if datacenter == "" {
		return errors.Wrap(err, "Failed to determine the correct datacenter for a cluster in a given project")
	}

	return c.DeleteClusterInDC(clusterID, projectID, datacenter, deleteVolumes, deleteLoadBalancers)
}

// DeleteClusterInDC deletes a cluster identified by id
func (c *Client) DeleteClusterInDC(clusterID string, projectID string, dc string, deleteVolumes bool, deleteLoadBalancers bool) error {
	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath, clusterID)

	headers := make(map[string]string)
	if deleteVolumes {
		headers["DeleteVolumes"] = "true"
	}

	if deleteLoadBalancers {
		headers["DeleteLoadBalancers"] = "true"
	}

	_, err := c.DeleteWithHeader(requestURL, headers)
	if err != nil {
		return err
	}

	return nil
}

// GetKubeConfig gets a clusters Kubeconfig
func (c *Client) GetKubeConfig(clusterID string, projectID string, dc string) (string, error) {
	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath, clusterID, kubeconfigPath)
	resp, err := c.Get(requestURL, nil)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// GetClusterHealth returns the health status of a cluster
func (c *Client) GetClusterHealth(clusterID string, projectID string, dc string) (*models.ClusterHealth, error) {
	var err error
	result := &models.ClusterHealth{}

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s/%s",
		projectPath,
		projectID,
		datacenterPath,
		dc,
		clusterPath,
		clusterID,
		clusterHealthPath,
	)
	_, err = c.Get(requestURL, &result)
	return result, err
}
