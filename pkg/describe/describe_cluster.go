package describe

import (
	"fmt"
	"strings"

	"github.com/cedi/kkpctl/pkg/output"
	"github.com/kubermatic/go-kubermatic/models"
)

// ClusterDescribeMeta contains all the necessary fields to describe a cluster
type ClusterDescribeMeta struct {
	Cluster         *models.Cluster
	NodeDeployments []models.NodeDeployment
	ClusterHealth   *models.ClusterHealth
}

// describeProject takes any KKP Cluster and describes it
func describeCluster(meta *ClusterDescribeMeta) (string, error) {
	cluster := meta.Cluster
	nd := meta.NodeDeployments
	ch := meta.ClusterHealth

	clusterTable, err := output.ParseOutput(cluster, output.Text, output.Name)
	if err != nil {
		return "", err
	}

	nodeDeploymentTable, err := output.ParseOutput(nd, output.Text, output.Name)
	if err != nil {
		return "", err
	}

	clusterHealthTable, err := output.ParseOutput(ch, output.Text, output.Name)
	if err != nil {
		return "", err
	}

	labels := make([]string, 0)
	for key, value := range cluster.Labels {
		labels = append(labels, fmt.Sprintf("%s=%s", key, value))
	}
	if len(labels) == 0 {
		labels = append(labels, "[None]")
	}

	inheritedLabels := make([]string, 0)
	for key, value := range cluster.InheritedLabels {
		inheritedLabels = append(inheritedLabels, fmt.Sprintf("%s=%s", key, value))
	}
	if len(inheritedLabels) == 0 {
		inheritedLabels = append(inheritedLabels, "[None]")
	}

	result := fmt.Sprintf("Cluster:\n%s\n\nHealth Status:\n%s\n\nNode Deployments:\n%s\n\nInherited Labels:\n%s\n\nLabels:\n%s\n",
		clusterTable,
		clusterHealthTable,
		nodeDeploymentTable,
		strings.Join(inheritedLabels, "; "),
		strings.Join(labels, "; "),
	)

	return result, err
}
