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

// datacenterRender is a intermediate struct to make use of lensesio/tableprinter, which relies on the header anotation
type datacenterRender struct {
	Name     string `header:"Name"`
	Provider string `header:"Provider"`
	Country  string `header:"Country"`
}

func parseDatacenter(object models.Datacenter, output string) (string, error) {
	return parseDatacenters([]models.Datacenter{object}, output, Name)
}

func parseDatacenters(objects []models.Datacenter, output string, sortBy string) (string, error) {
	var err error
	var parsedOutput []byte

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(objects, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(objects)

	case Text:
		rendered := make([]datacenterRender, 0)
		for _, object := range objects {
			if object.Spec.Country == "" && object.Spec.Provider == "" {
				continue
			}

			rendered = append(rendered, datacenterRender{
				Name:     object.Metadata.Name,
				Country:  object.Spec.Country,
				Provider: object.Spec.Provider,
			})
		}

		sort.Slice(rendered, func(i, j int) bool {
			return rendered[j].Name > rendered[i].Name
		})

		var bodyBuf io.ReadWriter
		bodyBuf = new(bytes.Buffer)

		tableprinter.Print(bodyBuf, rendered)
		parsedOutput, err = ioutil.ReadAll(bodyBuf)

	default:
		return "", errors.New("Unable to parse node deployment")
	}

	return string(parsedOutput), err
}
