package cmd

import (
	"github.com/spf13/cobra"
	"github.com/newclarity/scribeHelpers/loadTools"
	"path/filepath"
)


func ProcessArgs(cmd *cobra.Command, args []string) *loadTools.ArgTemplate {
	var tmpl *loadTools.ArgTemplate
	// tmpl := loadTools.NewArgTemplate()

	for range OnlyOnce {
		tmpl = Cmd

		_ = tmpl.Exec.SetArgs(cmd.Use)
		_ = tmpl.Exec.AddArgs(args...)

		ext := ""
		if len(args) >= 1 {
			ext := filepath.Ext(args[0])
			if ext == ".json" {
				tmpl.Json.Name = args[0]
			} else if ext == ".tmpl" {
				tmpl.Template.Name = args[0]
			}
		}

		if len(args) >= 2 {
			ext = filepath.Ext(args[1])
			if ext == ".json" {
				tmpl.Json.Name = args[1]
			} else if ext == ".tmpl" {
				tmpl.Template.Name = args[1]
			}
		}

		tmpl.ValidateArgs()
		if tmpl.State.IsNotOk() {
			break
		}
	}

	return tmpl
}
