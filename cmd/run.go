package cmd

import (
	"github.com/spf13/cobra"
	"github.com/newclarity/JsonToConfig/jtc"
	"github.com/newclarity/JsonToConfig/ux"
)


func init() {
	rootCmd.AddCommand(runCmd)
}


var runCmd = &cobra.Command{
	Use:   jtc.CmdRun,
	Short: ux.SprintfBlue("Execute resulting output file as a BASH script."),
	Long: ux.SprintfBlue(`Execute resulting output file as a BASH script.
You can also use this command as the start to '#!' scripts.
For example: #!/usr/bin/env jtc --json gearbox.json run
`),
	Run: Run,
}
func Run(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		Cmd.ExecShell = true
		Cmd.Output.Name = jtc.SelectConvert

		/*
			Allow this to be used as a UNIX script.
			The following should be placed on the first line.
			#!/usr/bin/env jtc load
		*/

		tmpl := ProcessArgs(cmd, args)
		Cmd.State = tmpl.State
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		ux.PrintflnOk("Executing file '%s' => '%s'", tmpl.Template.Name, tmpl.Output.Name)
		Cmd.State = tmpl.Load()
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		ux.PrintflnOk("Loading file '%s' => '%s'", tmpl.Template.Name, tmpl.Output.Name)
		Cmd.State = tmpl.Run()
		Cmd.State.PrintResponse()
	}
}
