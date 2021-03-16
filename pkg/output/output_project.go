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

func parseProject(object models.Project, output string) (string, error) {
	return parseProjects([]models.Project{object}, output, Name)
}

func parseProjects(objects []models.Project, output string, sortBy string) (string, error) {
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
		return "", fmt.Errorf("unable to parse objects")
	}
}
