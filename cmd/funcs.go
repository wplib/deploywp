package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wplib/deploywp/jsonTemplate"
	"github.com/wplib/deploywp/ux"
	"os"
	"path/filepath"
	"strings"
)


func ProcessArgs(cmd *cobra.Command, args []string) (*jsonTemplate.Template, ux.State) {
	var tmpl jsonTemplate.Template
	var state ux.State

	for range OnlyOnce {
		var err error
		var s string
		var b bool

		_ = tmpl.SetFullArgs(cmd.Use)
		_ = tmpl.AddFullArgs(args...)
		_ = tmpl.SetArgs(tmpl.GetFullArgs()[1:]...)

		fl := cmd.Flags()

		s, err = fl.GetString(argJsonFile)
		if err != nil {
			s = defaultJsonFile
		}
		err = tmpl.SetJsonFile(s)
		if err != nil {
			state.SetError("ERROR: %s", err)
			break
		}


		b, err = fl.GetBool(argChdir)
		if b {
			dir := tmpl.GetJsonFile()
			dir = filepath.Dir(dir)
			err = os.Chdir(dir)
			if err != nil {
				state.SetError("ERROR: %s", err)
				break
			}
		}


		for range OnlyOnce {
			s, err = fl.GetString(argTemplateFile)
			if err != nil {
				s = defaultTemplateFile
			}

			err = tmpl.SetTemplateFile(s)
			if err == nil {
				break
			}

			// Try again based on the json file.
			s = strings.TrimSuffix(tmpl.GetJsonFile(), defaultJsonFileSuffix) + defaultTemplateFileSuffix
			err = tmpl.SetTemplateFile(s)
			if err != nil {
				state.SetError("ERROR: %s", err)
				break
			}
		}


		err = tmpl.SetValid()
		if err != nil {
			state.SetError("ERROR: %s", err)
			break
		}

		state.SetOk("Processed arguments.")
	}

	return &tmpl, state
}
