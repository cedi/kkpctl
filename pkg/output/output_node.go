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

type nodeRender struct {
	Name                    string   `header:"Name"`
	Addresses               []string `header:"Addresses"`
	Architecture            string   `header:"Architecture"`
	ContainerRuntimeVersion string   `header:"ContainerRuntimeVersion"`
	KernelVersion           string   `header:"KernelVersion"`
	KubeletVersion          string   `header:"KubeletVersion"`
}

func parseNode(object models.Node, output string) (string, error) {
	return parseNodes([]models.Node{object}, output)
}

func parseNodes(objects []models.Node, output string) (string, error) {
	var err error
	var parsedOutput []byte

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(objects, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(objects)

	case Text:
		rendered := make([]nodeRender, 0)
		for _, node := range objects {
			addressRender := make([]string, 0)
			for _, address := range node.Status.Addresses {
				addressRender = append(addressRender, fmt.Sprintf("%s/%s", address.Type, address.Address))
			}

			rendered = append(rendered, nodeRender{
				Name:                    node.Name,
				Addresses:               addressRender,
				Architecture:            node.Status.NodeInfo.Architecture,
				ContainerRuntimeVersion: node.Status.NodeInfo.ContainerRuntimeVersion,
				KernelVersion:           node.Status.NodeInfo.KernelVersion,
				KubeletVersion:          node.Status.NodeInfo.KubeletVersion,
			})
		}

		sort.Slice(rendered, func(i, j int) bool {
			return rendered[j].Name < rendered[i].Name
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
