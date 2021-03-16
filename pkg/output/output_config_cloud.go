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

func parseConfigCloud(object config.CloudConfig, output string) (string, error) {
	var err error
	var parsedOutput []byte

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(object, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(object)

	case Text:
		rendered := make([]configCloudRender, 0)
		for key, value := range object {
			rendered = append(rendered, configCloudRender{
				Name: key,
				URL:  value,
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
