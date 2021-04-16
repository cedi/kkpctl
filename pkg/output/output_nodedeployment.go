package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"sort"

	"github.com/kubermatic/go-kubermatic/models"
	"github.com/lensesio/tableprinter"
	"gopkg.in/yaml.v3"
)

// nodeDeploymentRender is a intermediate struct to make use of lensesio/tableprinter, which relies on the header anotation
type nodeDeploymentRender struct {
	ID                string `header:"NodeDeploymentID"`
	Name              string `header:"Name"`
	Version           string `header:"Version"`
	Replicas          int32  `header:"Replicas"`
	ReadyReplicas     int32  `header:"Replicas Ready"`
	Paused            bool   `header:"Paused"`
	OperatingSystem   string `header:"OperatingSystem"`
	CreationTimestamp string `header:"Created"`
}

func (r nodeDeploymentRender) ParseObject(inputObj interface{}, output string) (string, error) {
	switch object := inputObj.(type) {
	case models.NodeDeployment:
		return r.ParseCollection([]models.NodeDeployment{object}, output, Name)

	case *models.NodeDeployment:
		return r.ParseCollection([]models.NodeDeployment{*object}, output, Name)

	default:
		return "", fmt.Errorf("inputObj is neighter a models.NodeDeployment nor a *models.NodeDeployment")
	}
}

func (r nodeDeploymentRender) ParseCollection(inputObj interface{}, output string, sortBy string) (string, error) {
	var err error
	var parsedOutput []byte

	objects, ok := inputObj.([]models.NodeDeployment)
	if !ok {
		return "", fmt.Errorf("inputObj is not a []models.NodeDeployment")
	}

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(objects, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(objects)

	case Text:
		rendered := make([]nodeDeploymentRender, len(objects))
		for idx, object := range objects {
			rendered[idx] = nodeDeploymentRender{
				ID:                object.ID,
				Name:              object.Name,
				CreationTimestamp: object.CreationTimestamp.String(),
				Version:           object.Spec.Template.Versions.Kubelet,
				Replicas:          *object.Spec.Replicas,
				ReadyReplicas:     object.Status.AvailableReplicas,
				Paused:            object.Spec.Paused,
				OperatingSystem:   getOperatingSystem(object.Spec.Template.OperatingSystem),
			}
		}

		sort.Slice(rendered, func(i, j int) bool {
			if sortBy == Date {
				return rendered[j].CreationTimestamp < rendered[i].CreationTimestamp
			}

			return rendered[j].Name > rendered[i].Name
		})

		var bodyBuf io.ReadWriter
		bodyBuf = new(bytes.Buffer)

		tableprinter.Print(bodyBuf, rendered)
		parsedOutput, err = ioutil.ReadAll(bodyBuf)

	default:
		return "", fmt.Errorf("unable to parse node deployment")
	}

	return string(parsedOutput), err
}

func getOperatingSystem(osSpec *models.OperatingSystemSpec) string {
	if osSpec.Centos != nil {
		return "Centos"
	}
	if osSpec.Flatcar != nil {
		return "Flatcar"
	}
	if osSpec.Rhel != nil {
		return "RHEL"
	}
	if osSpec.Sles != nil {
		return "SLES"
	}
	if osSpec.Ubuntu != nil {
		return "Ubuntu"
	}
	return ""
}

func init() {
	parser := GetParserFactory()
	parser.AddCollectionParser(reflect.TypeOf([]models.NodeDeployment{}), nodeDeploymentRender{})
	parser.AddObjectParser(reflect.TypeOf(models.NodeDeployment{}), nodeDeploymentRender{})
	parser.AddObjectParser(reflect.TypeOf(&models.NodeDeployment{}), nodeDeploymentRender{})
}
