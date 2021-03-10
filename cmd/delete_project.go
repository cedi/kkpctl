package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// delProjectsCmd represents the projects command
var delProjectsCmd = &cobra.Command{
	Use:   "project projectid",
	Short: "Delete a project.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		kkp, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			return errors.New("Could not initialize Kubermatic API client")
		}

		err = kkp.DeleteProject(args[0])
		if err != nil {
			return errors.Wrap(err, "Error fetching project")
		}

		fmt.Printf("Successfully deleted Project %s", args[0])
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(delProjectsCmd)
}
