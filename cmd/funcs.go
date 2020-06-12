package cmd

import (
	"github.com/newclarity/scribeHelpers/loadTools"
	"github.com/newclarity/scribeHelpers/ux"
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/deploywp"
	"path/filepath"
)


func ProcessArgs(toolArgs *loadTools.TypeScribeArgs, cmd *cobra.Command, args []string) *ux.State {
	state := Cmd.State

	for range onlyOnce {
		err := toolArgs.Runtime.SetArgs(cmd.Use)
		if err != nil {
			state.SetError(err)
			break
		}

		err = toolArgs.Runtime.AddArgs(args...)
		if err != nil {
			state.SetError(err)
			break
		}

		for range onlyTwice {
			if len(args) >= 1 {
				ext := filepath.Ext(args[0])
				if ext == ".json" {
					toolArgs.Json.Filename = args[0]
					args = args[1:]
				} else if ext == ".tmpl" {
					toolArgs.Template.Filename = args[0]
					args = args[1:]
				} else {
					break
				}
			}
		}
		_ = Cmd.Runtime.SetArgs(args...)

		state = toolArgs.ImportTools(&deploywp.GetHelpers)
		if state.IsNotOk() {
			break
		}

		state = toolArgs.ValidateArgs()
		if state.IsNotOk() {
			break
		}
	}

	return state
}


func cmdBuild(cmd *cobra.Command, args []string) {
	state := Cmd.State

	for range onlyOnce {
		Cmd.Chdir = true	// In this mode we always change directory to the JSON file.

		state = ProcessArgs(Cmd, cmd, args)
		if state.IsNotOk() {
			state.PrintResponse()
			break
		}

		state = Cmd.Load()
		if state.IsNotOk() {
			state.PrintResponse()
			break
		}

		ux.PrintflnOk("Building website via deploywp.")
		state = Cmd.Run()

		//dwp := deploywp.HelperLoadDeployWp(Cmd.JsonStruct.Json, Cmd.Runtime.GetArgs()...)
		//if dwp.State.IsNotOk() {
		//	dwp.State.PrintResponse()
		//	break
		//}
		//dwp.Runtime = Cmd.JsonStruct.Exec
		//Cmd.State = dwp.Run()
		//if Cmd.State.IsNotOk() {
		//	Cmd.State.SetExitCode(1)
		//	//Cmd.State.Exit(1)
		//	break
		//}

		state.PrintResponse()
		ux.PrintflnBlue("\n# FINISHED")
	}

	Cmd.State = state
}


func cmdTools(cmd *cobra.Command, args []string) {
	state := Cmd.State

	for range onlyOnce {
		state = ProcessArgs(Cmd, rootCmd, args)
		// Ignore errors as there's no args.

		Cmd.PrintTools()
		state.Clear()
	}

	Cmd.State = state
}


func cmdConvert(cmd *cobra.Command, args []string) {
	state := Cmd.State

	for range onlyOnce {
		Cmd.RemoveTemplate = true
		Cmd.Output.Filename = loadTools.CmdConvert

		state = ProcessArgs(Cmd, cmd, args)
		if state.IsNotOk() {
			state.PrintResponse()
			break
		}

		state = Cmd.Load()
		if state.IsNotOk() {
			state.PrintResponse()
			break
		}

		ux.PrintflnOk("Converting file '%s' => '%s'", Cmd.Template.Filename, Cmd.Output.Filename)
		state = Cmd.Run()
		state.PrintResponse()
	}

	Cmd.State = state
}


func cmdLoad(cmd *cobra.Command, args []string) {
	state := Cmd.State

	for range onlyOnce {
		state = ProcessArgs(Cmd, cmd, args)
		if state.IsNotOk() {
			state.PrintResponse()
			break
		}

		state = Cmd.Load()
		if state.IsNotOk() {
			state.PrintResponse()
			break
		}

		ux.PrintflnOk("Loading template '%s' and saving result to '%s'", Cmd.Template.Filename, Cmd.Output.Filename)
		state = Cmd.Run()
		state.PrintResponse()
	}

	Cmd.State = state
}


func cmdRun(cmd *cobra.Command, args []string) {
	state := Cmd.State

	for range onlyOnce {
		Cmd.ExecShell = true
		Cmd.Output.Filename = loadTools.SelectConvert
		Cmd.StripHashBang = true

		/*
			Allow this to be used as a UNIX script.
			The following should be placed on the first line.
			#!/usr/bin/env scribe load
		*/

		state = ProcessArgs(Cmd, cmd, args)
		if state.IsNotOk() {
			state.PrintResponse()
			break
		}

		ux.PrintflnOk("Executing file '%s' => '%s'", Cmd.Template.Filename, Cmd.Output.Filename)
		state = Cmd.Load()
		if state.IsNotOk() {
			state.PrintResponse()
			break
		}

		ux.PrintflnOk("Loading file '%s' => '%s'", Cmd.Template.Filename, Cmd.Output.Filename)
		state = Cmd.Run()
		state.PrintResponse()
	}

	Cmd.State = state
}
