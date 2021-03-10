package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/kubermatic/go-kubermatic/models"
)

const (
	clusterPath    string = "clusters"
	datacenterPath string = "dc"
	kubeconfigPath string = "kubeconfig"
)

// ListClusters lists all clusters for a given Project (identified by ID)
func (c *Client) ListClusters(projectID string) ([]models.Cluster, error) {
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

// ListClustersInDC lists all clusters for a given Project in a given datacenter
func (c *Client) ListClustersInDC(projectID string, dc string) ([]models.Cluster, error) {
	result := make([]models.Cluster, 0)

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath)
	_, err := c.Get(requestURL, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetCluster gets a clusters in a given Project in a given datacenter
func (c *Client) GetCluster(clusterID string, projectID string, dc string) (models.Cluster, error) {
	result := models.Cluster{}

	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath, clusterID)
	_, err := c.Get(requestURL, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// DeleteCluster deletes a cluster identified by id
func (c *Client) DeleteCluster(clusterID string, projectID string, dc string) error {
	requestURL := fmt.Sprintf("%s/%s/%s/seed-%s/%s/%s", projectPath, projectID, datacenterPath, dc, clusterPath, clusterID)
	_, err := c.Delete(requestURL)
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
