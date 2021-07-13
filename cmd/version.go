package cmd

import (
	"fmt"

	"github.com/cedi/kkpctl/pkg/output"
	"github.com/spf13/cobra"
)

var versionCMD = &cobra.Command{
	Use:     "version",
	Short:   "Shows version information",
	Example: "kkpctl version",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		render := make([]output.VersionRender, 0)
		defer func(render *[]output.VersionRender) {
			parsed, err := output.ParseOutput(*render, outputType, sortBy)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

			fmt.Print(parsed)
		}(&render)

		// Prepare kkpctl binary version
		render = append(render, output.VersionRender{
			Component: "kkpctl",
			Version:   Version,
			Date:      Date,
			Commit:    Commit,
			BuiltBy:   BuiltBy,
		})

		// Prepare KKP Server Version
		kkp, err := Config.GetKKPClient()
		if err != nil {
			return
		}

		versions, err := kkp.GetKKPVersion()
		if err != nil {
			return
		}

		render = append(render, output.VersionRender{
			Component: "KKP API",
			Version:   versions["api"],
		})
	},
}

func init() {
	rootCmd.AddCommand(versionCMD)
}
