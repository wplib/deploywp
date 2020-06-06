package cmd

import (
	"github.com/spf13/cobra"
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/ux"
)


func init() {
	rootCmd.AddCommand(loadCmd)
}


// releaseCmd represents the release command
var loadCmd = &cobra.Command{
	Use:   loadTools.CmdLoad,
	Short: ux.SprintfBlue("Load and execute a template file."),
	Long: ux.SprintfBlue("Load and execute a template file."),
	Run: LoadTemplate,
	DisableFlagParsing: false,
}
func LoadTemplate(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		tmpl := ProcessArgs(cmd, args)
		Cmd.State = tmpl.State
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		Cmd.State = tmpl.Load()
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		ux.PrintflnOk("Loading template '%s' and saving result to '%s'", tmpl.Template.Name, tmpl.Output.Name)
		Cmd.State = tmpl.Run()
		Cmd.State.PrintResponse()
	}
}
