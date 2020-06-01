package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/defaults"
	"github.com/newclarity/JsonToConfig/jtc"
	"github.com/newclarity/JsonToConfig/ux"
)


func init() {
	rootCmd.AddCommand(versionCmd)
}


var versionCmd = &cobra.Command{
	Use:   jtc.CmdVersion,
	Short: ux.SprintfBlue("Show version of binary."),
	Long:  ux.SprintfBlue(`...`),
	Run:   Version,
}
func Version(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		//Cmd.State = ux.NewState(Cmd.Debug)
		fmt.Printf("%s %s\n",
			ux.SprintfBlue(defaults.BinaryName),
			ux.SprintfCyan("v%s", defaults.BinaryVersion),
			)
		Cmd.State.Clear()
	}
}
