package describe

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/model"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/cedi/kkpctl/pkg/utils"
	"github.com/kubermatic/go-kubermatic/models"
)

// ClusterDescribeMeta contains all the necessary fields to describe a cluster
type ClusterDescribeMeta struct {
	Cluster            *model.ProjectCluster
	MachineDeployments []models.NodeDeployment
	ClusterHealth      *models.ClusterHealth
	ClusterEvents      []models.Event
}

// describeProject takes any KKP Cluster and describes it
func describeCluster(meta *ClusterDescribeMeta) (string, error) {
	cluster := meta.Cluster
	nd := meta.MachineDeployments
	ch := meta.ClusterHealth
	evnt := meta.ClusterEvents

	clusterTable, err := output.ParseOutput(cluster, output.Text, output.Name)
	if err != nil {
		return "", err
	}

	clusterHealthTable, err := output.ParseOutput(ch, output.Text, output.Name)
	if err != nil {
		return "", err
	}

	machineDeploymentTable, err := output.ParseOutput(nd, output.Text, output.Name)
	if err != nil || len(machineDeploymentTable) == 0 {
		machineDeploymentTable = "[None]"
	}

	clusterEventTable, err := output.ParseOutput(evnt, output.Text, output.Name)
	if err != nil || len(clusterEventTable) == 0 {
		clusterEventTable = "[None]"
	}

	result := fmt.Sprintf(`Cluster:
%s

Health Status:
%s

Machine Deployments:
%s

Inherited Labels:
%s

Labels:
%s

Events:
%s`,
		clusterTable,
		clusterHealthTable,
		machineDeploymentTable,
		utils.MergeLabels(cluster.Cluster.InheritedLabels),
		utils.MergeLabels(cluster.Cluster.Labels),
		clusterEventTable,
	)

	return result, nil
}
