package helperGit

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
)

type TypeGitClone struct {
	Base *helperSystem.TypeOsPath
	Cmd *helperTypes.TypeExecCommand
}


func (me *TypeGit) Clone(url interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		value := helperTypes.ReflectString(url)
		if value == nil {
			me.Cmd.SetError("URL is nil")
			break
		}

		me.Url = *value
		me.Cmd = me.Exec("clone", me.Url)
	}

	return me.Cmd
}


func (me *TypeGit) SetPath(path ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}

		p := helperTypes.ReflectPath(path...)
		if p == nil {
			me.Cmd.SetError("path repo is nil")
			break
		}

		ps := helperSystem.ResolveAbsPath(*p)
		if ps.IsError() {
			me.Cmd.Error = ps.Error
			break
		}
		//if !ps.Exists {
		//	me.Exe.SetError("path not found")
		//	break
		//}
		if ps.IsFile {
			me.Cmd.SetError("path is not a directory")
			break
		}

		me.Base = ps
	}

	return me.Cmd
}


func (me *TypeGit) SetUrl(u Url) {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}

		me.Url = u
	}
}
