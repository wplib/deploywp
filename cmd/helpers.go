package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/defaults"
	"github.com/wplib/deploywp/jsonTemplate"
	"github.com/wplib/deploywp/ux"
)


func init() {
	rootCmd.AddCommand(helpersCmd)
}


var helpersCmd = &cobra.Command{
	Use:   cmdHelpers,
	Short: ux.SprintfBlue("Show all built-in template helpers."),
	Long: ux.SprintfBlue(`...`),
	Run: Helpers,
}
func Helpers(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		//Cmd.State = ux.NewState(Cmd.Debug)
		var tmpl *jsonTemplate.ArgTemplate

		tmpl = ProcessArgs(cmd, args)
		// Ignore errors as there's no args.
		_ = tmpl.SetVersion(defaults.BinaryVersion)

		 tmpl.PrintHelpers()
	}
}