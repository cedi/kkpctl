package describe

import (
	"errors"
	"fmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// Object takes any KKP Object as an input and then describes it
func Object(object interface{}) (string, error) {
	// this is ugly and long, but it makes things kinda nicer to handle outside of the package
	project, ok := object.(models.Project)
	if ok {
		return describeProject(project)
	}

	return fmt.Sprintf("%v\n", object), errors.New("Unable to parse proper type of object")
}
