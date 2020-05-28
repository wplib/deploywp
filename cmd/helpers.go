package cmd

import (
	"github.com/spf13/cobra"
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
		//var tmpl *jtc.ArgTemplate

		tmpl := ProcessArgs(rootCmd, args)
		// Ignore errors as there's no args.

		tmpl.PrintHelpers()
	}
}