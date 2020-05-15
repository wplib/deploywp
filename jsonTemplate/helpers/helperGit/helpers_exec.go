package helperGit

import (
	"github.com/wplib/deploywp/jsonTemplate/helpers/helperTypes"
	"github.com/wplib/deploywp/only"
)


// Usage:
//		{{- $cmd := $git.GitClone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
// func (me *HelperGit) GitClone(args ...interface{}) *State {
func (me *HelperGit) GitClone(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandStatus, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitInit }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitInit(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandInit, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitAdd }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitAdd(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandAdd, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitMv }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitMv(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandMv, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitReset }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitReset(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandReset, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitRm }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitRm(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandRm, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitBisect }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitBisect(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandBisect, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitGrep }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitGrep(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandGrep, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitLog }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitLog(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandLog, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitShow }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitShow(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandShow, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitStatus }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitStatus(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandStatus, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitBranch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitBranch(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandBranch, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitCheckout }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitCheckout(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandCheckout, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitCommit }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitCommit(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandCommit, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitDiff }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitDiff(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandDiff, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitMerge }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitMerge(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandMerge, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitRebase }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitRebase(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandRebase, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitTag }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitTag(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandTag, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitFetch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitFetch(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = me.Exec(gitCommandFetch, args...)
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitPull }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitPull(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandPull, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}


// Usage:
//		{{- $cmd := $git.GitPush }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) GitPush(args ...string) *State {
	for range only.Once {
		if me.Reflect().IsNotOk() {
			break
		}

		me.Cmd = (*helperTypes.TypeExecCommand)(me.Exec(gitCommandPush, args...))
		if me.Cmd.IsError() {
			break
		}
	}

	return ReflectState(me.State)
}
