package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getFindProjectIDCmd = &cobra.Command{
	Use:     "project-by-clusterid [clusterid]",
	Short:   "Lists find the project for a clusterID",
	Example: "kkpctl find project-by-clusterid gshj9vt55p",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		clusterID := ""
		if len(args) == 1 {
			clusterID = args[0]
		}

		kkp, err := Config.GetKKPClient()
		if err != nil {
			return err
		}

		allProjects, err := kkp.ListProjects(true)
		if err != nil {
			return err
		}

		for _, project := range allProjects {
			allClusterInProject, err := kkp.ListClusters(project.ID)
			if err != nil {
				continue
			}

			for _, cluster := range allClusterInProject {
				if cluster.ID == clusterID {
					fmt.Println(project.ID)
					return nil
				}
			}
		}

		return nil
	},
}

func init() {
	findCmd.AddCommand(getFindProjectIDCmd)
}
