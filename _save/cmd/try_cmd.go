package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/app"
)

var TryCmd = &cobra.Command{
	Use:        "try",
	SuggestFor: []string{"test"},
	Short:      "A place to try things",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		app.Initialize()
	},
	Run: func(cmd *cobra.Command, args []string) {
		MakeDeployWP().Try()
	},
}

func init() {
	RootCmd.AddCommand(TryCmd)
	//fs := TryCmd.Flags()
	//fs.StringVar(&app.Domain, "domain", "", "Domain to deploy")
}
