package cmd

import (
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/defaults"
)


const onlyOnce = "1"
var onlyTwice = []string{"", ""}


func init() {
	rootCmd.AddCommand(buildCmd)
}


var buildCmd = &cobra.Command{
	Use:   "build",
	Short: ux.SprintfMagenta(defaults.BinaryName) + ux.SprintfBlue(" - Build a Pantheon website."),
	Long: ux.SprintfMagenta(defaults.BinaryName) + ux.SprintfBlue(" - Build a Pantheon website."),
	Args: cobra.RangeArgs(1, 2),
	Run: cmdBuild,
}

func cmdBuild(cmd *cobra.Command, args []string) {
	state := CmdScribe.State

	for range onlyOnce {
		CmdScribe.Chdir = true // In this mode we always change directory to the JSON file.

		state = CmdScribe.ProcessArgs(cmd.Use, args)
		if state.IsNotOk() {
			state.PrintResponse()
			break
		}

		state = CmdScribe.Load()
		if state.IsNotOk() {
			state.PrintResponse()
			break
		}

		ux.PrintflnOk("Building website via deploywp.")
		state = CmdScribe.Run()

		state.PrintResponse()
		ux.PrintflnBlue("\n# FINISHED")
	}

	CmdScribe.State = state
}
