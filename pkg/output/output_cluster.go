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

// clusterRender is a intermediate struct to make use of lensesio/tableprinter, which relies on the header anotation
type clusterRender struct {
	ID                string `header:"ClusterID"`
	Name              string `header:"Name"`
	Version           string `header:"Version"`
	CreationTimestamp string `header:"Created"`
	Datacenter        string `header:"Datacenter"`
	Provider          string `header:"Provider"`
}

func parseCluster(object models.Cluster, output string) (string, error) {
	return parseClusters([]models.Cluster{object}, output, Name)
}

func parseClusters(objects []models.Cluster, output string, sortBy string) (string, error) {
	var err error
	var parsedOutput []byte

	switch output {
	case JSON:
		parsedOutput, err = json.MarshalIndent(objects, "", "  ")

	case YAML:
		parsedOutput, err = yaml.Marshal(objects)

	case Text:
		rendered := make([]clusterRender, len(objects))
		for idx, object := range objects {
			rendered[idx] = clusterRender{
				ID:                object.ID,
				Name:              object.Name,
				CreationTimestamp: object.CreationTimestamp.String(),
				Datacenter:        object.Spec.Cloud.DatacenterName,
				Provider:          getProviderName(object.Spec.Cloud),
			}

			version, ok := object.Status.Version.(string)
			if !ok {
				// Honestly, I don't know wht to do here^^
				continue
			}

			rendered[idx].Version = version
		}

		sort.Slice(rendered, func(i, j int) bool {
			if sortBy == Date {
				return rendered[j].CreationTimestamp < rendered[i].CreationTimestamp
			}

			return rendered[j].Name > rendered[i].Name
		})

		var bodyBuf io.ReadWriter
		bodyBuf = new(bytes.Buffer)

		tableprinter.Print(bodyBuf, rendered)
		parsedOutput, err = ioutil.ReadAll(bodyBuf)

	default:
		return "", errors.New("Unable to parse objects")
	}

	return string(parsedOutput), err
}

func getProviderName(cloudSpec *models.CloudSpec) string {
	if cloudSpec.Alibaba != nil {
		return "Alibaba"
	}
	if cloudSpec.Anexia != nil {
		return "Anexia"
	}
	if cloudSpec.Aws != nil {
		return "AWS"
	}
	if cloudSpec.Azure != nil {
		return "Azure"
	}
	if cloudSpec.Bringyourown != nil {
		return "BringYourOwn"
	}
	if cloudSpec.Digitalocean != nil {
		return "DigitalOcean"
	}
	if cloudSpec.Fake != nil {
		return "Fake"
	}
	if cloudSpec.Gcp != nil {
		return "GCP"
	}
	if cloudSpec.Hetzner != nil {
		return "Hetzner"
	}
	if cloudSpec.Kubevirt != nil {
		return "Kubevirt"
	}
	if cloudSpec.Openstack != nil {
		return "OpenStack"
	}
	if cloudSpec.Packet != nil {
		return "Packet"
	}
	if cloudSpec.Vsphere != nil {
		return "Vsphere"
	}

	return ""
}
