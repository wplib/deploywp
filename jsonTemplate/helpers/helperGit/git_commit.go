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
func (me *TypeGit) Commit(format interface{}, a ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec("rev-parse", "--verify", "HEAD")
		if me.Cmd.IsError() {
			break
		}

		me.Cmd.Data = _NewCommit(me.Cmd.Output)
	}

	return me.Cmd
}
