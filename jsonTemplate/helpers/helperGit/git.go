package helperGit

import (
	"github.com/tsuyoshiwada/go-gitcmd"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperFile"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

//var _ helperTypes.TypeOsPathGetter = (*TypeOsPath)(nil)
//type TypeOsPath helperTypes.TypeOsPath

var _ helperTypes.TypeExecCommandGetter = (*TypeExecCommand)(nil)
type TypeExecCommand helperTypes.TypeExecCommand


type TypeGit struct {
	Url string
	Base *helperFile.TypeOsPath
	GitConfig *gitcmd.Config
	GitOptions []string
	skipDirCheck bool

	client gitcmd.Client
	repository *git.Repository

	Cmd *helperTypes.TypeExecCommand
}


// Usage:
//		{{ $git := GitLogin }}
func HelperGitLogin(path ...interface{}) *TypeGit {
	var ret TypeGit

	for range only.Once {
		ret.Cmd = ret.Cmd.EnsureNotNil()
		ret.Cmd = ret.SetPath(path...)
		ret.client = gitcmd.New(ret.GitConfig)

		ret.Cmd = ret.IsNil()
		if ret.Cmd.IsError() {
			break
		}

		ret.Cmd = ret.IsExec()
		if ret.Cmd.IsError() {
			break
		}
	}

	return &ret
}


// Usage:
//		{{ $cmd := $git.Chdir .Some.Directory }}
//		{{ if $git.IsOk }}Changed to directory {{ $git.Dir }}{{ end }}
func (me *TypeGit) Chdir(dir ...interface{}) *helperTypes.TypeOsPath {
	return helperSystem.HelperChdir(dir...)
}


// Usage:
//		{{ $git.SetDryRun }}
func (me *TypeGit) SetDryRun() bool {
	me.GitOptions = append(me.GitOptions, "-n")
	return true
}


// Usage:
//		{{- $cmd := $git.GetStatus }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GetStatus() (sts Status, err error) {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
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


// Usage:
//		{{- $cmd := $git.Lock }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) Lock() *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
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

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.IsNil }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) IsNil() *helperTypes.TypeExecCommand {
	for range only.Once {
		if me == nil {
			me.Cmd.SetError("`git` client is not configured")
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.IsExec }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) IsExec() *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd.ErrorValue = me.client.CanExec()
		if me.Cmd.IsError() {
			me.Cmd.SetError("`git` does not exist or its command file is not executable: %s", me.Cmd.ErrorValue)
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.IsNilRepository }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) IsNilRepository() *helperTypes.TypeExecCommand {
	for range only.Once {
		if me.repository == nil {
			me.Cmd.SetError("repository not open")
		}
	}

	return me.Cmd
}


// Usage:
//		{{ if $ret.IsError }}{{ $cmd.PrintError }}{{ end }}
func (me *TypeGit) SetError(format interface{}, a ...interface{}) {
	me.Cmd.SetError(format, a...)
}


// Usage:
//		{{ if $ret.IsError }}{{ $cmd.PrintError }}{{ end }}
func (me *TypeGit) IsError() bool {
	return me.Cmd.IsError()
}


// Usage:
//		{{ if $ret.IsOk }}OK{{ end }}
func (me *TypeGit) IsOk() bool {
	return me.Cmd.IsOk()
}


// Usage:
//		{{ if $ret.IsOk }}OK{{ end }}
func (me *TypeGit) PrintError() string {
	return me.Cmd.PrintError()
}


// Usage:
//		{{ if $ret.IsOk }}OK{{ end }}
func (me *TypeGit) ExitOnError() string {
	return me.Cmd.ExitOnError()
}
