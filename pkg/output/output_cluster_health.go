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
	Apiserver                    models.HealthStatus `header:"apiserver"`
	CloudProviderInfrastructure  models.HealthStatus `header:"cloudProviderInfrastructure"`
	Controller                   models.HealthStatus `header:"controller"`
	Etcd                         models.HealthStatus `header:"etcd"`
	MachineController            models.HealthStatus `header:"machineController"`
	Scheduler                    models.HealthStatus `header:"scheduler"`
	UserClusterControllerManager models.HealthStatus `header:"userClusterControllerManager"`
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
			Apiserver:                    object.Apiserver,
			CloudProviderInfrastructure:  object.CloudProviderInfrastructure,
			Controller:                   object.Controller,
			Etcd:                         object.Etcd,
			MachineController:            object.MachineController,
			Scheduler:                    object.Scheduler,
			UserClusterControllerManager: object.UserClusterControllerManager,
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
