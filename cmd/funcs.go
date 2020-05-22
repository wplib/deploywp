package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/jsonTemplate"
	"strings"
)


func ProcessArgs(cmd *cobra.Command, args []string) *jsonTemplate.ArgTemplate {
	tmpl := jsonTemplate.NewArgTemplate()

	for range OnlyOnce {
		var err error
		var s string
		var b bool

		_ = tmpl.Exec.SetFullArgs(cmd.Use)
		_ = tmpl.Exec.AddFullArgs(args...)
		_ = tmpl.Exec.SetArgs(tmpl.Exec.GetFullArgs()[1:]...)

		fl := cmd.Flags()

		//	flagConfigFile  = "config"
		//	flagVersion = "version"


		// Dry run mode.
		b, err = fl.GetBool(flagDryRun)
		if err != nil {
			tmpl.OverWrite = false
			tmpl.RemoveFiles = false
		}
		if b {
			tmpl.OverWrite = false
			tmpl.RemoveFiles = false
		} else {
			tmpl.OverWrite = true
			tmpl.RemoveFiles = true
		}


		// Json file.
		s, err = fl.GetString(flagJsonFile)
		if err != nil {
			s = defaultJsonFile
		}
		tmpl.State = tmpl.SetJsonFile(s)
		if tmpl.State.IsNotOk() {
			tmpl.State.SetError("ERROR: %s", err)
			break
		}


		// Template file.
		for range OnlyOnce {
			s, err = fl.GetString(flagTemplateFile)
			if err != nil {
				s = defaultTemplateFile
			}

			tmpl.State = tmpl.SetTemplateFile(s)
			if tmpl.State.IsOk() {
				break
			}

			// Try again based on the json file.
			s = strings.TrimSuffix(tmpl.JsonFile.GetPath(), defaultJsonFileSuffix) + defaultTemplateFileSuffix
			tmpl.State = tmpl.SetTemplateFile(s)
			if tmpl.State.IsNotOk() {
				tmpl.State.SetError("ERROR: %s", err)
				break
			}
		}


		// Output file.
		s, err = fl.GetString(flagOutputFile)
		if err != nil {
			s = ""
		}
		if s == "-" {
			tmpl.OutFile = nil
		}
		tmpl.State = tmpl.SetOutFile(s)
		if tmpl.State.IsNotOk() {
			tmpl.State.SetError("ERROR: %s", err)
			break
		}


		// Chdir.
		b, err = fl.GetBool(flagChdir)
		if b {
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
