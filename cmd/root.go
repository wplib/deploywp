package cmd

import (
	"fmt"
	"github.com/newclarity/JsonToConfig/jtc"
	"github.com/newclarity/JsonToConfig/ux"
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

var Cmd *jtc.ArgTemplate
var ConfigFile string
const flagConfigFile      = "config"
const DefaultJsonFile     = "deploywp.json"
const DefaultTemplateFile = `{{ $dwp := LoadDeployWp .Json }}{{ $dwp.ExitOnError }}{{ $dwp.Run }}`


func init() {
	Cmd = jtc.NewArgTemplate(defaults.BinaryName, defaults.BinaryVersion)

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&ConfigFile, flagConfigFile, fmt.Sprintf("%s-config.json", defaults.BinaryName), ux.SprintfBlue("%s: config file.", defaults.BinaryName))
	_ = rootCmd.PersistentFlags().MarkHidden(flagConfigFile)

	rootCmd.PersistentFlags().StringVarP(&Cmd.Json.Name, jtc.FlagJsonFile, "j", DefaultJsonFile, ux.SprintfBlue("Alternative JSON file."))
	rootCmd.PersistentFlags().StringVarP(&Cmd.Template.Name, jtc.FlagTemplateFile, "t", DefaultTemplateFile, ux.SprintfBlue("Alternative template file."))
	rootCmd.PersistentFlags().StringVarP(&Cmd.Output.Name, jtc.FlagOutputFile, "o", jtc.DefaultOutFile, ux.SprintfBlue("Output file."))

	rootCmd.PersistentFlags().BoolVarP(&Cmd.Chdir, jtc.FlagChdir, "c", false, ux.SprintfBlue("Change to directory containing %s", DefaultJsonFile))
	rootCmd.PersistentFlags().BoolVarP(&Cmd.QuietProgress, jtc.FlagQuiet, "q", false, ux.SprintfBlue("Silence progress in shell scripts."))
	//rootCmd.PersistentFlags().BoolVarP(&Cmd.DryRun, flagDryRun, "n", false, "Don't overwrite files.")

	rootCmd.PersistentFlags().BoolVarP(&Cmd.Debug, jtc.FlagDebug ,"d", false, ux.SprintfBlue("DEBUG mode."))

	rootCmd.Flags().BoolP(jtc.FlagVersion, "v", false, ux.SprintfBlue("Display version of " + defaults.BinaryName))
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
	Cmd.State = Cmd.State.EnsureNotNil()

	for range OnlyOnce {
		var ok bool
		fl := cmd.Flags()
		tmpl := jtc.NewArgTemplate(defaults.BinaryName, defaults.BinaryVersion)

		// ////////////////////////////////
		// Show version.
		ok, _ = fl.GetBool(jtc.FlagVersion)
		if ok {
			Version(cmd, args)
			Cmd.State.Clear()
			break
		}


		tmpl = ProcessArgs(cmd, args)
		Cmd.State = tmpl.State
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			_ = cmd.Help()
			break
		}

		// Show help if no commands specified.
		if len(args) == 0 {
			Version(cmd, args)
			_ = cmd.Help()
			Cmd.State.Clear()
			break
		}
	}
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() *ux.State {
	for range OnlyOnce {
		var err error

		if Cmd == nil {
			Cmd = jtc.NewArgTemplate(defaults.BinaryName, defaults.BinaryVersion)
		}

		err = rootCmd.Execute()
		if err != nil {
			Cmd.State.SetError(err)
			break
		}
	}

	return Cmd.State
}
