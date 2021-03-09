package output

import (
	"errors"
	"fmt"

	"github.com/kubermatic/go-kubermatic/models"
)

const (
	// Text specifies the output format as human readable text
	Text string = "text"

	// JSON specifies the output format as JSON
	JSON string = "json"

	// YAML specifies the output format as YAML (Experimental)
	YAML string = "yaml"
)

// ParseOutput takes any KKP Object as an input and then parses it to the appropriate output format
func ParseOutput(object interface{}, output string) (string, error) {
	// this is ugly and long, but it makes things kinda nicer to handle outside of the package
	projects, ok := object.([]models.Project)
	if ok {
		return parseProjects(projects, output)
	}

	project, ok := object.(models.Project)
	if ok {
		return parseProject(project, output)
	}

	return fmt.Sprintf("%v\n", object), errors.New("Unable to parse proper type of object")
}
