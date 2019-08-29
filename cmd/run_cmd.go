package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/app"
	"github.com/wplib/deploywp/deploywp"
	"github.com/wplib/deploywp/jsonfile"
)

var RunCmd = &cobra.Command{
	Use:        "run",
	SuggestFor: []string{"deploy", "run", "now"},
	Short:      "DeployWP a WordPress website to a Pantheon site",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		app.Initialize()
	},
	Run: func(cmd *cobra.Command, args []string) {
		c := app.Config()
		jf := jsonfile.NewJsonFile(*c)
		deploywp.NewDeployWP(c).Run(jf)
	},
}

func init() {
	RootCmd.AddCommand(RunCmd)
	//fs := RunCmd.Flags()
	//fs.StringVar(&app.Domain, "domain", "", "Domain to deploy")
}
