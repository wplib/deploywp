package cmd

import (
)


//func ProcessArgs(toolArgs *loadTools.TypeScribeArgs, cmd *cobra.Command, args []string) *ux.State {
//	state := CmdScribe.State
//
//	for range onlyOnce {
//		state = toolArgs.ImportTools(&deploywp.GetHelpers)
//		if state.IsNotOk() {
//			break
//		}
//
//		state = toolArgs.ValidateArgs()
//		if state.IsNotOk() {
//			break
//		}
//
//
//		//err := toolArgs.Runtime.SetArgs(cmd.Use)
//		//if err != nil {
//		//	state.SetError(err)
//		//	break
//		//}
//		//
//		//err = toolArgs.Runtime.AddArgs(args...)
//		//if err != nil {
//		//	state.SetError(err)
//		//	break
//		//}
//		//
//		//for range onlyTwice {
//		//	if len(args) >= 1 {
//		//		ext := filepath.Ext(args[0])
//		//		if ext == ".json" {
//		//			toolArgs.Json.Filename = args[0]
//		//			args = args[1:]
//		//		} else if ext == ".tmpl" {
//		//			toolArgs.Template.Filename = args[0]
//		//			args = args[1:]
//		//		} else {
//		//			break
//		//		}
//		//	}
//		//}
//		//_ = CmdScribe.Runtime.SetArgs(args...)
//		//
//		//state = toolArgs.ImportTools(&deploywp.GetHelpers)
//		//if state.IsNotOk() {
//		//	break
//		//}
//		//
//		//state = toolArgs.ValidateArgs()
//		//if state.IsNotOk() {
//		//	break
//		//}
//	}
//
//	return state
//}


//func cmdTools(cmd *cobra.Command, args []string) {
//	state := CmdScribe.State
//
//	for range onlyOnce {
//		state = CmdScribe.ProcessArgs(cmd.Use, args)
//		// Ignore errors as there's no args.
//
//		CmdScribe.PrintTools()
//		state.Clear()
//	}
//
//	CmdScribe.State = state
//}
//
//
//func cmdConvert(cmd *cobra.Command, args []string) {
//	state := CmdScribe.State
//
//	for range onlyOnce {
//		CmdScribe.RemoveTemplate = true
//		CmdScribe.Output.File = loadTools.CmdConvert
//
//		state = CmdScribe.ProcessArgs(cmd.Use, args)
//		if state.IsNotOk() {
//			state.PrintResponse()
//			break
//		}
//
//		state = CmdScribe.Load()
//		if state.IsNotOk() {
//			state.PrintResponse()
//			break
//		}
//
//		//ux.PrintflnOk("Converting file '%s' => '%s'", CmdScribe.Template.GetPath(), CmdScribe.Output.GetPath())
//		state = CmdScribe.Run()
//		state.PrintResponse()
//	}
//
//	CmdScribe.State = state
//}
//
//
//func cmdLoad(cmd *cobra.Command, args []string) {
//	state := CmdScribe.State
//
//	for range onlyOnce {
//		state = CmdScribe.ProcessArgs(cmd.Use, args)
//		if state.IsNotOk() {
//			state.PrintResponse()
//			break
//		}
//
//		state = CmdScribe.Load()
//		if state.IsNotOk() {
//			state.PrintResponse()
//			break
//		}
//
//		//ux.PrintflnOk("Loading template '%s' and saving result to '%s'", CmdScribe.Template.Filename, CmdScribe.Output.Filename)
//		state = CmdScribe.Run()
//		state.PrintResponse()
//	}
//
//	CmdScribe.State = state
//}
//
//
//func cmdRun(cmd *cobra.Command, args []string) {
//	state := CmdScribe.State
//
//	for range onlyOnce {
//		CmdScribe.ExecShell = true
//		CmdScribe.Output.File = loadTools.SelectConvert
//		CmdScribe.StripHashBang = true
//
//		/*
//			Allow this to be used as a UNIX script.
//			The following should be placed on the first line.
//			#!/usr/bin/env scribe load
//		*/
//
//		state = CmdScribe.ProcessArgs(cmd.Use, args)
//		if state.IsNotOk() {
//			state.PrintResponse()
//			break
//		}
//
//		//ux.PrintflnOk("Executing file '%s' => '%s'", CmdScribe.Template.Filename, CmdScribe.Output.Filename)
//		state = CmdScribe.Load()
//		if state.IsNotOk() {
//			state.PrintResponse()
//			break
//		}
//
//		//ux.PrintflnOk("Loading file '%s' => '%s'", CmdScribe.Template.Filename, CmdScribe.Output.Filename)
//		state = CmdScribe.Run()
//		state.PrintResponse()
//	}
//
//	CmdScribe.State = state
//}
