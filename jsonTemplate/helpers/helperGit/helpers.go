package helperGit

import (
	"github.com/tsuyoshiwada/go-gitcmd"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperExec"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperPath"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)

type HelperGit TypeGit
func (me *HelperGit) Reflect() *TypeGit {
	return (*TypeGit)(me)
}
func (me *TypeGit) Reflect() *HelperGit {
	return (*HelperGit)(me)
}


// Usage:
//		{{ $git := NewGit }}
func HelperNewGit(path ...interface{}) *HelperGit {
	var ret TypeGit

	for range only.Once {
		p := helperPath.ReflectAbsPath(path...)
		if p == nil {
			break
		}
		if ret.Base.SetPath(*p) {
			ret.State = ret.Base.StatPath().Reflect()
			if ret.Base.Exists() {

			}
			if ret.State.IsError() {
				break
			}

			// Can now set it after.
			//ret.State.SetError("%s destination empty", *p)
			//break
		}

		ret.Cmd = helperExec.NewExecCommand()
		ret.client = gitcmd.New(ret.GitConfig)

		if ret.IsNil() {
			break
		}

		if ret.IsNotAvailable() {
			break
		}
	}

	return ReflectHelperGit(&ret)
}


// Usage:
//		{{ $cmd := $git.Chdir .Some.Directory }}
//		{{ if $git.IsOk }}Changed to directory {{ $git.Dir }}{{ end }}
func (me *HelperGit) Chdir(dir ...interface{}) *helperPath.HelperOsPath {
	return helperPath.HelperChdir(dir...)
}


// Usage:
//		{{ $git.SetDryRun }}
func (me *HelperGit) SetDryRun() bool {
	me.GitOptions = append(me.GitOptions, "-n")
	return true
}


// Usage:
//		{{ $cmd := $git.Exec "tag" "-l" }}
//		{{ if $git.IsOk }}OK{{ end }}
// func (me *HelperGit) Exec(cmd interface{}, args ...interface{}) *State {
func (me *HelperGit) Exec(cmd string, args ...string) *State {
	for range only.Once {
		if me.Reflect().IsAvailable() {
			break
		}

		//c := helperTypes.ReflectString(cmd)
		//if c == nil {
		//	break
		//}
		//me.Cmd.Exe = *c
		//
		//a := helperTypes.ReflectStrings(args...)
		//if a == nil {
		//	break
		//}
		//
		//me.Cmd.Args = []string{}
		//me.Cmd.Args = append(me.Cmd.Args, me.GitOptions...)
		//me.Cmd.Args = append(me.Cmd.Args, *a...)

		me.Cmd.Exe = cmd
		me.Cmd.Args = []string{}
		me.Cmd.Args = append(me.Cmd.Args, me.GitOptions...)
		me.Cmd.Args = append(me.Cmd.Args, args...)

		for range only.Once {
			if me.skipDirCheck {
				break
			}
			if me.Base.IsCwd() {
				break
			}
			path := me.Base.GetPath()
			cd := helperPath.HelperChdir(path)
			if cd.State.IsError() {
				ux.PrintfError("Cannot change directory to '%s'", path)
				break
			}
		}

		var err error
		me.Cmd.Output, err = me.client.Exec(me.Cmd.Exe, me.Cmd.Args...)
		if err != nil {
			me.State.SetError(err)
			me.Cmd.Exit = 1	// Fake an exit code.
			break
		}

		me.State.SetOk("")
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.IsExec }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) IsAvailable() *State {
	for range only.Once {
		if me.Reflect().IsAvailable() {
			break
		}
	}

	//foo := &State{}
	//foo = (*State)(me.State)
	//foo = (*State)(me.Reflect().State)
	//foo = ReflectState(me.State)

	return ReflectState(me.State)
}


// Usage:
//		{{ if $ret.IsError }}{{ $cmd.PrintError }}{{ end }}
func (me *HelperGit) SetError(error ...interface{}) {
	me.State.SetError(error...)
}


// Usage:
//		{{ if $ret.IsError }}{{ $cmd.PrintError }}{{ end }}
func (me *HelperGit) IsError() bool {
	return me.State.IsError()
}


// Usage:
//		{{ if $ret.IsOk }}OK{{ end }}
func (me *HelperGit) IsOk() bool {
	return me.State.IsOk()
}


// Usage:
//		{{ if $ret.IsOk }}OK{{ end }}
func (me *HelperGit) PrintError() string {
	return me.Cmd.PrintError()
}


// Usage:
//		{{ if $ret.IsOk }}OK{{ end }}
func (me *HelperGit) ExitOnError() string {
	return me.Cmd.ExitOnError()
}
