package client

import (
	"fmt"

	"github.com/kubermatic/go-kubermatic/models"
)

const eventsPath string = "events"

// GetEvents returns the events of a cluster
func (c *Client) GetEvents(clusterID string, projectID string) ([]models.Event, error) {
	var err error
	result := make([]models.Event, 0)

	requestURL := fmt.Sprintf("%s/%s/%s/%s/%s",
		projectPath,
		projectID,
		clusterPath,
		clusterID,
		eventsPath,
	)
	_, err = c.Get(requestURL, &result, V2API)
	return result, err
}
