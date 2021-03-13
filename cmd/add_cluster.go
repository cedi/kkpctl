package cmd

import (
	"fmt"
	"strings"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/output"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var createClusterCmd = &cobra.Command{
	Use:     "cluster name",
	Short:   "Lets you create a new cluster",
	Example: "kkpctl create cluster test --project z7qbzk5mn4 --labels=\"stage=dev,costcentre=123456\"",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.Wrap(err, "Could not initialize Kubermatic API client")
		}

		mapLabels := make(map[string]string, 0)
		if labels != "" {
			slicedLabels := strings.Split(labels, ",")
			for _, slicedLabel := range slicedLabels {
				splitLabel := strings.Split(slicedLabel, "=")
				mapLabels[splitLabel[0]] = splitLabel[1]
			}
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
	addCmd.AddCommand(createClusterCmd)

	getClustersCmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project.")
	getClustersCmd.MarkFlagRequired("project")
	getClustersCmd.RegisterFlagCompletionFunc("project", getValidProjectArgs)

	getClustersCmd.Flags().StringVarP(&datacenter, "datacenter", "d", "", "Name of the datacenter.")
	getClustersCmd.MarkFlagRequired("datacenter")
	getClustersCmd.RegisterFlagCompletionFunc("datacenter", getValidDatacenterArgs)

	createClusterCmd.Flags().StringVarP(&labels, "labels", "l", "", "A comma separated list of labels in the format key=value")
}
