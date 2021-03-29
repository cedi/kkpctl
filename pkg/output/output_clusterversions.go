package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"sort"

	"github.com/cedi/kkpctl/pkg/model"
	"github.com/lensesio/tableprinter"
	"gopkg.in/yaml.v3"
)

// clusterRender is a intermediate struct to make use of lensesio/tableprinter, which relies on the header anotation
type clusterVersionRender struct {
	Version string `header:"Version"`
}

func (r clusterVersionRender) ParseObject(inputObj interface{}, output string) (string, error) {

	switch clusterVersion := inputObj.(type) {
	case model.Version:
		return r.ParseCollection(model.VersionList{clusterVersion}, output, Name)

	case *model.Version:
		return r.ParseCollection(model.VersionList{*clusterVersion}, output, Name)

	default:
		return "", fmt.Errorf("inputObj is neighter a models.Version nor a *models.Version")
	}
}

func (r clusterVersionRender) ParseCollection(inputObj interface{}, output string, sortBy string) (string, error) {
	var err error
	parsedOutput := make([]byte, 0)

	objects, ok := inputObj.(model.VersionList)
	if !ok {
		return "", fmt.Errorf("inputObj is not a model.VersionList")
	}

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(objects, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(objects)

	case Text:
		// Sort first, so the highes K8s version is returned first
		sort.Sort(objects)

		rendered := make([]clusterVersionRender, 0)
		for idx, version := range objects {
			rendered = append(rendered, clusterVersionRender{
				Version: version.Version,
			})

			if version.Default {
				rendered[idx].Version = fmt.Sprintf("%s *", rendered[idx].Version)
			}
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
