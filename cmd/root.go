package cmd

import (
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/toolSelfUpdate"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/defaults"
	"github.com/wplib/deploywp/deploywp"
)


const DefaultJsonFile       = "deploywp.json"
const DefaultTemplateFile   = "deploywp.tmpl"
const DefaultTemplateString = `{{ BuildDeployWp .Json .Exec.Args }}`
//const DefaultTemplateString = `{{ $dwp := LoadDeployWp .Json (.Exec.GetArg 1) }}{{ $dwp.ExitOnError }}{{ $dwp.Run }}`
//const DefaultTemplateString = `{{ $dwp := LoadDeployWp .Json .Exec.Args }}{{ $dwp.ExitOnError }}{{ $state := $dwp.Run }}{{ $state.ExitOnError }}`

var CmdSelfUpdate *toolSelfUpdate.TypeSelfUpdate
var CmdScribe *loadTools.TypeScribeArgs


var rootCmd = &cobra.Command{
	Use:   defaults.BinaryName,
	Short: "Pantheon release tool.",
	Long: `Feed me a deploywp.json file and I'll do the rest.`,
	Run: gbRootFunc,
}


func init() {
	for range onlyOnce {
		if CmdScribe == nil {
			CmdScribe = loadTools.New(defaults.BinaryName, defaults.BinaryVersion, false)
			CmdScribe.Runtime.SetRepos(defaults.SourceRepo, defaults.BinaryRepo)
			if CmdScribe.State.IsNotOk() {
				break
			}

			CmdScribe.Json.SetDefaults(DefaultJsonFile, "")
			CmdScribe.Template.SetDefaults(DefaultTemplateFile, DefaultTemplateString)

			// Import additional tools.
			CmdScribe.ImportTools(&deploywp.GetTools)
			if CmdScribe.State.IsNotOk() {
				break
			}

			CmdScribe.LoadCommands(rootCmd, false)
			if CmdScribe.State.IsNotOk() {
				break
			}

			CmdScribe.AddConfigOption(false, false)
			if CmdScribe.State.IsNotOk() {
				break
			}

			// This executable is based on Scribe, but we are going to disable some things that we don't need.
			CmdScribe.FlagHide(loadTools.FlagScribeFile)
			CmdScribe.FlagSetDefault(loadTools.FlagJsonFile, DefaultJsonFile)
			CmdScribe.FlagSetDefault(loadTools.FlagTemplateFile, DefaultTemplateString)
			//CmdScribe.FlagSetDefault(loadTools.FlagOutputFile, loadTools.DefaultOutFile)
			//CmdScribe.FlagSetDefault(loadTools.FlagWorkingPath, loadTools.DefaultWorkingPath)

			//CmdScribe.FlagSetDefault(loadTools.FlagChdir, "false")
			CmdScribe.FlagHide(loadTools.FlagRemoveTemplate)
			CmdScribe.FlagHide(loadTools.FlagForce)
			CmdScribe.FlagHide(loadTools.FlagRemoveOutput)
			//CmdScribe.FlagSetDefault(loadTools.FlagQuiet, "false")

			//CmdScribe.FlagSetDefault(loadTools.FlagDebug, "false")

			CmdScribe.FlagHide(loadTools.FlagHelpAll)
			CmdScribe.FlagHide(loadTools.FlagHelpVariables)
			CmdScribe.FlagHide(loadTools.FlagHelpFunctions)
			CmdScribe.FlagHide(loadTools.FlagHelpExamples)
		}

		if CmdSelfUpdate == nil {
			CmdSelfUpdate = toolSelfUpdate.New(CmdScribe.Runtime)
			CmdSelfUpdate.LoadCommands(rootCmd, false)
			if CmdSelfUpdate.State.IsNotOk() {
				break
			}
		}

		CmdScribe.SetHelp(rootCmd)
	}
}


func gbRootFunc(cmd *cobra.Command, args []string) {
	for range onlyOnce {
		var ok bool
		fl := cmd.Flags()

		// Show version.
		ok, _ = fl.GetBool(loadTools.FlagVersion)
		if ok {
			CmdSelfUpdate.VersionShow()
			CmdScribe.State.SetOk()
			break
		}

		if CmdScribe.ParseScribeFlags(cmd) {
			break
		}

		CmdScribe.State = CmdScribe.ProcessArgs(cmd.Use, args)
		if CmdScribe.State.IsNotOk() {
			_ = cmd.Help()
			break
		}

		// Show help if no commands specified.
		if len(args) == 0 {
			_ = cmd.Help()
			CmdScribe.State.SetOk()
			break
		}
	}
}


func Execute() *ux.State {
	for range onlyOnce {
		err := rootCmd.Execute()
		if err != nil {
			CmdScribe.State.SetError(err)
			break
		}

		CmdScribe.State = CheckReturns()
	}

	return CmdScribe.State
}


func CheckReturns() *ux.State {
	state := CmdScribe.State
	for range onlyOnce {
		if CmdScribe.State.IsNotOk() {
			state = CmdScribe.State
			break
		}

		if CmdSelfUpdate.State.IsNotOk() {
			state = CmdSelfUpdate.State
			break
		}

		//if CobraHelp.State.IsNotOk() {
		//	state = CobraHelp.State
		//	break
		//}
	}
	return state
}
