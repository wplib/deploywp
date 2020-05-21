package helperGit

import (
	"github.com/wplib/deploywp/ux"
)


// Usage:
//		{{- $cmd := $git.GitClone }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
// func (me *HelperGit) GitClone(args ...interface{}) *ux.State {
func (g *HelperGit) GitClone(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandClone, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitInit }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitInit(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandInit, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitAdd }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitAdd(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandAdd, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitMv }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitMv(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandMv, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitReset }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitReset(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandReset, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitRm }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitRm(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandRm, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitBisect }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitBisect(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandBisect, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitGrep }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitGrep(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandGrep, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitLog }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitLog(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandLog, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitShow }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitShow(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandShow, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitStatus }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitStatus(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandStatus, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitBranch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitBranch(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandBranch, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitCheckout }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitCheckout(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandCheckout, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitCommit }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitCommit(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandCommit, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitDiff }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitDiff(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandDiff, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitMerge }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitMerge(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandMerge, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitRebase }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitRebase(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandRebase, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitTag }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitTag(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandTag, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitFetch }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitFetch(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandFetch, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitPull }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitPull(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandPull, args...))
	return g.State
}


// Usage:
//		{{- $cmd := $git.GitPush }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) GitPush(args ...string) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")
	g.State.SetState(g.Exec(gitCommandPush, args...))
	return g.State
}
