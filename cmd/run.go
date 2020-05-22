package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/deploywp"
	"github.com/wplib/deploywp/ux"
)


func init() {
	rootCmd.AddCommand(runCmd)
}


var runCmd = &cobra.Command{
	Use:   cmdRun,
	Short: ux.SprintfBlue("Process and release a template file"),
	Long: ux.SprintfBlue(`...`),
	Run: Run,
}
func Run(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		//Cmd.State = ux.NewState(Cmd.Debug)
		//var tmpl *jsonTemplate.ArgTemplate

		tmpl := ProcessArgs(rootCmd, args)
		Cmd.State = tmpl.State
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		Cmd.State = tmpl.LoadTemplate()
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		//dwp := deploywp.TypeDeployWp{}
		dwp := deploywp.HelperLoadDeployWp(tmpl.JsonStruct.Json, tmpl.Exec.GetArgs()...)
		if dwp.State.IsNotOk() {
			dwp.State.PrintResponse()
			break
		}

		dwp.Exec = tmpl.JsonStruct.Exec
		Cmd.State = dwp.Run()

		if Cmd.State.IsNotOk() {
			Cmd.State.SetExitCode(1)
			//Cmd.State.Exit(1)
			break
		}

		fmt.Printf("\n%s\nFINISHED\n", Cmd.State.SprintResponse())
	}
}
