package helperGit

import (
	"github.com/wplib/deploywp/only"
	"github.com/newclarity/scribeHelpers/ux"
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
	for range OnlyOnce {
		g.State.SetFunction("")

		if g.Reflect().IsNotOk() {
			break
		}

		g.State.SetState(g.Exec("rev-parse", "--verify", "HEAD"))
		if g.State.IsError() {
			break
		}

		g.State.Response = _NewCommit(g.State.Output)
	}

	return g.State
}
