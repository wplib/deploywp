package helperGit

import (
	"github.com/wplib/deploywp/only"
	"gopkg.in/src-d/go-git.v4"
)


func (me *TypeGit) Pull(opts ...*PullOptions) (err error) {
	for range only.Once {
		me.Cmd = me.IsNil()
		if me.Cmd.IsError() {
			break
		}
		me.Cmd = me.IsNilRepository()
		if me.Cmd.IsError() {
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

	return err
}
