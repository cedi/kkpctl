package output

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"sort"

	"github.com/cedi/kkpctl/pkg/model"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/lensesio/tableprinter"
)

type clusterProjectRender struct {
	ProjectID         string `header:"ProjectID"`
	ClusterID         string `header:"ClusterID"`
	Name              string `header:"Name"`
	Version           string `header:"Version"`
	CreationTimestamp string `header:"Created"`
	Datacenter        string `header:"Datacenter"`
	Provider          string `header:"Provider"`
}

func (r clusterProjectRender) ParseObject(inputObj interface{}, output string) (string, error) {
	switch cluster := inputObj.(type) {
	case model.ProjectCluster:
		return r.ParseCollection([]model.ProjectCluster{cluster}, output, Name)

	case *model.ProjectCluster:
		return r.ParseCollection([]model.ProjectCluster{*cluster}, output, Name)

	default:
		return "", fmt.Errorf("inputObj is neighter a ClusterProjectRender nor a *ClusterProjectRender")
	}
}

func (r clusterProjectRender) ParseCollection(inputObj interface{}, output string, sortBy string) (string, error) {
	var err error
	var parsedOutput []byte

	objects, ok := inputObj.([]model.ProjectCluster)
	if !ok {
		return "", fmt.Errorf("inputObj is not a []ClusterProjectRender")
	}

	switch output {
	case Text:
		rendered := make([]clusterProjectRender, len(objects))
		for idx, object := range objects {
			rendered[idx] = clusterProjectRender{
				ProjectID:         object.ProjectID,
				ClusterID:         object.Cluster.ID,
				Name:              object.Cluster.Name,
				CreationTimestamp: object.Cluster.CreationTimestamp.String(),
				Datacenter:        object.Cluster.Spec.Cloud.DatacenterName,
				Provider:          model.GetProviderNameFromCloudSpec(object.Cluster.Spec.Cloud),
			}

			version, ok := object.Cluster.Status.Version.(string)
			if !ok {
				// Honestly, I don't know what to do here^^
				continue
			}

			rendered[idx].Version = version
		}

		sort.Slice(rendered, func(i, j int) bool {
			if sortBy == Date {
				return rendered[j].CreationTimestamp > rendered[i].CreationTimestamp
			}

			return rendered[j].Name > rendered[i].Name
		})

		var bodyBuf io.ReadWriter
		bodyBuf = new(bytes.Buffer)

		tableprinter.Print(bodyBuf, rendered)
		parsedOutput, err = ioutil.ReadAll(bodyBuf)

	default:
		clusters := make([]models.Cluster, 0)
		for _, renderObj := range objects {
			clusters = append(clusters, renderObj.Cluster)
		}

		return ParseOutput(clusters, output, sortBy)
	}

	return string(parsedOutput), err
}

func init() {
	parser := GetParserFactory()
	parser.AddCollectionParser(reflect.TypeOf([]model.ProjectCluster{}), clusterProjectRender{})
	parser.AddObjectParser(reflect.TypeOf(model.ProjectCluster{}), clusterProjectRender{})
	parser.AddObjectParser(reflect.TypeOf(&model.ProjectCluster{}), clusterProjectRender{})
}
