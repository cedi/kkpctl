package client

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kubermatic/go-kubermatic/models"
)

const (
	clusterPath    string = "clusters"
	datacenterPath string = "dc"
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

// GetClusterKubeconfig fetches kubeconfig for a given cluster
//func (c *Client) GetClusterKubeconfig(projectID string, seed string, clusterID string) (clientcmdapi.Config, error) {
//	req, err := c.newRequest("GET", projectPath+"/"+projectID+datacenterSubPath+"/"+seed+clustersSubPath+"/"+clusterID+kubeconfigSubPath, nil)
//	if err != nil {
//		return clientcmdapi.Config{}, err
//	}
//
//	result := clientcmdapi.Config{}
//
//	resp, err := c.do(req, &result)
//	if err != nil {
//		return clientcmdapi.Config{}, err
//	}
//
//	// StatusCodes 401 and 403 mean empty response and should be treated as such
//	if resp.StatusCode == 401 || resp.StatusCode == 403 {
//		return clientcmdapi.Config{}, nil
//	}
//
//	if resp.StatusCode >= 299 {
//		return clientcmdapi.Config{}, errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
//	}
//
//	return result, nil
//}
