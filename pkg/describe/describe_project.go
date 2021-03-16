package describe

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/cedi/kkpctl/pkg/output"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/lensesio/tableprinter"
)

type ownerStruct struct {
	ID                string `header:"UserID"`
	Name              string `header:"Name"`
	Email             string `header:"Email"`
	CreationTimestamp string `header:"Created"`
}

// describeProject takes any KKP Project and describes it
func describeProject(project *models.Project) (string, error) {
	projectTable, err := output.ParseOutput(*project, output.Text, output.Name)
	if err != nil {
		return "", err
	}

	ownerMeta := make([]ownerStruct, len(project.Owners))
	for idx, owner := range project.Owners {
		ownerMeta[idx] = ownerStruct{
			ID:                owner.ID,
			Name:              owner.Name,
			Email:             owner.Email,
			CreationTimestamp: owner.CreationTimestamp.String(),
		}
	}

	ownerRenderBuf := new(bytes.Buffer)
	tableprinter.Print(ownerRenderBuf, ownerMeta)
	ownerRenderBytes, err := ioutil.ReadAll(ownerRenderBuf)

	labels := make([]string, 0)
	for key, value := range project.Labels {
		labels = append(labels, fmt.Sprintf("%s=%s", key, value))
	}
	if len(labels) == 0 {
		labels = append(labels, "[None]")
	}

	result := fmt.Sprintf("Project:\n%s\n\nOwners:\n%s\n\nLabels:\n%s\n\nClusters in this Project: %d\n",
		string(projectTable),
		string(ownerRenderBytes),
		strings.Join(labels, "; "),
		project.ClustersNumber,
	)

	return result, err
}
