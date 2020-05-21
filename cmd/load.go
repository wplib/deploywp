package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/jsonTemplate"
	"github.com/wplib/deploywp/ux"
	"strings"
)


// releaseCmd represents the release command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: ux.SprintfBlue("Load and execute a template file."),
	Long: ux.SprintfBlue(`...`),
	Run: LoadTemplate,
}

func init() {
	rootCmd.AddCommand(loadCmd)
}

func LoadTemplate(cmd *cobra.Command, args []string) {
	for range OnlyOnce {
		var state ux.State
		var tmpl *jsonTemplate.Template

		/*
		Allow this to be used as a UNIX script.
		The following should be placed on the first line.
		#!/usr/bin/env deploywp load
		*/
		if len(args) > 0 {
			t := args[0]
			args = args[1:]
			_ = cmd.Flags().Set(argTemplateFile, t)

			t = strings.TrimSuffix(t, "tmpl") + "json"
			_ = cmd.Flags().Set(argJsonFile, t)
		}

		tmpl, state = ProcessArgs(cmd, args)
		if !state.IsOk() {
			state.PrintResponse()
			break
		}

		_ = tmpl.SetVersion(Version)

		state = tmpl.ProcessTemplate()
		if !state.IsOk() {
			state.PrintResponse()
			break
		}
	}
}