package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/app"
	"github.com/wplib/deploywp/deploywp"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	SuggestFor: []string{"deploy","run","now"},
	Short: "Deploy a WordPress website to a Pantheon site",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		app.Initialize()
	},
	Run: func(cmd *cobra.Command, args []string) {
		deploywp.NewDeploy(app.Config()).Run()
	},
}

func init() {
	RootCmd.AddCommand(RunCmd)
	//fs := RunCmd.Flags()
	//fs.StringVar(&app.Domain, "domain", "", "Domain to deploy")
}
