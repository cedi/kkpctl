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
	Message           string `header:"Message"`
	Count             int32  `header:"Count"`
}

func parseEvents(objects []models.Event, output string) (string, error) {
	var err error
	var parsedOutput []byte

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
				Message:           evnt.Message,
				Count:             evnt.Count,
			})
		}

		sort.Slice(rendered, func(i, j int) bool {
			return rendered[j].LastTimestamp > rendered[i].LastTimestamp
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
