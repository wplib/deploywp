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
	//rootCmd.AddCommand(toolsCmd)
	//rootCmd.AddCommand(loadCmd)
	//rootCmd.AddCommand(runCmd)
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


//var toolsCmd = &cobra.Command{
//	Use:   loadTools.CmdTools,
//	Short: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Show all built-in template helpers."),
//	Long:  ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Show all built-in template helpers."),
//	Run:   cmdTools,
//}
//
//var convertCmd = &cobra.Command{
//	Use:   loadTools.CmdConvert,
//	Short: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Convert a template file to the resulting output file."),
//	Long: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Convert a template file to the resulting output file."),
//	Run: cmdConvert,
//}
//var loadCmd = &cobra.Command{
//	Use:   loadTools.CmdLoad,
//	Short: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Load and execute a template file."),
//	Long: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Load and execute a template file."),
//	Run: cmdLoad,
//	DisableFlagParsing: false,
//}
//var runCmd = &cobra.Command{
//	Use:   loadTools.CmdRun,
//	Short: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Execute resulting output file as a BASH script."),
//	Long: ux.SprintfMagenta("scribe") + ux.SprintfBlue(`Execute resulting output file as a BASH script.
//You can also use this command as the start to '#!' scripts.
//For example: #!/usr/bin/env scribe --json gearbox.json run
//`),
//	Run: cmdRun,
//}
