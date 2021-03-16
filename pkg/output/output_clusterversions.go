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

func parseClusterVersion(object model.Version, output string) (string, error) {
	return parseClusterVersions(model.VersionList{object}, output)
}

func parseClusterVersions(objects model.VersionList, output string) (string, error) {
	var err error
	var parsedOutput []byte

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
