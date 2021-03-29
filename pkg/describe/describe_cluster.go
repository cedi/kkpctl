package describe

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/output"
	"github.com/cedi/kkpctl/pkg/utils"
	"github.com/kubermatic/go-kubermatic/models"
)

// ClusterDescribeMeta contains all the necessary fields to describe a cluster
type ClusterDescribeMeta struct {
	Cluster         *models.Cluster
	NodeDeployments []models.NodeDeployment
	ClusterHealth   *models.ClusterHealth
	ClusterEvents   []models.Event
}

// describeProject takes any KKP Cluster and describes it
func describeCluster(meta *ClusterDescribeMeta) (string, error) {
	cluster := meta.Cluster
	nd := meta.NodeDeployments
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

	nodeDeploymentTable, err := output.ParseOutput(nd, output.Text, output.Name)
	if err != nil || len(nodeDeploymentTable) == 0 {
		nodeDeploymentTable = "[None]"
	}

	clusterEventTable, err := output.ParseOutput(evnt, output.Text, output.Name)
	if err != nil || len(clusterEventTable) == 0 {
		clusterEventTable = "[None]"
	}

	result := fmt.Sprintf(`Cluster:
%s

Health Status:
%s

Node Deployments:
%s

Inherited Labels:
%s

Labels:
%s

Events:
%s`,
		clusterTable,
		clusterHealthTable,
		nodeDeploymentTable,
		utils.MergeLabels(cluster.InheritedLabels),
		utils.MergeLabels(cluster.Labels),
		clusterEventTable,
	)

	return result, nil
}
