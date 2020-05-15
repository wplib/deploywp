package helperGit

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"strings"
)


// Usage:
//		{{- $cmd := $git.ChangedFiles }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) ChangedFiles() *State {
	for range only.Once {
		if (*TypeGit)(me).IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec("status", "--porcelain"))
		if me.Cmd.IsError() {
			break
		}

		var fps Filepaths
		sts := strings.Split(strings.TrimSpace(me.Cmd.Output), "\n")
		fps = make(Filepaths, len(sts))
		for i, fp := range sts {
			s := strings.Fields(fp)
			if len(s) > 1 {
				fps[i] = s[1]
			}
		}

		me.Cmd.Data = fps
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.AddFiles }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) AddFiles() *State {
	for range only.Once {
		if (*TypeGit)(me).IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec("add", "--porcelain"))
		if me.Cmd.IsError() {
			break
		}

		var fps Filepaths
		sts := strings.Split(strings.TrimSpace(me.Cmd.Output), "\n")
		fps = make(Filepaths, len(sts))
		for i, fp := range sts {
			s := strings.Fields(fp)
			if len(s) > 1 {
				fps[i] = s[1]
			}
		}

		me.Cmd.Data = fps
	}

	return ReflectState(me.State)
}
