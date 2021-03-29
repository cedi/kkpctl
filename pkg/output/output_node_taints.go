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

func (r taintRender) ParseObject(inputObj interface{}, output string) (string, error) {
	switch object := inputObj.(type) {
	case models.TaintSpec:
		return r.ParseCollection([]models.TaintSpec{object}, output, Name)

	case *models.TaintSpec:
		return r.ParseCollection([]models.TaintSpec{*object}, output, Name)

	default:
		return "", fmt.Errorf("inputObj is neighter a models.TaintSpec nor a *models.TaintSpec")
	}
}

func (r taintRender) ParseCollection(inputObj interface{}, output string, sortBy string) (string, error) {
	var err error
	var parsedOutput []byte

	objects, ok := inputObj.([]models.TaintSpec)
	if !ok {
		return "", fmt.Errorf("inputObj is not a []models.TaintSpec")
	}

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(objects, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(objects)

	case Text:
		rendered := make([]taintRender, 0)
		for _, taint := range objects {
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
