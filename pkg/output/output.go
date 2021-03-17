package output

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/config"
	"github.com/cedi/kkpctl/pkg/model"
	"github.com/cedi/kkpctl/pkg/utils"
	"github.com/kubermatic/go-kubermatic/models"
)

const (
	// Text as output parameter specifies the output format as human readable text
	Text string = "text"

	// JSON as output parameter specifies the output format as JSON
	JSON string = "json"

	// YAML as output parameter specifies the output format as YAML (Experimental)
	YAML string = "yaml"

	// Name as sortBy parameter specified the output (if in a list) should be sorted by name
	Name string = "name"

	// Date as sortBy parameter specified the output (if in a list) should be sorted by Date
	Date string = "date"
)

// ParseOutput takes any KKP Object as an input and then parses it to the appropriate output format
func ParseOutput(object interface{}, output string, sortBy string) (string, error) {

	err := validateOutput(output)
	if err != nil {
		return "", err
	}

	err = validateSorting(sortBy)
	if err != nil {
		return "", err
	}

	return parseOutput(object, output, sortBy)
}

// parseOutput is ugly and long, but it makes things kinda nicer to handle outside of the package
func parseOutput(outputObject interface{}, output string, sortBy string) (string, error) {

	switch o := outputObject.(type) {
	// KKP Project
	case []models.Project:
		return parseProjects(o, output, sortBy)
	case models.Project:
		return parseProject(o, output)

	// KKP Clusters
	case []models.Cluster:
		return parseClusters(o, output, sortBy)
	case models.Cluster:
		return parseCluster(o, output)
	case *models.Cluster:
		return parseCluster(*o, output)

	// Node Deployments
	case []models.NodeDeployment:
		return parseNodeDeployments(o, output, sortBy)
	case models.NodeDeployment:
		return parseNodeDeployment(o, output)
	case *models.NodeDeployment:
		return parseNodeDeployment(*o, output)

	// Datacenter
	case []models.Datacenter:
		return parseDatacenters(o, output, sortBy)
	case models.Datacenter:
		return parseDatacenter(o, output)

	// ClusterHealth
	case *models.ClusterHealth:
		return parseClusterHealth(o, output)

	// Cluster Versions
	case model.VersionList:
		return parseClusterVersions(o, output)
	case model.Version:
		return parseClusterVersion(o, output)

	// Config
	case config.CloudConfig:
		return parseConfigCloud(o, output)

	// Events
	case []models.Event:
		return parseEvents(o, output)

	// Node Taints
	case *models.TaintSpec:
		return parseNodeTaint(o, output)
	case []*models.TaintSpec:
		return parseNodeTaints(o, output)

	// Node
	case models.Node:
		return parseNode(o, output)
	case []models.Node:
		return parseNodes(o, output)
	}

	return fmt.Sprintf("%v\n", outputObject), fmt.Errorf("unable to parse proper type of object")
}

func validateOutput(output string) error {
	if !utils.IsOneOf(output, Text, JSON, YAML) {
		return fmt.Errorf("the output type '%s' is not a valid output", output)
	}
	return nil
}

func validateSorting(sort string) error {
	if !utils.IsOneOf(sort, Name, Date) {
		return fmt.Errorf("the sort parameter '%s' is not a valid sorting criteria", sort)
	}

	return nil
}
