package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/describe"
	"github.com/cedi/kkpctl/pkg/errors"
	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var describeProjectsCmd = &cobra.Command{
	Use:               "project projectid",
	Short:             "Describe a project",
	Example:           "kkpctl describe project dw2s9jk28z",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: getValidProjectArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := args[0]

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		project, err := kkp.GetProject(projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to get project %s", projectID)
		}

		parsed, err := describe.Object(&project)
		if err != nil {
			return errors.Wrapf(err, "failed to describe project %s", projectID)
		}
		fmt.Print(parsed)

		return nil
	},
}

func init() {
	describeCmd.AddCommand(describeProjectsCmd)
}
