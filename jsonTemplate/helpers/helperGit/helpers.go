package helperGit

import (
	"github.com/tsuyoshiwada/go-gitcmd"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperExec"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperPath"
	"github.com/wplib/deploywp/ux"
)

type HelperGit TypeGit
func (g *HelperGit) Reflect() *TypeGit {
	return (*TypeGit)(g)
}
func (g *TypeGit) Reflect() *HelperGit {
	return (*HelperGit)(g)
}

func (c *HelperGit) IsNil() *ux.State {
	if state := ux.IfNilReturnError(c); state.IsError() {
		return state
	}
	c.State = c.State.EnsureNotNil()
	return c.State
}


// Usage:
//		{{ $git := NewGit }}
func HelperNewGit(path ...interface{}) *HelperGit {
	ret := NewGit()

	for range OnlyOnce {
		p := helperPath.ReflectAbsPath(path...)
		if p == nil {
			break
		}
		if ret.Base.SetPath(*p) {
			state := ret.Base.StatPath()
			ret.State.SetState(state)
			if ret.Base.Exists() {

			}
			if ret.State.IsError() {
				break
			}

			// Can now set it after.
			//ret.State.SetError("%s destination empty", *p)
			//break
		}

		ret.Cmd = helperExec.NewExecCommand(false)
		ret.client = gitcmd.New(ret.GitConfig)

		if ret.IsNotAvailable() {
			break
		}
	}

	return ReflectHelperGit(ret)
}


// Usage:
//		{{ $cmd := $git.Chdir .Some.Directory }}
//		{{ if $git.IsOk }}Changed to directory {{ $git.Dir }}{{ end }}
func (g *HelperGit) Chdir() *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	return helperPath.HelperChdir(g.Base.GetPath()).State
}


// Usage:
//		{{ $git.SetDryRun }}
func (g *HelperGit) SetDryRun() bool {
	g.GitOptions = append(g.GitOptions, "-n")
	return true
}


// Usage:
//		{{ $cmd := $git.Exec "tag" "-l" }}
//		{{ if $git.IsOk }}OK{{ end }}
// func (me *HelperGit) Exec(cmd interface{}, args ...interface{}) *ux.State {
func (g *HelperGit) Exec(cmd string, args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		if g.Reflect().IsNotAvailable() {
			break
		}

		//c := helperTypes.ReflectString(cmd)
		//if c == nil {
		//	break
		//}
		//g.Cmd.Exe = *c
		//
		//a := helperTypes.ReflectStrings(args...)
		//if a == nil {
		//	break
		//}
		//
		//g.Cmd.Args = []string{}
		//g.Cmd.Args = append(g.Cmd.Args, g.GitOptions...)
		//g.Cmd.Args = append(g.Cmd.Args, *a...)

		g.Cmd.Exe = cmd
		g.Cmd.Args = []string{}
		g.Cmd.Args = append(g.Cmd.Args, g.GitOptions...)
		g.Cmd.Args = append(g.Cmd.Args, args...)

		for range OnlyOnce {
			if g.skipDirCheck {
				break
			}
			if g.Base.IsCwd() {
				break
			}
			path := g.Base.GetPath()
			cd := helperPath.HelperChdir(path)
			if cd.State.IsError() {
				ux.PrintfError("Cannot change directory to '%s'", path)
				break
			}
		}

		out, err := g.client.Exec(g.Cmd.Exe, g.Cmd.Args...)
		g.State.SetOutput(out)
		g.State.OutputTrim()
		g.State.SetError(err)
		if g.State.IsError() {
			g.State.SetExitCode(1) // Fake an exit code.
			break
		}

		g.State.SetOk("")
	}

	return g.State
}


//// Usage:
////		{{- $cmd := $git.IsExec }}
////		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
//func (g *HelperGit) IsAvailable() *ux.State {
//	for range OnlyOnce {
//		if g.Reflect().IsNotAvailable() {
//			break
//		}
//	}
//
//	//foo := &State{}
//	//foo = (*State)(g.State)
//	//foo = (*State)(g.Reflect().State)
//	//foo = ReflectState(g.State)
//
//	return g.State
//}
//
//
//// Usage:
////		{{ if $ret.IsError }}{{ $cmd.PrintError }}{{ end }}
//func (g *HelperGit) SetError(error ...interface{}) {
//	g.State.SetError(error...)
//}
//
//
//// Usage:
////		{{ if $ret.IsError }}{{ $cmd.PrintError }}{{ end }}
//func (g *HelperGit) IsError() bool {
//	return g.State.IsError()
//}
//
//
//// Usage:
////		{{ if $ret.IsOk }}OK{{ end }}
//func (g *HelperGit) IsOk() bool {
//	return g.State.IsOk()
//}
//
//
//// Usage:
////		{{ if $ret.IsOk }}OK{{ end }}
//func (g *HelperGit) PrintError() string {
//	return g.Cmd.PrintError()
//}
//
//
//// Usage:
////		{{ if $ret.IsOk }}OK{{ end }}
//func (g *HelperGit) ExitOnError() string {
//	return g.State.ExitOnError()
//}
