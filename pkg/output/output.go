package output

import (
	"fmt"
	"reflect"

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

// make the parser factory a singleton
var parser *ParserFactory

func init() {
	parser = NewParserFactory()

	// KKP Project
	parser.AddCollectionParser(reflect.TypeOf([]models.Project{}), projectRender{})
	parser.AddObjectParser(reflect.TypeOf(models.Project{}), projectRender{})
	parser.AddObjectParser(reflect.TypeOf(&models.Project{}), projectRender{})

	// Datacenter
	parser.AddCollectionParser(reflect.TypeOf([]models.Datacenter{}), datacenterRender{})
	parser.AddObjectParser(reflect.TypeOf(models.Datacenter{}), datacenterRender{})
	parser.AddObjectParser(reflect.TypeOf(&models.Datacenter{}), datacenterRender{})

	// KKP Clusters
	parser.AddCollectionParser(reflect.TypeOf([]models.Cluster{}), clusterRender{})
	parser.AddObjectParser(reflect.TypeOf(models.Cluster{}), clusterRender{})
	parser.AddObjectParser(reflect.TypeOf(&models.Cluster{}), clusterRender{})

	// Cluster Versions
	parser.AddCollectionParser(reflect.TypeOf(model.VersionList{}), clusterVersionRender{})
	parser.AddObjectParser(reflect.TypeOf(model.Version{}), clusterVersionRender{})
	parser.AddObjectParser(reflect.TypeOf(&model.Version{}), clusterVersionRender{})

	// ClusterHealth
	parser.AddObjectParser(reflect.TypeOf(&models.ClusterHealth{}), clusterHealthRender{})

	// Events
	parser.AddCollectionParser(reflect.TypeOf([]models.Event{}), eventRender{})

	// Node Deployments
	parser.AddCollectionParser(reflect.TypeOf([]models.NodeDeployment{}), nodeDeploymentRender{})
	parser.AddObjectParser(reflect.TypeOf(models.NodeDeployment{}), nodeDeploymentRender{})
	parser.AddObjectParser(reflect.TypeOf(&models.NodeDeployment{}), nodeDeploymentRender{})

	// Node
	parser.AddCollectionParser(reflect.TypeOf([]models.Node{}), nodeRender{})
	parser.AddObjectParser(reflect.TypeOf(models.Node{}), nodeRender{})
	parser.AddObjectParser(reflect.TypeOf(&models.Node{}), nodeRender{})

	// Node Taints
	parser.AddCollectionParser(reflect.TypeOf([]*models.TaintSpec{}), taintRender{})
	parser.AddObjectParser(reflect.TypeOf(&models.TaintSpec{}), taintRender{})

	// Config
	parser.AddObjectParser(reflect.TypeOf(config.CloudConfig{}), configCloudRender{})
	parser.AddObjectParser(reflect.TypeOf(&config.CloudConfig{}), configCloudRender{})
}

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

	collectionsParser, ok := parser.GetCollectionParser(reflect.TypeOf(object))
	if ok {
		return collectionsParser.ParseCollection(object, output, sortBy)
	}

	objectParser, ok := parser.GetObjectParser(reflect.TypeOf(object))
	if ok {
		return objectParser.ParseObject(object, output)
	}

	return fmt.Sprintf("%v\n", object), fmt.Errorf("unable to determine proper output type")
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
