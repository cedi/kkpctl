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

type eventRender struct {
	CreationTimestamp string `header:"Created"`
	LastTimestamp     string `header:"Last"`
	Type              string `header:"Type"`
	Message           string `header:"Message"`
	Count             int32  `header:"Count"`
}

func (r eventRender) ParseCollection(inputObj interface{}, output string, sortBy string) (string, error) {
	var err error
	var parsedOutput []byte

	objects, ok := inputObj.([]models.Event)
	if !ok {
		return "", fmt.Errorf("inputObj is not a []models.Event")
	}

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(objects, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(objects)

	case Text:
		rendered := make([]eventRender, 0)
		for _, evnt := range objects {
			rendered = append(rendered, eventRender{
				CreationTimestamp: evnt.CreationTimestamp.String(),
				LastTimestamp:     evnt.LastTimestamp.String(),
				Type:              evnt.Type,
				Message:           evnt.Message,
				Count:             evnt.Count,
			})
		}

		sort.Slice(rendered, func(i, j int) bool {
			if sortBy == Date {
				return rendered[j].LastTimestamp > rendered[i].LastTimestamp
			}

			return rendered[j].Type > rendered[i].Type
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
