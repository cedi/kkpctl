package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var listAll bool

// projectsCmd represents the projects command
var getProjectsCmd = &cobra.Command{
	Use:               "project [projectid]",
	Short:             "List a project.",
	Long:              `If no projectid is specified, all projects of an account are listed. If a projectid is specified only this project is shown`,
	Example:           "kkpctl get project --all",
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: getValidProjectArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.New("Could not initialize Kubermatic API client")
		}

		var result interface{}
		if len(args) == 0 || listAll {
			result, err = kkp.ListProjects(listAll)
		} else {
			result, err = kkp.GetProject(args[0])
		}

		if err != nil {
			return errors.Wrap(err, "Error fetching project")
		}

		parsed, err := output.ParseOutput(result, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "Error parsing project")
		}
		fmt.Println(parsed)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getProjectsCmd)
	getProjectsCmd.Flags().BoolVarP(&listAll, "all", "a", false, "Display all projects the users is allowed to see.")
}
