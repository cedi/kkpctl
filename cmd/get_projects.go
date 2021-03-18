package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var listAll bool

// projectsCmd represents the projects command
var getProjectsCmd = &cobra.Command{
	Use:               "project [projectid]",
	Short:             "List a project",
	Example:           "kkpctl get project",
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: getValidProjectArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
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
		fmt.Print(parsed)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getProjectsCmd)
	getProjectsCmd.Flags().BoolVarP(&listAll, "all", "a", false, "Display all projects if the users is allowed to see.")
}
