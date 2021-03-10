package cmd

import (
	"fmt"
	"strings"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/output"
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
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.Wrap(err, "Could not initialize Kubermatic API client")
		}

		slicedLabels := strings.Split(labels, ",")
		mapLabels := make(map[string]string, 0)

		for _, slicedLabel := range slicedLabels {
			splitLabel := strings.Split(slicedLabel, "=")
			mapLabels[splitLabel[0]] = splitLabel[1]
		}

		project, err := kkp.CreateProject(args[0], mapLabels)
		if err != nil {
			return errors.Wrap(err, "Error fetching projects")
		}

		parsed, err := output.ParseOutput(project, outputType, sortBy)
		if err != nil {
			return errors.Wrap(err, "Error parsing projects")
		}
		fmt.Print(parsed)
		return nil
	},
}

func init() {
	addCmd.AddCommand(createProjectCmd)
	createProjectCmd.Flags().StringVarP(&labels, "labels", "l", "", "A comma separated list of labels in the format key=value")
}
