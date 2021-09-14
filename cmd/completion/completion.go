package completion

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/cedi/kkpctl/pkg/config"
	"github.com/cedi/kkpctl/pkg/model"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/spf13/cobra"
)

// Config holds the global configuration for kkpctl (a pointer which is assigned in root.go)
var Config *config.Config

// GetValidProjectArgs TODO
func GetValidProjectArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	kkp, err := Config.GetKKPClient()
	if err != nil {
		return completions, cobra.ShellCompDirectiveError
	}

	listAll, err := cmd.Flags().GetBool("all")
	if err != nil {
		return completions, cobra.ShellCompDirectiveError
	}

	projects, err := kkp.ListProjects(listAll)
	if err != nil {
		return completions, cobra.ShellCompDirectiveError
	}

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for _, project := range projects {
		if toCompleteRegexp.MatchString(project.ID) {
			completions = append(completions, project.ID)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

// GetValidClusterArgs TODO
func GetValidClusterArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	kkp, err := Config.GetKKPClient()
	if err != nil {
		return completions, cobra.ShellCompDirectiveError
	}

	projects := make([]models.Project, 0)

	projectStr, err := cmd.Flags().GetString("project")
	if err != nil {
		return completions, cobra.ShellCompDirectiveNoFileComp
	}

	if projectStr == "" {
		projects, err = kkp.ListProjects(false)
		if err == nil {
			projects = append(projects, projects...)
		}
	} else {
		project, err := kkp.GetProject(projectStr)
		if err == nil {
			projects = append(projects, *project)
		}
	}

	clusters := make([]models.Cluster, 0)

	for _, projectTmp := range projects {
		clusterTmps, _ := kkp.ListClusters(projectTmp.ID)
		for _, clusterTmp := range clusterTmps {
			clusters = append(clusters, clusterTmp.Cluster)
		}
	}

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for _, cluster := range clusters {
		if toCompleteRegexp.MatchString(cluster.ID) {
			completions = append(completions, cluster.ID)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

// GetValidDatacenterArgs TODO
func GetValidDatacenterArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	kkp, err := Config.GetKKPClient()
	if err != nil {
		return completions, cobra.ShellCompDirectiveError
	}

	var datacenters []models.Datacenter

	projectStr, err := cmd.Flags().GetString("project")
	if err == nil {
		if projectStr == "" {
			datacenters, _ = kkp.ListDatacenter()
		} else {
			cluster, err := kkp.GetCluster(args[0], projectStr)
			if err == nil {
				datacenter, _ := kkp.GetDatacenter(cluster.Cluster.Spec.Cloud.DatacenterName)
				datacenters = append(datacenters, *datacenter)
			}
		}
	}

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for _, dc := range datacenters {
		if toCompleteRegexp.MatchString(dc.Metadata.Name) {
			if dc.Spec.Country == "" && dc.Spec.Provider == "" {
				continue
			}
			completions = append(completions, dc.Metadata.Name)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

// GetValidCloudContextArgs TODO
func GetValidCloudContextArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for key := range Config.Cloud {
		if toCompleteRegexp.MatchString(key) {
			completions = append(completions, key)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

// GetValidKubernetesVersions TODO
func GetValidKubernetesVersions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	kkp, err := Config.GetKKPClient()
	if err != nil {
		return completions, cobra.ShellCompDirectiveError
	}

	clusterVersions, err := kkp.ListClusterVersions()
	if err != nil {
		return completions, cobra.ShellCompDirectiveNoFileComp
	}

	// Sort first, so the highes K8s version is returned first
	sort.Sort(clusterVersions)

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for _, version := range clusterVersions {
		if toCompleteRegexp.MatchString(version.Version) {
			completions = append(completions, version.Version)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

// GetValidClusterTypes TODO
func GetValidClusterTypes(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{"kubernetes"}, cobra.ShellCompDirectiveNoFileComp
}

// GetValidProvider TODO
func GetValidProvider(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for _, name := range Config.Provider.GetAllProviderNames() {
		if toCompleteRegexp.MatchString(name) {
			completions = append(completions, name)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

// GetValidOperatingSystem TODO
func GetValidOperatingSystem(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for _, name := range Config.OSSpec.GetValidOSSpecNames() {
		if toCompleteRegexp.MatchString(string(name)) {
			completions = append(completions, string(name))
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

// GetValidFlavorArgs TODO
func GetValidFlavorArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{}, cobra.ShellCompDirectiveNoFileComp
}

// GetValidNodeSpecArgs TODO
func GetValidNodeSpecArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for _, name := range Config.NodeSpec.GetAllNodeSpecNames() {
		if toCompleteRegexp.MatchString(name) {
			completions = append(completions, name)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

// GetValidMachineDeploymentArgs TODO
func GetValidMachineDeploymentArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	kkp, err := Config.GetKKPClient()
	if err != nil {
		return completions, cobra.ShellCompDirectiveError
	}

	projectID, err := cmd.Flags().GetString("project")
	if err != nil {
		return completions, cobra.ShellCompDirectiveNoFileComp
	}

	clusterID, err := cmd.Flags().GetString("cluster")
	if err != nil {
		return completions, cobra.ShellCompDirectiveNoFileComp
	}

	var cluster *model.ProjectCluster
	cluster, err = kkp.GetCluster(clusterID, projectID)

	if err != nil {
		return completions, cobra.ShellCompDirectiveNoFileComp
	}

	machineDeployments, err := kkp.GetMachineDeployments(cluster.ID, projectID)
	if err != nil {
		return completions, cobra.ShellCompDirectiveNoFileComp
	}

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for _, nd := range machineDeployments {
		if toCompleteRegexp.MatchString(nd.ID) {
			completions = append(completions, nd.ID)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

// GetValidToVersionArgs TODO
func GetValidToVersionArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	kkp, err := Config.GetKKPClient()
	if err != nil {
		return completions, cobra.ShellCompDirectiveError
	}

	// clusterID
	if len(args) == 0 {
		return completions, cobra.ShellCompDirectiveNoFileComp
	}
	clusterID := args[0]

	projectID, err := cmd.Flags().GetString("project")
	if err != nil {
		return completions, cobra.ShellCompDirectiveNoFileComp
	}

	if err != nil {
		return completions, cobra.ShellCompDirectiveNoFileComp
	}

	upgradeVersions, err := kkp.GetClusterUpgradeVersions(clusterID, projectID)
	if err != nil {
		return completions, cobra.ShellCompDirectiveNoFileComp
	}

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for _, version := range upgradeVersions {
		if toCompleteRegexp.MatchString(version.Version) {
			completions = append(completions, version.Version)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}
