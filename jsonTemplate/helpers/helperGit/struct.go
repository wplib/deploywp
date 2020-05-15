package helperGit

import (
	"github.com/tsuyoshiwada/go-gitcmd"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperExec"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperPath"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

var _ helperExec.TypeExecCommandGetter = (*TypeExecCommand)(nil)
type TypeExecCommand helperExec.TypeExecCommand


type TypeGit struct {
	State        *ux.State

	Url          string
	Base         *helperPath.TypeOsPath

	GitConfig    *gitcmd.Config
	GitOptions   []string

	skipDirCheck bool

	client       gitcmd.Client
	repository   *git.Repository

	Cmd          *helperExec.TypeExecCommand
}


type State ux.State
func (p *State) Reflect() *ux.State {
	return (*ux.State)(p)
}
func ReflectState(p *ux.State) *State {
	return (*State)(p)
}
func ReflectHelperGit(p *TypeGit) *HelperGit {
	return (*HelperGit)(p)
}


func (me *TypeGit) IsOk() bool {
	var ok bool

	for range only.Once {
		if me.IsNil() {
			break
		}
		if me.IsAvailable() {
			break
		}
		if me.IsNilRepository() {
			break
		}
		me.State.Clear()
		ok = true
	}

	return ok
}
func (me *TypeGit) IsNotOk() bool {
	return !me.IsOk()
}

func (me *TypeGit) IsNil() bool {
	var ok bool

	for range only.Once {
		if me == nil {
			me.Cmd.SetError("`git` client is not configured")
			break
		}
		me.State.Clear()
		ok = true
	}

	return ok
}


func (me *TypeGit) IsNilRepository() bool {
	var ok bool

	for range only.Once {
		if me.IsNil() {
			break
		}
		if me.repository == nil {
			me.Cmd.SetError("repository not open")
			break
		}
		me.State.Clear()
		ok = true
	}

	return ok
}


func (me *TypeGit) IsAvailable() bool {
	var ok bool

	for range only.Once {
		if me.IsNil() {
			break
		}
		me.Cmd.ErrorValue = me.client.CanExec()
		if me.Cmd.IsError() {
			me.Cmd.SetError("`git` does not exist or its command file is not executable: %s", me.Cmd.ErrorValue)
			break
		}
		me.State.Clear()
		ok = true
	}

	return ok
}
func (me *TypeGit) IsNotAvailable() bool {
	return !me.IsAvailable()
}


type (
	Dir          = string
	Url          = string
	Filepath     = string
	Filepaths    []Filepath
	ReadableName = string
	Tagname      = string
)

type (
	PullOptions  = git.PullOptions
	LogOptions   = git.LogOptions
	Tag          = object.Tag
	Reference    = plumbing.Reference
	CloneOptions = git.CloneOptions
	Status       = git.Status
)
