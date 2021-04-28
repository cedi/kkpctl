package describe

import (
	"fmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// Object takes any KKP Object as an input and then describes it
func Object(object interface{}) (string, error) {
	switch describeObj := object.(type) {
	case *models.Project:
		return describeProject(describeObj)

	case *ClusterDescribeMeta:
		return describeCluster(describeObj)

	case *MachineDeploymentDescribeMeta:
		return describeNodeDeployment(describeObj)

	default:
		return fmt.Sprintf("%v\n", object), fmt.Errorf("unable to parse proper type of object")
	}
}
