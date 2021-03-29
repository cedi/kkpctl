package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/kubermatic/go-kubermatic/models"
	"github.com/lensesio/tableprinter"
	"gopkg.in/yaml.v3"
)

// projectRender is a intermediate struct to make use of lensesio/tableprinter, which relies on the header anotation
type projectRender struct {
	ID                string `header:"ProjectID"`
	Name              string `header:"Name"`
	Owner             string `header:"Owner"`
	CreationTimestamp string `header:"Created"`
	Status            string `header:"Status"`
	Clusters          int64  `header:"Clusters"`
}

func (r projectRender) ParseObject(inputObj interface{}, output string) (string, error) {
	switch object := inputObj.(type) {
	case models.Project:
		return r.ParseCollection([]models.Project{object}, output, Name)

	case *models.Project:
		return r.ParseCollection([]models.Project{*object}, output, Name)

	default:
		return "", fmt.Errorf("inputObj is neighter a models.Project nor a *models.Project")
	}
}

func (r projectRender) ParseCollection(inputObj interface{}, output string, sortBy string) (string, error) {
	var err error
	var parsedOutput []byte

	objects, ok := inputObj.([]models.Project)
	if !ok {
		return "", fmt.Errorf("inputObj is not a []models.Project")
	}

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(objects, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(objects)

	case Text:
		rendered := make([]projectRender, len(objects))
		for idx, object := range objects {
			rendered[idx] = projectRender{
				ID:                object.ID,
				Name:              object.Name,
				CreationTimestamp: object.CreationTimestamp.String(),
				Status:            object.Status,
				Clusters:          object.ClustersNumber,
			}

			owners := make([]string, len(object.Owners))
			for idx, owner := range object.Owners {
				owners[idx] = owner.Name
			}
			rendered[idx].Owner = strings.Join(owners, ", ")
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
