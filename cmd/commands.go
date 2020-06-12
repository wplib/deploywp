package cmd

import (
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
)

const onlyOnce = "1"
var onlyTwice = []string{"", ""}


func init() {
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(convertCmd)
	rootCmd.AddCommand(helpersCmd)
	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(runCmd)
}


var buildCmd = &cobra.Command{
	Use:   loadTools.CmdBuild,
	Short: ux.SprintfBlue("Build a Pantheon website."),
	Long: ux.SprintfBlue("Build a Pantheon website."),
	Run: cmdBuild,
}

var helpersCmd = &cobra.Command{
	Use:   loadTools.CmdTools,
	Short: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Show all built-in template helpers."),
	Long:  ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Show all built-in template helpers."),
	Run:   cmdTools,
}

var convertCmd = &cobra.Command{
	Use:   loadTools.CmdConvert,
	Short: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Convert a template file to the resulting output file."),
	Long: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Convert a template file to the resulting output file."),
	Run: cmdConvert,
}
var loadCmd = &cobra.Command{
	Use:   loadTools.CmdLoad,
	Short: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Load and execute a template file."),
	Long: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Load and execute a template file."),
	Run: cmdLoad,
	DisableFlagParsing: false,
}
var runCmd = &cobra.Command{
	Use:   loadTools.CmdRun,
	Short: ux.SprintfMagenta("scribe") + ux.SprintfBlue(" - Execute resulting output file as a BASH script."),
	Long: ux.SprintfMagenta("scribe") + ux.SprintfBlue(`Execute resulting output file as a BASH script.
You can also use this command as the start to '#!' scripts.
For example: #!/usr/bin/env scribe --json gearbox.json run
`),
	Run: cmdRun,
}


func init() {
	rootCmd.AddCommand(helpersCmd)
}