package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/kubermatic/go-kubermatic/models"
	"github.com/lensesio/tableprinter"
	"gopkg.in/yaml.v3"
)

// clusterRender is a intermediate struct to make use of lensesio/tableprinter, which relies on the header anotation
type clusterHealthRender struct {
	Apiserver                    string `header:"apiserver"`
	CloudProviderInfrastructure  string `header:"cloudProviderInfrastructure"`
	Controller                   string `header:"controller"`
	Etcd                         string `header:"etcd"`
	MachineController            string `header:"machineController"`
	Scheduler                    string `header:"scheduler"`
	UserClusterControllerManager string `header:"userClusterControllerManager"`
}

func parseClusterHealth(object *models.ClusterHealth, output string) (string, error) {
	var err error
	var parsedOutput []byte

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(object, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(object)

	case Text:
		rendered := clusterHealthRender{
			Apiserver:                    getHealthStatus(object.Apiserver),
			CloudProviderInfrastructure:  getHealthStatus(object.CloudProviderInfrastructure),
			Controller:                   getHealthStatus(object.Controller),
			Etcd:                         getHealthStatus(object.Etcd),
			MachineController:            getHealthStatus(object.MachineController),
			Scheduler:                    getHealthStatus(object.Scheduler),
			UserClusterControllerManager: getHealthStatus(object.UserClusterControllerManager),
		}

		var bodyBuf io.ReadWriter
		bodyBuf = new(bytes.Buffer)

		tableprinter.Print(bodyBuf, rendered)
		parsedOutput, err = ioutil.ReadAll(bodyBuf)

	default:
		return "", fmt.Errorf("unable to parse objects")
	}

	return string(parsedOutput), err
}

func getHealthStatus(status models.HealthStatus) string {
	switch status {
	case 0: //kubermaticv1.HealthStatusDown:
		return "Down"

	case 1: //kubermaticv1.HealthStatusUp:
		return "Up"

	case 2: //kubermaticv1.HealthStatusProvisioning:
		return "Provisioning"
	}
	return "Unknown"
}
