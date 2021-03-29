package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/cedi/kkpctl/pkg/config"
	"github.com/lensesio/tableprinter"
	"gopkg.in/yaml.v3"
)

// clusterRender is a intermediate struct to make use of lensesio/tableprinter, which relies on the header anotation
type configCloudRender struct {
	Name string `header:"Name"`
	URL  string `header:"Url"`
}

func (r configCloudRender) ParseObject(inputObj interface{}, output string) (string, error) {
	var err error
	var parsedOutput []byte
	var cfg config.CloudConfig

	switch object := inputObj.(type) {
	case config.CloudConfig:
		cfg = object

	case *config.CloudConfig:
		cfg = *object

	default:
		return "", fmt.Errorf("inputObj is neighter a config.CloudConfig nor a *config.CloudConfig")
	}

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(cfg, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(cfg)

	case Text:
		rendered := make([]configCloudRender, 0)
		for key, value := range cfg {
			rendered = append(rendered, configCloudRender{
				Name: key,
				URL:  value.URL,
			})
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
