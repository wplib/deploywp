package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/jtc"
	"strings"
)


func ProcessArgs(cmd *cobra.Command, args []string) *jtc.ArgTemplate {
	tmpl := jtc.NewArgTemplate()

	for range OnlyOnce {
		var err error

		_ = tmpl.Exec.SetArgs(cmd.Use)
		_ = tmpl.Exec.AddArgs(args...)

		fl := cmd.Flags()

		//	flagConfigFile  = "config"
		//	flagVersion = "version"


		// Dry run mode.
		Cmd.DryRun, err = fl.GetBool(flagDryRun)
		if err != nil {
			tmpl.OverWrite = false
			tmpl.RemoveFiles = false
		}
		if Cmd.DryRun {
			tmpl.OverWrite = false
			tmpl.RemoveFiles = false
		} else {
			tmpl.OverWrite = true
			tmpl.RemoveFiles = true
		}


		// Json file.
		Cmd.JsonFile, err = fl.GetString(flagJsonFile)
		if err != nil {
			Cmd.JsonFile = defaultJsonFile
		}
		if Cmd.JsonFile == "" {
			Cmd.JsonFile = defaultJsonFile
		}
		tmpl.State = tmpl.SetJsonFile(Cmd.JsonFile)
		if tmpl.State.IsNotOk() {
			tmpl.State.SetError("ERROR: %s", err)
			break
		}
		Cmd.JsonFile = tmpl.JsonFile.GetPath()


		// Template file.
		for range OnlyOnce {
			Cmd.TemplateFile, err = fl.GetString(flagTemplateFile)
			if err != nil {
				Cmd.TemplateFile = defaultTemplateFile
			}
			if Cmd.TemplateFile == "" {
				Cmd.TemplateFile = defaultTemplateFile
			}

			tmpl.State = tmpl.SetTemplateFile(Cmd.TemplateFile)
			if tmpl.State.IsOk() {
				break
			}

			// Try again based on the json file.
			Cmd.TemplateFile = strings.TrimSuffix(tmpl.JsonFile.GetPath(), defaultJsonFileSuffix) + defaultTemplateFileSuffix
			tmpl.State = tmpl.SetTemplateFile(Cmd.TemplateFile)
			if tmpl.State.IsNotOk() {
				tmpl.State.SetError("ERROR: %s", err)
				break
			}
		}


		// Output file.
		Cmd.OutFile, err = fl.GetString(flagOutputFile)
		if err != nil {
			Cmd.OutFile = ""
		}
		if Cmd.OutFile == defaultOutFile {
			tmpl.OutFile = nil
		}
		tmpl.State = tmpl.SetOutFile(Cmd.OutFile)
		if tmpl.State.IsNotOk() {
			tmpl.State.SetError("ERROR: %s", err)
			break
		}


		// Chdir.
		Cmd.Chdir, err = fl.GetBool(flagChdir)
		if Cmd.Chdir {
			tmpl.State = tmpl.JsonFile.Chdir()
			if tmpl.State.IsNotOk() {
				//tmpl.State.SetError("ERROR: %s", err)
				break
			}
		}


		tmpl.SetValid()
		tmpl.State.SetOk("Processed arguments.")
	}

	return tmpl
}
