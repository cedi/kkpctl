package client

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kubermatic/go-kubermatic/models"
)

const (
	projectPath string = "/projects"
)

// ListProjects lists all projects a user has permission to see
// if all is true, all projects the user has access to, will be listed
// if all is false (default), only clusters owned by the user will be listed
func (c *Client) ListProjects(all bool) ([]models.Project, error) {
	var resp *http.Response
	var err error
	result := make([]models.Project, 0)

	if all {
		resp, err = c.GetWithQueryParams(projectPath, URLParams{"displayAll": "true"}, &result)
	} else {
		resp, err = c.GetWithQueryParams(projectPath, URLParams{"displayAll": "false"}, &result)
	}

	if err != nil {
		return result, err
	}

	// Something non-2xx - not good...
	if resp.StatusCode >= 299 {
		return nil, errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
	}

	return result, nil
}

// GetProject gets a specific project
func (c *Client) GetProject(projectID string) (models.Project, error) {
	result := models.Project{}
	resp, err := c.Get(fmt.Sprintf("%s/%s", projectPath, projectID), &result)
	if err != nil {
		return result, err
	}

	// StatusCodes 401 and 409 mean empty response and should be treated as such
	if resp.StatusCode == 401 || resp.StatusCode == 409 {
		return result, nil
	}

	// Something non-2xx - not good...
	if resp.StatusCode >= 299 {
		return result, errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
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
	resp, err := c.Post(projectPath, contentTypeJSON, newProject, &result)
	if err != nil {
		return models.Project{}, err
	}

	// StatusCodes 401 and 409 mean empty response and should be treated as such
	if resp.StatusCode == 401 || resp.StatusCode == 409 {
		return models.Project{}, nil
	}

	// Something non-2xx - not good...
	if resp.StatusCode >= 299 {
		return models.Project{}, errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
	}

	return result, nil
}

// DeleteProject deletes a project identified by id
func (c *Client) DeleteProject(projectID string) error {
	resp, err := c.Delete(fmt.Sprintf("%s/%s", projectPath, projectID))
	if err != nil {
		return err
	}

	// StatusCodes 401 and 403 mean empty response and should be treated as such
	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return nil
	}

	// Something non-2xx - not good...
	if resp.StatusCode >= 299 {
		return errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}
