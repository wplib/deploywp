package cmd

import (
	"github.com/spf13/cobra"
	"github.com/newclarity/scribeHelpers/scribeLoader"
	"github.com/wplib/deploywp/deploywp"
	"github.com/newclarity/scribeHelpers/ux"
)


func init() {
	rootCmd.AddCommand(buildCmd)
}


var buildCmd = &cobra.Command{
	Use:   scribeLoader.CmdBuild,
	Short: ux.SprintfBlue("Build a Pantheon website."),
	Long: ux.SprintfBlue("Build a Pantheon website."),
	Run: Build,
}
func Build(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		tmpl := ProcessArgs(cmd, args)
		Cmd.State = tmpl.State
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		tmpl.LoadHelpers(deploywp.GetHelpers)

		Cmd.State = tmpl.Load()
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		ux.PrintflnOk("Running build.")
		Cmd.State = tmpl.Run()
		Cmd.State.PrintResponse()

		//dwp := deploywp.TypeDeployWp{}
		//dwp := deploywp.HelperLoadDeployWp(tmpl.JsonStruct.Json, tmpl.Exec.GetArgs()...)
		//if dwp.State.IsNotOk() {
		//	dwp.State.PrintResponse()
		//	break
		//}
		//
		//dwp.Exec = tmpl.JsonStruct.Exec
		//Cmd.State = dwp.Run()
		//
		//if Cmd.State.IsNotOk() {
		//	Cmd.State.SetExitCode(1)
		//	//Cmd.State.Exit(1)
		//	break
		//}
		//
		//fmt.Printf("\n%s\nFINISHED\n", Cmd.State.SprintResponse())
	}
}
