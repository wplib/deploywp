package helperGit

import (
	"github.com/wplib/deploywp/only"
	"gopkg.in/src-d/go-git.v4"
)


// Usage:
//		{{- $cmd := $git.Open }}
//		{{- if $cmd.IsError }}{{ $cmd.PrintError }}{{- end }}
func (me *HelperGit) Pull(opts ...*PullOptions) *State {
	for range only.Once {
		if (*TypeGit)(me).IsNotOk() {
			break
		}

		if len(opts) == 0 {
			opts = []*PullOptions{}
		}

		for range only.Once {
			var h *git.Repository
			h, err = me.getHandle()
			if err != nil {
				break
			}
			wt, err := h.Worktree()
			if err != nil {
				break
			}
			err = wt.Pull(opts[0])
		}
	}

	return ReflectState(me.State)
}
