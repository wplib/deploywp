package helperGit

import (
	"fmt"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperPath"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"strings"
)


//type Repository struct {
//	Dir    Dir
//	Url    Url
//	Handle *git.Repository
//}


// Usage:
//		{{- $cmd := $git.Clone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
// func (me *HelperGit) Clone(url interface{}, dir ...interface{}) *TypeExecCommand {
func (me *HelperGit) Clone(url string, dir ...interface{}) *State {
	for range only.Once {
		if me._IsNil() {
			break
		}

		u := helperTypes.ReflectString(url)
		if u == nil {
			me.Cmd.SetError("URL is nil")
			break
		}
		me.SetUrl(*u)


		d := helperPath.ReflectAbsPath(dir...)
		if d == nil {
			me.Cmd.SetError("dir is nil")
			break
		}

		if !me.Base.SetPath(*d) {
			me.Cmd.SetError("error setting path to %s", me.Base.GetPath())
			break
		}

		me.Base.StatPath()
		if me.Base.Exists() {
			me.Cmd.SetError("cannot clone as path %s already exists", me.Base.GetPath())
			me.Cmd.Exit = 1
			break
		}


		ux.PrintfWhite("Cloning %s into %s\n", me.Url, me.Base.GetPath())
		me.skipDirCheck = true
		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandClone, me.Url, me.Base.GetPath()))
		me.skipDirCheck = false
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.Open }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) Open() *State {
	for range only.Once {
		if me._IsNil() {
			break
		}

		me.Cmd = (*TypeExecCommand)(me.Exec("rev-parse", "--is-inside-work-tree"))
		if me.Cmd.Output != "true" {
			if me.Cmd.IsError() {
				me.Cmd.SetError("current directory does not contain a valid .Git repository: %s", me.Cmd.ErrorValue)
				break
			}

			me.Cmd.SetError("current directory does not contain a valid Git repository")
			break
		}

		var err error
		me.repository, err = git.PlainOpen(me.Base.GetPath())
		if err != nil {
			me.Cmd.SetError(err)
			break
		}

		c, _ := me.repository.Config()
		me.Url = c.Remotes["origin"].URLs[0]

		me.Cmd.Output = fmt.Sprintf("Opened directory %s.\nRemote origin is set to %s\n", me.Base.GetPath(), me.Url)
		me.Cmd.Data = true
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.SetPath }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) SetPath(path ...interface{}) *State {
	for range only.Once {
		if me._IsNil() {
			break
		}

		p := helperPath.ReflectPath(path...)
		if p == nil {
			me.Cmd.SetError("path repo is nil")
			break
		}

		ps := helperPath.ResolveAbsPath(*p)
		if ps.IsError() {
			me.Cmd.ErrorValue = ps.ErrorValue
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

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GetUrl }}
//		{{- if $cmd.IsOk }}{{ $cmd.Data }}{{- end }}
func (me *HelperGit) GetUrl() *State {
	for range only.Once {
		if me._IsNil() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec("config", "--get", "remote.origin.url"))
		if me.Cmd.IsError() {
			break
		}

		me.Url = strings.TrimSpace(me.Cmd.Output)
		me.Cmd.Data = me.Url
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.SetUrl }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) SetUrl(u Url) *State {
	for range only.Once {
		if me._IsNil() {
			break
		}

		me.Url = u
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.Clone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
// func (me *HelperGit) Clone(url interface{}, dir ...interface{}) *TypeExecCommand {
func (me *HelperGit) Remove() *State {
	for range only.Once {
		if (*TypeGit)(me).IsNotOk() {
			break
		}

		me.State = me.Base.StatPath().Reflect()
		if me.State.IsError() {
			break
		}

		// @TODO - TO BE IMPLEMENTED

		//me.Base.Exists
		//ps := helperSystem.ResolveAbsPath(*d)
		//if ps.IsFile {
		//	break
		//}
		//if ps.IsDir {
		//	if ps.Exists {
		//		me.Cmd.SetError("Repository exists for directory '%s'", ps.Path)
		//		me.Cmd.Exit = 1
		//		break
		//	}
		//}
		//
		//me.SetUrl(*u)
		//me.Base = ps
		//ux.PrintfWhite("Cloning %s into %s\n", me.Url, me.Base.Path)
		//
		//me.skipDirCheck = true
		//me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandClone, me.Url, me.Base.Path))
		//me.skipDirCheck = false
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.Lock }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) Lock() *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = me.GetTagObject(LockTag)
		if me.Cmd.IsError() {
			break
		}

		var to *object.Tag
		to = me.Cmd.Data.(*object.Tag)

		_ = to.ID()
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GetStatus }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GetStatus() (sts Status, err error) {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		var wt *git.Worktree
		wt, me.Cmd.ErrorValue = me.repository.Worktree()
		if me.Cmd.IsError() {
			break
		}

		sts, me.Cmd.ErrorValue = wt.Status()
		if me.Cmd.IsError() {
			break
		}
	}

	return sts, err
}
