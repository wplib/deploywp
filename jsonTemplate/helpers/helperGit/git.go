package helperGit

import (
	"github.com/tsuyoshiwada/go-gitcmd"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

//var _ helperTypes.TypeOsPathGetter = (*TypeOsPath)(nil)
//type TypeOsPath helperTypes.TypeOsPath

//var _ helperTypes.TypeExecCommandGetter = (*TypeExecCommand)(nil)
//type TypeExecCommand helperTypes.TypeExecCommand


type TypeGit struct {
	Url string
	Base *helperSystem.TypeOsPath
	GitConfig *gitcmd.Config
	GitOptions []string

	client gitcmd.Client
	repository *git.Repository

	Cmd *helperTypes.TypeExecCommand
}


//// Usage:
////		{{ $git := GitLogin }}
////		{{ if $git.IsError }}FAILED{{ end }}
//func (me *TypeGit) IsError() bool {
//	return me.Exe.Error.IsError()
//}
//
//
//// Usage:
////		{{ $git := GitLogin }}
////		{{ if $git.IsOk }}OK{{ end }}
//func (me *TypeGit) IsOk() bool {
//	return me.Exe.Error.IsOk()
//}


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
//		{{ $ret := $git.Chdir .Some.Directory }}
//		{{ if $ret.IsOk }}Changed to directory {{ $ret.Dir }}{{ end }}
func (me *TypeGit) Chdir(dir interface{}) *helperSystem.TypeOsPath {
	return helperSystem.HelperChdir(dir)
}


func (me *TypeGit) SetDryRun() bool {
	me.GitOptions = append(me.GitOptions, "-n")
	return true
}


func (me *TypeGit) Exec(cmd interface{}, args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}

		c := helperTypes.ReflectString(cmd)
		if c == nil {
			break
		}
		me.Cmd.Exe = *c

		a := helperTypes.ReflectStrings(args...)
		if a == nil {
			break
		}
		me.Cmd.Args = append(me.Cmd.Args, me.GitOptions...)
		me.Cmd.Args = append(me.Cmd.Args, *a...)

		//me.Exe = me.NonHelperExec(ret.Exe, ret.Args...)

		cwd := helperSystem.HelperGetwd()
		cd := helperSystem.HelperChdir(me.Base.Path)
		if cd.IsError() {
			ux.PrintfError("Cannot change directory to '%s'", me.Base.Path)
			break
		}

		me.Cmd.Output, me.Cmd.Error = me.client.Exec(me.Cmd.Exe, me.Cmd.Args...)
		if me.Cmd.Error != nil {
			me.Cmd.Exit = 1	// Fake an exit code.
		}

		cd = helperSystem.HelperChdir(me.Base.Path)
		if cd.IsError() {
			ux.PrintfError("Cannot change back to directory '%s'", cwd.Path)
			break
		}
	}

	return me.Cmd
}


// @TODO
func (me *TypeGit) Open() *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec("rev-parse", "--is-inside-work-tree")
		if me.Cmd.Output != "true" {
			if me.Cmd.Error != nil {
				me.Cmd.SetError("current directory does not contain a valid .Git repository: %s", me.Cmd.Error)
				break
			}

			me.Cmd.SetError("current directory does not contain a valid Git repository")
			break
		}

		me.repository, me.Cmd.Error = git.PlainOpen(me.Base.Path)
		if me.Cmd.IsError() {
			break
		}

		me.Cmd.Data = true
	}

	return me.Cmd
}


// @TODO
func (me *TypeGit) Status() (sts Status, err error) {
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
		wt, me.Cmd.Error = me.repository.Worktree()
		if me.Cmd.IsError() {
			break
		}

		sts, me.Cmd.Error = wt.Status()
		if me.Cmd.IsError() {
			break
		}
	}

	return sts, err
}


func (me *TypeGit) Lock() (ok bool, err error) {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		var to *object.Tag
		to, err = me.GetTagObject(LockTag)
		if err != nil {
			break
		}
		_ = to.ID()
	}

	return ok, err
}


func (me *TypeGit) IsNil() *helperTypes.TypeExecCommand {
	for range only.Once {
		if me == nil {
			me.Cmd.SetError("`git` client is not configured")
			break
		}
	}

	return me.Cmd
}


func (me *TypeGit) IsExec() *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd.Error = me.client.CanExec()
		if me.Cmd.Error != nil {
			me.Cmd.SetError("`git` does not exist or its command file is not executable: %s", me.Cmd.Error)
			break
		}
	}

	return me.Cmd
}


func (me *TypeGit) IsNilRepository() *helperTypes.TypeExecCommand {
	for range only.Once {
		if me.repository == nil {
			me.Cmd.SetError("repository not open")
		}
	}

	return me.Cmd
}


// Usage:
//		{{ if $ret.IsError }}ERROR{{ end }}
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
