package output

import (
	"fmt"

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
func parseOutput(object interface{}, output string, sortBy string) (string, error) {
	// KKP Projects
	projects, ok := object.([]models.Project)
	if ok {
		return parseProjects(projects, output, sortBy)
	}

	project, ok := object.(models.Project)
	if ok {
		return parseProject(project, output)
	}

	// KKP Clusters
	clusters, ok := object.([]models.Cluster)
	if ok {
		return parseClusters(clusters, output, sortBy)
	}

	cluster, ok := object.(models.Cluster)
	if ok {
		return parseCluster(cluster, output)
	}

	clusterP, ok := object.(*models.Cluster)
	if ok {
		return parseCluster(*clusterP, output)
	}

	// Node Deployments
	nodeDeployments, ok := object.([]models.NodeDeployment)
	if ok {
		return parseNodeDeployments(nodeDeployments, output, sortBy)
	}

	nodeDeployment, ok := object.(models.NodeDeployment)
	if ok {
		return parseNodeDeployment(nodeDeployment, output)
	}

	// Datacenter
	datacenters, ok := object.([]models.Datacenter)
	if ok {
		return parseDatacenters(datacenters, output, sortBy)
	}

	datacenter, ok := object.(models.Datacenter)
	if ok {
		return parseDatacenter(datacenter, output)
	}

	return fmt.Sprintf("%v\n", object), fmt.Errorf("unable to parse proper type of object")
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
