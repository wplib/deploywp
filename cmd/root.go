package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string
var DeployWpJsonFile string
var DeployWpTemplateFile string
var Debug bool


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "deploywp",
	Short: "Pantheon release tool.",
	Long: `...`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, argConfigFile, "", "deploywp config file, (default is deploywp-config.json).")
	rootCmd.PersistentFlags().StringVar(&DeployWpJsonFile, argJsonFile, defaultJsonFile, "deploywp JSON file.")
	rootCmd.PersistentFlags().StringVar(&DeployWpTemplateFile, argTemplateFile, defaultTemplateFile, "deploywp template file.")
	rootCmd.PersistentFlags().BoolVarP(&Debug, argDebug ,"d", false, "DEBUG mode.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("debug", "d", false, "DEBUG mode.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		//home, err := homedir.Dir()
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}

		viper.AddConfigPath(".")
		viper.SetConfigName("deploywp-config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
