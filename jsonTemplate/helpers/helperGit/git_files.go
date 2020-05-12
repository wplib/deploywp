package helperGit

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"strings"
)


func (me *TypeGit) ChangedFiles() *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec("status", "--porcelain")
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

	return me.Cmd
}


func (me *TypeGit) AddFiles() *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec("add", "--porcelain")
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

	return me.Cmd
}
