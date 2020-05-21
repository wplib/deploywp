package helperGit

import (
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"strings"
)


// Usage:
//		{{- $cmd := $git.ChangedFiles }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) ChangedFiles() *ux.State {
	for range OnlyOnce {
		g.State.SetFunction("")

		if g.Reflect().IsNotOk() {
			break
		}

		g.State.SetState(g.Exec(gitCommandStatus, "--porcelain"))
		if g.State.IsError() {
			break
		}

		var fps Filepaths
		fps = make(Filepaths, len(g.State.OutputArray))
		for i, fp := range g.State.OutputArray {
			s := strings.Fields(fp)
			if len(s) > 1 {
				fps[i] = s[1]
			}
		}

		g.State.Response = fps
	}

	return g.State
}


// Usage:
//		{{- $cmd := $git.AddFiles }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) AddFiles() *ux.State {
	for range OnlyOnce {
		g.State.SetFunction("")

		if g.Reflect().IsNotOk() {
			break
		}

		g.State.SetState(g.Exec(gitCommandAdd, "--porcelain"))
		if g.State.IsError() {
			break
		}

		var fps Filepaths
		fps = make(Filepaths, len(g.State.OutputArray))
		for i, fp := range g.State.OutputArray {
			s := strings.Fields(fp)
			if len(s) > 1 {
				fps[i] = s[1]
			}
		}

		g.State.Response = fps
	}

	return g.State
}
