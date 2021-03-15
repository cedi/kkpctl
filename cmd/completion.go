package cmd

import (
	"fmt"
	"os"
	"regexp"
	"sort"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/utils"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

  $ source <(kkpctl completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ kkpctl completion bash > /etc/bash_completion.d/kkpctl
  # macOS:
  $ kkpctl completion bash > /usr/local/etc/bash_completion.d/kkpctl

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ kkpctl completion zsh > "${fpath[1]}/_kkpctl"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ kkpctl completion fish | source

  # To load completions for each session, execute once:
  $ kkpctl completion fish > ~/.config/fish/completions/kkpctl.fish

PowerShell:

  PS> kkpctl completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> kkpctl completion powershell > kkpctl.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}

func getValidProjectArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	baseURL, apiToken := Config.GetCloudFromContext()
	kkp, err := client.NewClient(baseURL, apiToken)
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

func getValidClusterArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	baseURL, apiToken := Config.GetCloudFromContext()
	kkp, err := client.NewClient(baseURL, apiToken)
	if err != nil {
		fmt.Println(err.Error())
		return completions, cobra.ShellCompDirectiveError
	}

	var projects []models.Project

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
			projects = append(projects, project)
		}
	}

	var clusters []models.Cluster

	for _, projectTmp := range projects {
		clusterTmp, _ := kkp.ListClustersInProject(projectTmp.ID)
		clusters = append(clusters, clusterTmp...)
	}

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for _, cluster := range clusters {
		if toCompleteRegexp.MatchString(cluster.ID) {
			completions = append(completions, cluster.ID)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func getValidDatacenterArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	baseURL, apiToken := Config.GetCloudFromContext()
	kkp, err := client.NewClient(baseURL, apiToken)
	if err != nil {
		return completions, cobra.ShellCompDirectiveError
	}

	var datacenters []models.Datacenter

	if utils.IsOneOf(cmd.Name(), getClustersCmd.Name(), describeClusterCmd.Name(), delClusterCmd.Name()) {
		projectStr, _ := cmd.Flags().GetString("project")
		if projectStr == "" {
			datacenters, _ = kkp.ListDatacenter()
		} else {
			cluster, err := kkp.GetClusterInProject(args[0], projectStr)
			if err == nil {
				datacenter, _ := kkp.GetDatacenter(cluster.Spec.Cloud.DatacenterName)
				datacenters = append(datacenters, datacenter)
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

func getValidCloudContextArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for key := range Config.Cloud {
		if toCompleteRegexp.MatchString(key) {
			completions = append(completions, key)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}

func getValidKubernetesVersions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	baseURL, apiToken := Config.GetCloudFromContext()
	kkp, err := client.NewClient(baseURL, apiToken)
	if err != nil {
		fmt.Println(err.Error())
		return completions, cobra.ShellCompDirectiveError
	}

	clusterVersions, err := kkp.GetClusterVersions()
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

func getValidClusterTypes(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{"kubernetes"}, cobra.ShellCompDirectiveNoFileComp
}

func getValidProvider(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for _, name := range Config.Provider.GetAllProviderNames() {
		if toCompleteRegexp.MatchString(name) {
			completions = append(completions, name)
		}
	}

	return completions, cobra.ShellCompDirectiveNoFileComp
}
