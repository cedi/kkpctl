package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listAll bool

// projectsCmd represents the projects command
var getProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Lists projects.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not yet implemented")
	},
}

func init() {
	getCmd.AddCommand(getProjectsCmd)
	getProjectsCmd.Flags().BoolVarP(&listAll, "all", "a", false, "Display all projects the users is allowed to see.")
}
