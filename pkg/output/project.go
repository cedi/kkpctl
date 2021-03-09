package output

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"sort"

	"github.com/kubermatic/go-kubermatic/models"
	"github.com/lensesio/tableprinter"
	"gopkg.in/yaml.v3"
)

// projectRender is a intermediate struct to make use of lensesio/tableprinter, which relies on the header anotation
type projectRender struct {
	ID                string `header:"ProjectID"`
	Name              string `header:"Name"`
	CreationTimestamp string `header:"Created"`
	Status            string `header:"Status"`
	Clusters          int64  `header:"Clusters"`
}

func parseProjects(objects []models.Project, output string) (string, error) {
	switch output {
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
		}

		sort.Slice(rendered, func(i, j int) bool {
			return rendered[j].CreationTimestamp < rendered[i].CreationTimestamp
		})

		var bodyBuf io.ReadWriter
		bodyBuf = new(bytes.Buffer)

		tableprinter.Print(bodyBuf, rendered)
		bodyBytes, err := ioutil.ReadAll(bodyBuf)
		return string(bodyBytes), err

	case JSON:
		buf, err := json.MarshalIndent(objects, "", "  ")
		if err != nil {
			return "", err
		}

		return string(buf), nil

	case YAML:
		buf, err := yaml.Marshal(objects)
		return "---\n" + string(buf), err

	default:
		return "", errors.New("Unable to parse objects")
	}
}

func parseProject(object models.Project, output string) (string, error) {
	switch output {
	case Text:
		rendered := projectRender{
			ID:                object.ID,
			Name:              object.Name,
			CreationTimestamp: object.CreationTimestamp.String(),
			Status:            object.Status,
			Clusters:          object.ClustersNumber,
		}

		var bodyBuf io.ReadWriter
		bodyBuf = new(bytes.Buffer)

		tableprinter.Print(bodyBuf, rendered)
		bodyBytes, err := ioutil.ReadAll(bodyBuf)
		return string(bodyBytes), err

	case JSON:
		buf, err := json.MarshalIndent(object, "", "  ")
		if err != nil {
			return "", err
		}

		return string(buf), nil

	case YAML:
		buf, err := yaml.Marshal(object)
		return string(buf), err

	default:
		return "", errors.New("Unable to parse object")
	}
}
