package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"sort"

	"github.com/cedi/kkpctl/pkg/model"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/lensesio/tableprinter"
	"gopkg.in/yaml.v3"
)

// clusterRender is a intermediate struct to make use of lensesio/tableprinter, which relies on the header anotation
type clusterRender struct {
	ID                string `header:"ClusterID"`
	Name              string `header:"Name"`
	Version           string `header:"Version"`
	CreationTimestamp string `header:"Created"`
	Datacenter        string `header:"Datacenter"`
	Provider          string `header:"Provider"`
}

func (r clusterRender) ParseObject(inputObj interface{}, output string) (string, error) {
	switch cluster := inputObj.(type) {
	case models.Cluster:
		return r.ParseCollection([]models.Cluster{cluster}, output, Name)

	case *models.Cluster:
		return r.ParseCollection([]models.Cluster{*cluster}, output, Name)

	default:
		return "", fmt.Errorf("inputObj is neighter a models.Cluster nor a *models.Cluster")
	}
}

func (r clusterRender) ParseCollection(inputObj interface{}, output string, sortBy string) (string, error) {
	var err error
	var parsedOutput []byte

	objects, ok := inputObj.([]models.Cluster)
	if !ok {
		return "", fmt.Errorf("inputObj is not a []models.Cluster")
	}

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(objects, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(objects)

	case Text:
		rendered := make([]clusterRender, len(objects))
		for idx, object := range objects {
			rendered[idx] = clusterRender{
				ID:                object.ID,
				Name:              object.Name,
				CreationTimestamp: object.CreationTimestamp.String(),
				Datacenter:        object.Spec.Cloud.DatacenterName,
				Provider:          model.GetProviderNameFromCloudSpec(object.Spec.Cloud),
			}

			version, ok := object.Status.Version.(string)
			if !ok {
				// Honestly, I don't know what to do here^^
				continue
			}

			rendered[idx].Version = version
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
		return "", fmt.Errorf("unable to parse objects")
	}

	return string(parsedOutput), err
}

func init() {
	parser := GetParserFactory()
	parser.AddCollectionParser(reflect.TypeOf([]models.Cluster{}), clusterRender{})
	parser.AddObjectParser(reflect.TypeOf(models.Cluster{}), clusterRender{})
	parser.AddObjectParser(reflect.TypeOf(&models.Cluster{}), clusterRender{})
}
