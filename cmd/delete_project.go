package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// delProjectsCmd represents the projects command
var delProjectsCmd = &cobra.Command{
	Use:               "project projectid",
	Short:             "Delete a project",
	Example:           "kkpctl delete project dw2s9jk28z",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completion.GetValidProjectArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		projectID := args[0]

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		project, err := kkp.GetProject(projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to find project %s", projectID)
		}

		err = kkp.DeleteProject(args[0])
		if err != nil {
			return errors.Wrapf(err, "failed to delete project %s (%s)", project.Name, projectID)
		}

		fmt.Printf("Successfully deleted Project %s (%s)\n", project.Name, projectID)
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(delProjectsCmd)
}
