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

type taintRender struct {
	Key    string `header:"key"`
	Value  string `header:"value"`
	Effect string `header:"effect"`
}

func parseNodeTaint(object *models.TaintSpec, output string) (string, error) {
	return parseNodeTaints([]*models.TaintSpec{object}, output)
}

func parseNodeTaints(objects []*models.TaintSpec, output string) (string, error) {
	var err error
	var parsedOutput []byte

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(objects, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(objects)

	case Text:
		rendered := make([]taintRender, 0)
		for _, taint := range objects {
			if taint == nil {
				continue
			}

			rendered = append(rendered, taintRender{
				Key:    taint.Key,
				Value:  taint.Value,
				Effect: taint.Effect,
			})
		}

		sort.Slice(rendered, func(i, j int) bool {
			return rendered[j].Key < rendered[i].Key
		})

		var bodyBuf io.ReadWriter
		bodyBuf = new(bytes.Buffer)

		tableprinter.Print(bodyBuf, rendered)
		parsedOutput, err = ioutil.ReadAll(bodyBuf)

	default:
		return "", fmt.Errorf("unable to parse node taint")
	}

	return string(parsedOutput), err
}
