package output

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (r datacenterRender) ParseObject(inputObject interface{}, output string) (string, error) {
	switch object := inputObject.(type) {
	case models.Datacenter:
		return r.ParseCollection([]models.Datacenter{object}, output, Name)

	case *models.Datacenter:
		return r.ParseCollection([]models.Datacenter{*object}, output, Name)

	default:
		return "", fmt.Errorf("inputObj is neighter a models.Datacenter nor a *models.Datacenter")
	}
}

func (r datacenterRender) ParseCollection(inputObj interface{}, output string, sortBy string) (string, error) {
	var err error
	var parsedOutput []byte

	objects, ok := inputObj.([]models.Datacenter)
	if !ok {
		return "", fmt.Errorf("inputObj is not a []models.Datacenter")
	}

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
		return "", fmt.Errorf("unable to parse node deployment")
	}

	return string(parsedOutput), err
}
