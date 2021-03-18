package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/describe"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var describeProjectsCmd = &cobra.Command{
	Use:               "project projectid",
	Short:             "Describe a project.",
	Example:           "kkpctl describe project dw2s9jk28z",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: getValidProjectArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		project, err := kkp.GetProject(args[0])
		if err != nil {
			return errors.Wrap(err, "Error fetching project")
		}

		parsed, err := describe.Object(&project)
		if err != nil {
			return errors.Wrap(err, "Error describing project")
		}
		fmt.Print(parsed)

		return nil
	},
}

func init() {
	describeCmd.AddCommand(describeProjectsCmd)
}
