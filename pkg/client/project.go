package client

import (
	"fmt"

	"github.com/kubermatic/go-kubermatic/models"
)

const (
	projectPath string = "/projects"
)

// ListProjects lists all projects a user has permission to see
// if all is true, all projects the user has access to, will be listed
// if all is false (default), only clusters owned by the user will be listed
func (c *Client) ListProjects(all bool) ([]models.Project, error) {
	var err error
	result := make([]models.Project, 0)

	if all {
		_, err = c.GetWithQueryParams(projectPath, URLParams{"displayAll": "true"}, &result)
	} else {
		_, err = c.GetWithQueryParams(projectPath, URLParams{"displayAll": "false"}, &result)
	}

	if err != nil {
		return result, err
	}

	return result, nil
}

// GetProject gets a specific project
func (c *Client) GetProject(projectID string) (models.Project, error) {
	result := models.Project{}
	_, err := c.Get(fmt.Sprintf("%s/%s", projectPath, projectID), &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// CreateProject creates a new project with the given name and labels
func (c *Client) CreateProject(name string, labels map[string]string) (models.Project, error) {
	newProject := models.Project{
		Name:   name,
		Labels: labels,
	}

	result := models.Project{}
	_, err := c.Post(projectPath, contentTypeJSON, newProject, &result)
	if err != nil {
		return models.Project{}, err
	}

	return result, nil
}

// DeleteProject deletes a project identified by id
func (c *Client) DeleteProject(projectID string) error {
	_, err := c.Delete(fmt.Sprintf("%s/%s", projectPath, projectID))
	if err != nil {
		return err
	}

	return nil
}

// GetProjectIDForCluster deletes a project identified by id
func (c *Client) GetProjectIDForCluster(clusterID string) (string, error) {
	projects, err := c.ListProjects(true)
	if err != nil {
		return "", err
	}

	for _, project := range projects {
		clusters, err := c.ListClustersInProject(project.ID)
		if err != nil {
			return "", err
		}

		for _, cluster := range clusters {
			if cluster.ID == clusterID {
				return project.ID, nil
			}
		}
	}

	return "", fmt.Errorf("no Project for ClusterID %s found", clusterID)
}
