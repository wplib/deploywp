package helperGit

import (
	"github.com/wplib/deploywp/ux"
	"gopkg.in/src-d/go-git.v4"
)


// Usage:
//		{{- $cmd := $git.Open }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (g *HelperGit) Pull(opts ...*PullOptions) *ux.State {
	if state := g.IsNil(); state.IsError() {
		return state
	}
	g.State.SetFunction("")

	for range OnlyOnce {
		if len(opts) == 0 {
			opts = []*PullOptions{}
		}

		var wt *git.Worktree
		var err error
		wt, err = g.repository.Worktree()
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}

		err = wt.Pull(opts[0])
		g.State.SetError(err)
		if g.State.IsError() {
			break
		}
	}

	if g.State.IsError() {
		g.State.SetError("Pull() - %s", g.State.GetError())
	}
	return g.State
}
