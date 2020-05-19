package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/jsonTemplate"
	"github.com/wplib/deploywp/jsonTemplate/helpers/deploywp"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)


// releaseCmd represents the release command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: ux.SprintfBlue("Process and release a template file"),
	Long: ux.SprintfBlue(`...`),
	Run: Run,
}

func init() {
	rootCmd.AddCommand(runCmd)
}


// This is intended to replace functional template driven workflow with pure GoLang.
// All the same structures and methods are available in both.
func Run(cmd *cobra.Command, args []string) {
	for range only.Once {
		state := ux.NewState()
		var tmpl *jsonTemplate.Template

		tmpl, *state = ProcessArgs(cmd, args)
		if state.IsNotOk() {
			fmt.Printf(state.PrintResponse())
			break
		}

		_ = tmpl.SetVersion(Version)

		state = tmpl.LoadJson()
		if state.IsNotOk() {
			fmt.Printf(state.PrintResponse())
			break
		}

		//dwp := deploywp.TypeDeployWp{}
		dwp := deploywp.HelperLoadDeployWp(tmpl.JsonStruct.Json, tmpl.GetArgs()...)
		if dwp.State.IsNotOk() {
			fmt.Printf(dwp.State.PrintResponse())
			break
		}

		dwp.Exec = tmpl.JsonStruct.Exec
		state = dwp.Run()

		if state.IsNotOk() {
			state.SetExitCode(1)
			state.ExitOnNotOk()
			break
		}

		fmt.Printf("\n%s\nFINISHED\n", state.PrintResponse())
	}
}