package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/app"
	"github.com/wplib/deploywp/util"
)

var RootCmd = &cobra.Command{
	Use:   "deploywp",
	Short: "Deploy a WordPress site to Pantheon.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return err
	},
}

func init() {
	pf := RootCmd.PersistentFlags()
	//pf.BoolVarP(&app.NoCache, "no-cache", "", false, "Disable caching")

	pf.StringVarP(&app.ConfigFile,
		"config",
		"",
		app.Config().GetConfigFile(),
		"Filepath to a config.json to load for deploying",
	)


	pf.StringVarP(&app.DeployDir,
		"deploy",
		"",
		util.GetCurrentDir(),
		fmt.Sprintf("Directory to deploy; e.g. containing '%s'",app.DeployFile),
	)
}

