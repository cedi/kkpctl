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

// ListClustersForProjectAndDatacenter lists all clusters for a given Project
// (identified by ID) and Seed (identified by name)
//func (c *Client) ListClustersForProjectAndDatacenter(projectID string, seed string) ([]models.Cluster, error) {
//	req, err := c.newRequest("GET", projectPath+"/"+projectID+datacenterSubPath+"/"+seed+clustersSubPath, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	result := make([]models.Cluster, 0)
//
//	resp, err := c.do(req, &result)
//	if err != nil {
//		return nil, err
//	}
//
//	// StatusCodes 401 and 403 mean empty response and should be treated as such
//	if resp.StatusCode == 401 || resp.StatusCode == 403 {
//		return nil, nil
//	}
//
//	if resp.StatusCode >= 299 {
//		return nil, errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
//	}
//
//	return result, nil
//}

// GetCluster fetches cluster
//func (c *Client) GetCluster(projectID string, seed string, clusterID string) (models.Cluster, error) {
//	req, err := c.newRequest("GET", projectPath+"/"+projectID+datacenterSubPath+"/"+seed+clustersSubPath+"/"+clusterID, nil)
//	if err != nil {
//		return models.Cluster{}, err
//	}
//
//	result := models.Cluster{}
//
//	resp, err := c.do(req, &result)
//	if err != nil {
//		return models.Cluster{}, err
//	}
//
//	// StatusCodes 401 and 403 mean empty response and should be treated as such
//	if resp.StatusCode == 401 || resp.StatusCode == 403 {
//		return models.Cluster{}, nil
//	}
//
//	if resp.StatusCode >= 299 {
//		return models.Cluster{}, errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
//	}
//
//	return result, nil
//}

// DeleteCluster deletes a given cluster
//func (c *Client) DeleteCluster(projectID string, seed string, clusterID string) error {
//	req, err := c.newRequest("DELETE", projectPath+"/"+projectID+datacenterSubPath+"/"+seed+clustersSubPath+"/"+clusterID, nil)
//	if err != nil {
//		return err
//	}
//
//	resp, err := c.do(req, nil)
//	if err != nil {
//		return err
//	}
//
//	// StatusCodes 401 and 403 mean empty response and should be treated as such
//	if resp.StatusCode == 401 || resp.StatusCode == 403 {
//		return nil
//	}
//
//	if resp.StatusCode >= 299 {
//		return errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
//	}
//
//	return nil
//}

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
