package helperGit

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperSystem"
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
	"github.com/wplib/deploywp/ux"
)


// Usage:
//		{{ $cmd := $git.Exec "tag" "-l" }}
//		{{ if $git.IsOk }}OK{{ end }}
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

		me.Cmd.Args = []string{}
		me.Cmd.Args = append(me.Cmd.Args, me.GitOptions...)
		me.Cmd.Args = append(me.Cmd.Args, *a...)

		//me.Exe = me.NonHelperExec(ret.Exe, ret.Args...)

		cwd := helperSystem.HelperGetwd()
		if cwd.Path != me.Base.Path {
			cd := helperSystem.HelperChdir(me.Base.Path)
			if cd.IsError() {
				ux.PrintfError("Cannot change directory to '%s'", me.Base.Path)
				break
			}
		}

		me.Cmd.Output, me.Cmd.ErrorValue = me.client.Exec(me.Cmd.Exe, me.Cmd.Args...)
		if me.Cmd.ErrorValue != nil {
			me.Cmd.Exit = 1	// Fake an exit code.
		}

		//cd = helperSystem.HelperChdir(me.Base.Path)
		//if cd.IsError() {
		//	ux.PrintfError("Cannot change back to directory '%s'", cwd.Path)
		//	break
		//}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitClone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitClone(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandStatus, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitInit }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitInit(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandInit, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitAdd }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitAdd(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandAdd, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitMv }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitMv(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandMv, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitReset }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitReset(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandReset, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitRm }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitRm(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandRm, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitBisect }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitBisect(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandBisect, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitGrep }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitGrep(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandGrep, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitLog }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitLog(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandLog, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitShow }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitShow(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandShow, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitStatus }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitStatus(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandStatus, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitBranch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitBranch(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandBranch, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitCheckout }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitCheckout(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandCheckout, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitCommit }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitCommit(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandCommit, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitDiff }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitDiff(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandDiff, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitMerge }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitMerge(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandMerge, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitRebase }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitRebase(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandRebase, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitTag }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitTag(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandTag, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitFetch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitFetch(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandFetch, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitPull }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitPull(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandPull, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}


// Usage:
//		{{- $cmd := $git.GitPush }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *TypeGit) GitPush(args ...interface{}) *helperTypes.TypeExecCommand {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
			break
		}

		me.Cmd = me.Exec(gitCommandPush, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return me.Cmd
}
