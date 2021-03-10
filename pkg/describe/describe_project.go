package describe

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/cedi/kkpctl/pkg/output"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/lensesio/tableprinter"
)

// projectRender is a intermediate struct to make use of lensesio/tableprinter, which relies on the header anotation
type projectMetaStruct struct {
	ID                string `header:"ProjectID"`
	Name              string `header:"Name"`
	Status            string `header:"Status"`
	CreationTimestamp string `header:"Created"`
}

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

	var ownerRenderBuf io.ReadWriter
	ownerRenderBuf = new(bytes.Buffer)
	tableprinter.Print(ownerRenderBuf, ownerMeta)
	ownerRenderBytes, err := ioutil.ReadAll(ownerRenderBuf)

	labels := make([]string, 0)
	for key, value := range project.Labels {
		labels = append(labels, fmt.Sprintf("%s=%s", key, value))
	}
	if len(labels) == 0 {
		labels = append(labels, "[None]")
	}

	result := fmt.Sprintf("Project:\n%s\n\nOwners:\n%s\n\nLabels:\n%s\n\nClusters in this Project: %d",
		string(projectTable),
		string(ownerRenderBytes),
		strings.Join(labels, "; "),
		project.ClustersNumber,
	)

	return result, err
}
