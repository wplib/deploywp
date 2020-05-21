package helperGit

import (
	"github.com/wplib/deploywp/ux"
)

type Commit struct {
	Hash string
}


func _NewCommit(hash string) *Commit {
	return &Commit{
		Hash: hash,
	}
}


// Usage:
//		{{- $cmd := $git.Commit }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) Commit(format interface{}, a ...interface{}) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		g.State.SetState(g.Exec("rev-parse", "--verify", "HEAD"))
		if g.State.IsError() {
			break
		}

		g.State.Response = _NewCommit(g.State.Output)
	}

	return g.State
}
