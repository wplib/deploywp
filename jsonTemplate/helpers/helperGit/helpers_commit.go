package helperGit

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
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
func (me *HelperGit) Commit(format interface{}, a ...interface{}) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec("rev-parse", "--verify", "HEAD"))
		if me.Cmd.IsError() {
			break
		}

		me.Cmd.Data = _NewCommit(me.Cmd.Output)
	}

	return ReflectState(me.State)
}
