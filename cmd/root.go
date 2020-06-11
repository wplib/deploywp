package cmd

import (
	"fmt"
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wplib/deploywp/defaults"
)


type TypeCmd struct {
	ConfigFile   string

	JsonFile     string
	TemplateFile string
	OutFile      string

	Debug        bool
	Chdir        bool
	DryRun       bool

	State        *ux.State
}

const DefaultJsonFile     = "deploywp.json"
//const DefaultTemplateFile = `{{ $dwp := LoadDeployWp .Json (.Exec.GetArg 1) }}{{ $dwp.ExitOnError }}{{ $dwp.Run }}`
const DefaultTemplateFile = `{{ $dwp := LoadDeployWp .Json .Exec.Args }}{{ $dwp.ExitOnError }}{{ $state := $dwp.Run }}{{ $state.ExitOnError }}`

var Cmd *loadTools.TypeScribeArgs
var ConfigFile string
const 	flagConfigFile  	= "config"
func SetCmd() {
	if Cmd == nil {
		Cmd = loadTools.New(defaults.BinaryName, defaults.BinaryVersion, false)
		Cmd.Runtime.SetRepos(defaults.SourceRepo, defaults.BinaryRepo)
	}
}


func init() {
	SetCmd()

	//Cmd = loadTools.NewArgTemplate(defaults.BinaryName, defaults.BinaryVersion)

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&ConfigFile, flagConfigFile, fmt.Sprintf("%s-config.json", defaults.BinaryName), ux.SprintfBlue("%s: config file.", defaults.BinaryName))
	_ = rootCmd.PersistentFlags().MarkHidden(flagConfigFile)

	rootCmd.PersistentFlags().StringVarP(&Cmd.Json.Filename, loadTools.FlagJsonFile, "j", DefaultJsonFile, ux.SprintfBlue("Alternative JSON file."))
	rootCmd.PersistentFlags().StringVarP(&Cmd.Template.Filename, loadTools.FlagTemplateFile, "t", DefaultTemplateFile, ux.SprintfBlue("Alternative template file."))
	rootCmd.PersistentFlags().StringVarP(&Cmd.Output.Filename, loadTools.FlagOutputFile, "o", loadTools.DefaultOutFile, ux.SprintfBlue("Output file."))
	rootCmd.PersistentFlags().StringVarP(&Cmd.WorkingPath.Filename, loadTools.FlagWorkingPath, "p", loadTools.DefaultWorkingPath, ux.SprintfBlue("Set working path."))

	rootCmd.PersistentFlags().BoolVarP(&Cmd.Chdir, loadTools.FlagChdir, "c", false, ux.SprintfBlue("Change to directory containing %s", DefaultJsonFile))
	//rootCmd.PersistentFlags().BoolVarP(&Cmd.RemoveTemplate, loadTools.FlagRemoveTemplate, "", false, ux.SprintfBlue("Remove template file afterwards."))
	//rootCmd.PersistentFlags().BoolVarP(&Cmd.ForceOverwrite, loadTools.FlagForce, "f", false, ux.SprintfBlue("Force overwrite of output files."))
	//rootCmd.PersistentFlags().BoolVarP(&Cmd.RemoveOutput, loadTools.FlagRemoveOutput, "", false, ux.SprintfBlue("Remove output file afterwards."))
	rootCmd.PersistentFlags().BoolVarP(&Cmd.QuietProgress, loadTools.FlagQuiet, "q", false, ux.SprintfBlue("Silence progress in shell scripts."))
	//rootCmd.PersistentFlags().BoolVarP(&Cmd.DryRun, flagDryRun, "n", false, "Don't overwrite files.")

	rootCmd.PersistentFlags().BoolVarP(&Cmd.Debug, loadTools.FlagDebug ,"d", false, ux.SprintfBlue("DEBUG mode."))

	//rootCmd.PersistentFlags().BoolVarP(&Cmd.HelpAll, loadTools.FlagHelpAll, "", false, ux.SprintfBlue("Show all help."))
	//rootCmd.PersistentFlags().BoolVarP(&Cmd.HelpVariables, loadTools.FlagHelpVariables, "", false, ux.SprintfBlue("Help on template variables."))
	//rootCmd.PersistentFlags().BoolVarP(&Cmd.HelpFunctions, loadTools.FlagHelpFunctions, "", false, ux.SprintfBlue("Help on template functions."))
	//rootCmd.PersistentFlags().BoolVarP(&Cmd.HelpExamples, loadTools.FlagHelpExamples, "", false, ux.SprintfBlue("Help on template examples."))

	rootCmd.Flags().BoolP(loadTools.FlagVersion, "v", false, ux.SprintfBlue("Display version of " + defaults.BinaryName))
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if ConfigFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(ConfigFile)
	} else {
		// Find home directory.
		//home, err := homedir.Dir()
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}

		viper.AddConfigPath(".")
		viper.SetConfigName(defaults.BinaryName + "-config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   defaults.BinaryName,
	Short: "Pantheon release tool.",
	Long: `...`,
	Run: gbRootFunc,
}

func gbRootFunc(cmd *cobra.Command, args []string) {
	//Cmd.State = Cmd.State.EnsureNotNil()

	for range onlyOnce {
		var ok bool
		fl := cmd.Flags()
		//tmpl := loadTools.NewArgTemplate(defaults.BinaryName, defaults.BinaryVersion)

		// ////////////////////////////////
		// Show version.
		ok, _ = fl.GetBool(loadTools.FlagVersion)
		if ok {
			VersionShow()
			Cmd.State.Clear()
			break
		}

		// Show HelpVariables.
		ok, _ = fl.GetBool(loadTools.FlagHelpVariables)
		if ok {
			HelpVariables()
			break
		}

		// Show HelpFunctions.
		ok, _ = fl.GetBool(loadTools.FlagHelpFunctions)
		if ok {
			HelpFunctions()
			break
		}

		// Show HelpExamples.
		ok, _ = fl.GetBool(loadTools.FlagHelpExamples)
		if ok {
			HelpExamples()
			break
		}

		// Show all help.
		ok, _ = fl.GetBool(loadTools.FlagHelpAll)
		if ok {
			HelpAll()
			break
		}


		Cmd.State = ProcessArgs(Cmd, cmd, args)
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			_ = cmd.Help()
			break
		}

		// Show help if no commands specified.
		if len(args) == 0 {
			_ = cmd.Help()
			Cmd.State.Clear()
			break
		}
	}
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() *ux.State {
	for range onlyOnce {
		SetHelp(rootCmd)
		SetCmd()

		//if Cmd == nil {
		//	Cmd = loadTools.NewArgTemplate(defaults.BinaryName, defaults.BinaryVersion)
		//}

		err := rootCmd.Execute()
		if err != nil {
			Cmd.State.SetError(err)
			break
		}
	}

	return Cmd.State
}
