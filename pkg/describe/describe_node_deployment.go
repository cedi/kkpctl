package describe

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/output"
	"github.com/cedi/kkpctl/pkg/utils"
	"github.com/kubermatic/go-kubermatic/models"
)

// NodeDeploymentDescribeMeta contains all the necessary fields to describe a cluster
type NodeDeploymentDescribeMeta struct {
	NodeDeployment *models.NodeDeployment
	Nodes          []models.Node
	NodeEvents     []models.Event
}

// describeProject takes any KKP Cluster and describes it
func describeNodeDeployment(meta *NodeDeploymentDescribeMeta) (string, error) {
	nd := meta.NodeDeployment
	evnt := meta.NodeEvents
	nodes := meta.Nodes

	nodeDeploymentTable, err := output.ParseOutput(nd, output.Text, output.Name)
	if err != nil || len(nodeDeploymentTable) == 0 {
		return "", err
	}

	nodeTable, err := output.ParseOutput(nodes, output.Text, output.Name)
	if err != nil || len(nodeTable) == 0 {
		return "", err
	}

	nodeTaintsTable, err := output.ParseOutput(nd.Spec.Template.Taints, output.Text, output.Name)
	if err != nil || len(nodeTaintsTable) == 0 {
		nodeTaintsTable = "[None]"
	}

	nodeEventTable, err := output.ParseOutput(evnt, output.Text, output.Name)
	if err != nil || len(nodeEventTable) == 0 {
		nodeEventTable = "[None]"
	}

	result := fmt.Sprintf(`Node Deployment:
%s

Nodes:
%s

Taints:
%s

Labels:
%s

Events:
%s`,
		nodeDeploymentTable,
		nodeTable,
		nodeTaintsTable,
		utils.MergeLabels(nd.Spec.Template.Labels),
		nodeEventTable,
	)

	return result, nil
}
