package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/ux"
	"strings"
)


func init() {
	rootCmd.AddCommand(loadCmd)
}

// releaseCmd represents the release command
var loadCmd = &cobra.Command{
	Use:   cmdLoad,
	Short: ux.SprintfBlue("Load and execute a template file."),
	Long: ux.SprintfBlue(`...`),
	Run: LoadTemplate,
}
func LoadTemplate(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		//Cmd.State = ux.NewState(Cmd.Debug)
		//var tmpl *jsonTemplate.ArgTemplate

		/*
		Allow this to be used as a UNIX script.
		The following should be placed on the first line.
		#!/usr/bin/env deploywp load
		*/
		if len(args) > 0 {
			t := args[0]
			args = args[1:]
			_ = cmd.Flags().Set(flagJsonFile, t)

			t = strings.TrimSuffix(t, "json") + "tmpl"
			_ = cmd.Flags().Set(flagTemplateFile, t)
		}

		tmpl := ProcessArgs(rootCmd, args)
		Cmd.State = tmpl.State
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		Cmd.State = tmpl.LoadTemplate()
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}

		Cmd.State = tmpl.RunTemplate()
		if Cmd.State.IsNotOk() {
			Cmd.State.PrintResponse()
			break
		}
	}
}