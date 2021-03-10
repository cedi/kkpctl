package cmd

import (
	"fmt"
	"regexp"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/spf13/cobra"
)

func getValidProjectArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	kkp, err := client.NewClient(baseURL, apiToken)
	if err != nil {
		return completions, cobra.ShellCompDirectiveError
	}

	projects, err := kkp.ListProjects(listAll)

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for _, project := range projects {
		if toCompleteRegexp.MatchString(project.ID) {
			completions = append(completions, project.ID)
		}
	}

	return completions, cobra.ShellCompDirectiveDefault
}

func getValidClusterArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := make([]string, 0)

	kkp, err := client.NewClient(baseURL, apiToken)
	if err != nil {
		return completions, cobra.ShellCompDirectiveError
	}

	projects, err := kkp.ListProjects(listAll)

	toCompleteRegexp := regexp.MustCompile(fmt.Sprintf("^%s.*$", toComplete))
	for _, project := range projects {
		if toCompleteRegexp.MatchString(project.ID) {
			completions = append(completions, project.ID)
		}
	}

	return completions, cobra.ShellCompDirectiveDefault
}
