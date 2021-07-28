package cmd

import (
	"github.com/cedi/kkpctl/cmd/completion"
	"github.com/spf13/cobra"
)

// AddProjectFlag adds the --project flag to the cobra.Command
// 	this ensures that the --project flag is always added in the same way and
// 	having the completion functino working
func AddProjectFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&projectID, "project", "p", "", "ID of the project.")
	cmd.MarkFlagRequired("project")
	cmd.RegisterFlagCompletionFunc("project", completion.GetValidProjectArgs)
}

// AddClusterFlag adds the --cluster flag to the cobra.Command
// 	this ensures that the --cluster flag is always added in the same way and
// 	having the completion function working
func AddClusterFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&clusterID, "cluster", "c", "", "ID of the cluster")
	cmd.MarkFlagRequired("cluster")
	cmd.RegisterFlagCompletionFunc("cluster", completion.GetValidClusterArgs)
}

// AddLabelsFlag adds the --labels flag to the cobra.Command
// 	this ensures that the --datacenter flag is always added in the same way
func AddLabelsFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&labels, "labels", "l", "", "A comma separated list of labels in the format key=value")
}
