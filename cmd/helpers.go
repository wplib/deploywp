package cmd

import (
	"github.com/spf13/cobra"
	"github.com/newclarity/scribeHelpers/scribeLoader"
	"github.com/newclarity/scribeHelpers/ux"
)


func init() {
	rootCmd.AddCommand(helpersCmd)
}


var helpersCmd = &cobra.Command{
	Use:   scribeLoader.CmdHelpers,
	Short: ux.SprintfBlue("Show all built-in template helpers."),
	Long:  ux.SprintfBlue("Show all built-in template helpers."),
	Run:   Helpers,
}
func Helpers(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		tmpl := ProcessArgs(rootCmd, args)
		// Ignore errors as there's no args.

		tmpl.PrintHelpers()
		Cmd.State.Clear()
	}
}
