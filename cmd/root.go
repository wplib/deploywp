package cmd

import (
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/toolSelfUpdate"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/defaults"
	"github.com/wplib/deploywp/deploywp"
)


//type TypeCmd struct {
//	ConfigFile   string
//
//	JsonFile     string
//	TemplateFile string
//	OutFile      string
//
//	Debug        bool
//	Chdir        bool
//	DryRun       bool
//
//	State        *ux.State
//}


const DefaultJsonFile     = "deploywp.json"
//const DefaultTemplateFile = `{{ $dwp := LoadDeployWp .Json (.Exec.GetArg 1) }}{{ $dwp.ExitOnError }}{{ $dwp.Run }}`
//const DefaultTemplateFile = `{{ $dwp := LoadDeployWp .Json .Exec.Args }}{{ $dwp.ExitOnError }}{{ $state := $dwp.Run }}{{ $state.ExitOnError }}`
const DefaultTemplateFile = `{{ BuildDeployWp .Json .Exec.Args }}`


var CmdSelfUpdate *toolSelfUpdate.TypeSelfUpdate
var CmdScribe *loadTools.TypeScribeArgs
//var ConfigFile string
//const 	flagConfigFile  	= "config"


var rootCmd = &cobra.Command{
	Use:   defaults.BinaryName,
	Short: "Pantheon release tool.",
	Long: `Feed me a deploywp.json file and I'll do the rest.`,
	Run: gbRootFunc,
}


func init() {
	SetCmd()
	//cobra.OnInitialize(initConfig)
	//cobra.EnableCommandSorting = false
	//
	//rootCmd.PersistentFlags().StringVar(&ConfigFile, flagConfigFile, fmt.Sprintf("%s-config.json", defaults.BinaryName), ux.SprintfBlue("%s: config file.", defaults.BinaryName))
	//_ = rootCmd.PersistentFlags().MarkHidden(flagConfigFile)

	//rootCmd.PersistentFlags().StringVarP(&CmdScribe.Json.File, loadTools.FlagJsonFile, "j", DefaultJsonFile, ux.SprintfBlue("Alternative JSON file."))
	//rootCmd.PersistentFlags().StringVarP(&CmdScribe.Template.File, loadTools.FlagTemplateFile, "t", DefaultTemplateFile, ux.SprintfBlue("Alternative template file."))
	//rootCmd.PersistentFlags().StringVarP(&CmdScribe.Output.File, loadTools.FlagOutputFile, "o", loadTools.DefaultOutFile, ux.SprintfBlue("Output file."))
	//rootCmd.PersistentFlags().StringVarP(&CmdScribe.WorkingPath.File, loadTools.FlagWorkingPath, "p", loadTools.DefaultWorkingPath, ux.SprintfBlue("Set working path."))
	//
	//rootCmd.PersistentFlags().BoolVarP(&CmdScribe.Chdir, loadTools.FlagChdir, "c", false, ux.SprintfBlue("Change to directory containing %s", DefaultJsonFile))
	////rootCmd.PersistentFlags().BoolVarP(&CmdScribe.RemoveTemplate, loadTools.FlagRemoveTemplate, "", false, ux.SprintfBlue("Remove template file afterwards."))
	////rootCmd.PersistentFlags().BoolVarP(&CmdScribe.ForceOverwrite, loadTools.FlagForce, "f", false, ux.SprintfBlue("Force overwrite of output files."))
	////rootCmd.PersistentFlags().BoolVarP(&CmdScribe.RemoveOutput, loadTools.FlagRemoveOutput, "", false, ux.SprintfBlue("Remove output file afterwards."))
	//rootCmd.PersistentFlags().BoolVarP(&CmdScribe.QuietProgress, loadTools.FlagQuiet, "q", false, ux.SprintfBlue("Silence progress in shell scripts."))
	////rootCmd.PersistentFlags().BoolVarP(&CmdScribe.DryRun, flagDryRun, "n", false, "Don't overwrite files.")
	//
	//rootCmd.PersistentFlags().BoolVarP(&CmdScribe.Debug, loadTools.FlagDebug ,"d", false, ux.SprintfBlue("DEBUG mode."))
	//
	////rootCmd.PersistentFlags().BoolVarP(&CmdScribe.HelpAll, loadTools.FlagHelpAll, "", false, ux.SprintfBlue("Show all help."))
	////rootCmd.PersistentFlags().BoolVarP(&CmdScribe.HelpVariables, loadTools.FlagHelpVariables, "", false, ux.SprintfBlue("Help on template variables."))
	////rootCmd.PersistentFlags().BoolVarP(&CmdScribe.HelpFunctions, loadTools.FlagHelpFunctions, "", false, ux.SprintfBlue("Help on template functions."))
	////rootCmd.PersistentFlags().BoolVarP(&CmdScribe.HelpExamples, loadTools.FlagHelpExamples, "", false, ux.SprintfBlue("Help on template examples."))
}


// initConfig reads in config file and ENV variables if set.
//func initConfig() {
//	if ConfigFile != "" {
//		// Use config file from the flag.
//		viper.SetConfigFile(ConfigFile)
//	} else {
//		// Find home directory.
//		//home, err := homedir.Dir()
//		//if err != nil {
//		//	fmt.Println(err)
//		//	os.Exit(1)
//		//}
//
//		viper.AddConfigPath(".")
//		viper.SetConfigName(defaults.BinaryName + "-config")
//	}
//
//	viper.AutomaticEnv() // read in environment variables that match
//
//	// If a config file is found, read it in.
//	if err := viper.ReadInConfig(); err == nil {
//		fmt.Println("Using config file:", viper.ConfigFileUsed())
//	}
//}


func SetCmd() {
	for range onlyOnce {
		if CmdScribe == nil {
			CmdScribe = loadTools.New(defaults.BinaryName, defaults.BinaryVersion, false)
			CmdScribe.Runtime.SetRepos(defaults.SourceRepo, defaults.BinaryRepo)
			if CmdScribe.State.IsNotOk() {
				break
			}

			// Import additional tools.
			CmdScribe.ImportTools(&deploywp.GetHelpers)
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
			CmdScribe.FlagSetDefault(loadTools.FlagTemplateFile, DefaultTemplateFile)
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


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() *ux.State {
	for range onlyOnce {
		//SetHelp(rootCmd)
		//SetCmd()

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
	astate := CmdScribe.State
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
