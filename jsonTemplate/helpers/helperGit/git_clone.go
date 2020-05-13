package helperGit

import (
	"fmt"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"gopkg.in/src-d/go-git.v4"
	"strings"
)

type TypeGitClone struct {
	Base *helperTypes.TypeOsPath
	Cmd *helperTypes.TypeExecCommand
}


// Usage:
//		{{- $cmd := $git.Clone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) Clone(url interface{}, dir ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		//me.Cmd = me.IsNilRepository()
		//if me.Cmd.IsOk() {
		//	break
		//}

		u := helperTypes.ReflectString(url)
		if u == nil {
			me.Cmd.SetError("URL is nil")
			break
		}
		me.SetUrl(*u)

		d := helperTypes.ReflectPath(dir...)
		if d == nil {
			me.Cmd.SetError("URL is nil")
			break
		}
		ps := helperSystem.ResolveAbsPath(*d)
		if ps.IsFile {
			break
		}
		if ps.IsDir {
			if ps.Exists {
				me.Cmd.SetError("Repository exists for directory '%s'", me.Base.Path)
				me.Cmd.Exit = 1
				break
			}
		}
		me.Base = ps

		ux.PrintfWhite("Cloning %s into %s\n", me.Url, me.Base.Path)

		me.skipDirCheck = true
		me.Cmd = me.Exec(gitCommandClone, me.Url, me.Base.Path)
		me.skipDirCheck = false
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.Open }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) Open() *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec("rev-parse", "--is-inside-work-tree")
		if me.Cmd.Output != "true" {
			if me.Cmd.IsError() {
				me.Cmd.SetError("current directory does not contain a valid .Git repository: %s", me.Cmd.ErrorValue)
				break
			}

			me.Cmd.SetError("current directory does not contain a valid Git repository")
			break
		}

		var err error
		me.repository, err = git.PlainOpen(me.Base.Path)
		if err != nil {
			me.Cmd.SetError(err)
			break
		}

		c, _ := me.repository.Config()
		me.Url = c.Remotes["origin"].URLs[0]

		me.Cmd.Output = fmt.Sprintf("Opened directory %s.\nRemote origin is set to %s\n", me.Base.Path, me.Url)
		me.Cmd.Data = true
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.SetPath }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
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

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GetUrl }}
//		{{- if $cmd.IsOk }}{{ $cmd.Data }}{{- end }}
func (me *TypeGit) GetUrl() *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec("config", "--get", "remote.origin.url")
		if me.Cmd.IsError() {
			break
		}

		me.Url = strings.TrimSpace(me.Cmd.Output)
		me.Cmd.Data = me.Url
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.SetUrl }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) SetUrl(u Url) {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}

		me.Url = u
	}
}
