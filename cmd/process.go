package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/jsonTemplate"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)


// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "process",
	Short: "Process and release a template file",
	Long: `...`,
	Run: Process,
}

func init() {
	rootCmd.AddCommand(releaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// releaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// releaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Process(cmd *cobra.Command, args []string) {
	for range only.Once {
		var state ux.State
		var tmpl *jsonTemplate.Template

		tmpl, state = ProcessArgs(cmd, args)
		if !state.IsOk() {
			state.Print()
			break
		}

		_ = tmpl.SetVersion(Version)

		state = tmpl.ProcessTemplate()
		if !state.IsOk() {
			state.Print()
			break
		}
	}
}