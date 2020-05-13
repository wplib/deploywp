package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/jsonTemplate"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)


// releaseCmd represents the release command
var helpersCmd = &cobra.Command{
	Use:   "helpers",
	Short: ux.SprintfBlue("Show all built-in template helpers."),
	Long: ux.SprintfBlue(`...`),
	Run: Helpers,
}

func init() {
	rootCmd.AddCommand(helpersCmd)
}

func Helpers(cmd *cobra.Command, args []string) {
	for range only.Once {
		var tmpl *jsonTemplate.Template

		tmpl, _ = ProcessArgs(cmd, args)
		// Ignore errors as there's no args.
		_ = tmpl.SetVersion(Version)

		 tmpl.PrintHelpers()
	}
}