package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/jsonTemplate"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)


// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: ux.SprintfBlue("Process and release a template file"),
	Long: ux.SprintfBlue(`...`),
	Run: Release,
}

func init() {
	rootCmd.AddCommand(releaseCmd)
}

func Release(cmd *cobra.Command, args []string) {
	for range only.Once {
		var state ux.State
		var tmpl *jsonTemplate.Template

		tmpl, state = ProcessArgs(cmd, args)
		if !state.IsOk() {
			fmt.Printf(state.PrintResponse())
			break
		}

		_ = tmpl.SetVersion(Version)

		state = tmpl.ProcessTemplate()
		if !state.IsOk() {
			fmt.Printf(state.PrintResponse())
			break
		}
	}
}