package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/output"
	"github.com/cedi/kkpctl/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	labels string
)

// projectCmd represents the project command
var createProjectCmd = &cobra.Command{
	Use:     "project name",
	Short:   "Lets you create a new project",
	Example: "kkpctl create project test --labels=\"stage=dev,costcentre=123456\"",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		project, err := kkp.CreateProject(projectName, utils.SplitLabelString(labels))
		if err != nil {
			return errors.Wrapf(err, "failed to create project %s", projectName)
		}

		parsed, err := output.ParseOutput(project, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "failed to parse output")
		}
		fmt.Print(parsed)
		return nil
	},
}

func init() {
	addCmd.AddCommand(createProjectCmd)
	AddLabelsFlag(createProjectCmd)
}
