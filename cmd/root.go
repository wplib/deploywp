package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wplib/deploywp/defaults"
	"github.com/wplib/deploywp/ux"
	"strings"
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

var Cmd TypeCmd


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
		var err error
		var ok bool

		fl := cmd.Flags()

		// ////////////////////////////////
		// Modifier flags.

		// Quiet flag.
		ok, err = fl.GetBool(flagQuiet)
		if err != nil {
			Cmd.State.DebugSet(false)
		} else {
			Cmd.State.DebugSet(ok)
		}

		// DEBUG flag.
		ok, err = fl.GetBool(flagDebug)
		if err != nil {
			Cmd.State.DebugSet(false)
		} else {
			Cmd.State.DebugSet(ok)
		}
		if ok {
			flargs := fl.Args()
			ux.Printf("flargs: %s\n", strings.Join(flargs, " "))
			ux.Printf("args: %s\n", strings.Join(args, " "))
		}


		// ////////////////////////////////
		// Flag versions of commands.

		// Show version.
		ok, _ = fl.GetBool(flagVersion)
		if ok {
			Version(cmd, args)
			Cmd.State.Clear()
			break
			//os.Exit(0)
		}

		// Show help if no commands specified.
		if len(args) == 0 {
			Version(cmd, args)
			_ = cmd.Help()
			Cmd.State.Clear()
			break
			//os.Exit(0)
		}
	}
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() *ux.State {
	for range OnlyOnce {
		var err error

		Cmd = TypeCmd {
			State:        ux.NewState(false),
			ConfigFile:   "",
			JsonFile:     "",
			TemplateFile: "",
			Debug:        false,
		}

		err = rootCmd.Execute()
		if err != nil {
			Cmd.State.SetError(err)
			break
		}
	}

	return Cmd.State
}


func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&Cmd.ConfigFile, flagConfigFile, "", defaults.BinaryName + " config file, (default is deploywp-config.json).")

	rootCmd.PersistentFlags().StringVar(&Cmd.JsonFile, flagJsonFile, defaultJsonFile, defaults.BinaryName + " JSON file.")
	rootCmd.PersistentFlags().StringVar(&Cmd.TemplateFile, flagTemplateFile, defaultTemplateFile, defaults.BinaryName + " template file.")
	rootCmd.PersistentFlags().StringVar(&Cmd.OutFile, flagOutputFile, defaultOutFile, defaults.BinaryName + " output file.")

	rootCmd.PersistentFlags().BoolVarP(&Cmd.Debug, flagDebug ,"d", false, "DEBUG mode.")
	rootCmd.PersistentFlags().BoolVarP(&Cmd.Chdir, flagChdir, "c", false, "Change to directory containing deploywp.json.")
	rootCmd.PersistentFlags().BoolVarP(&Cmd.DryRun, flagDryRun, "n", false, "Don't overwrite files.")

	rootCmd.Flags().BoolP(flagVersion, "v", false, ux.SprintfBlue("Display version of " + defaults.BinaryName))
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if Cmd.ConfigFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(Cmd.ConfigFile)
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
