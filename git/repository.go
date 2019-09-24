package git

import (
	"errors"
	"fmt"
	"github.com/tsuyoshiwada/go-gitcmd"
	"github.com/wplib/deploywp/app"
	"github.com/wplib/deploywp/util"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
	"strings"
)

type Repository struct {
	Dir    Dir
	Url    Url
	Handle *git.Repository
	Client gitcmd.Client
	//commit *object.Commit
}

func NewRepository(d Dir) *Repository {
	r := Repository{
		Dir: d,
	}
	return &r
}

func (me *Repository) Exec(subcmd string, args ...string) (out string, err error) {
	cd := util.GetCurrentDir()
	err = os.Chdir(me.Dir)
	if err != nil {
		app.Fail("Cannot change directory to '%s'", me.Dir)
	}
	out, err = me.Client.Exec(subcmd, args...)
	err = os.Chdir(cd)
	if err != nil {
		app.Fail("Cannot restore directory to '%s'", cd)
	}
	return out, err
}

func (me *Repository) Open() (err error) {
	for range Once {
		me.Client = gitcmd.New(nil)
		err = me.Client.CanExec()
		if err != nil {
			err = fmt.Errorf("`git` does not exist or its command file is not executable: %s", err)
			break
		}
		out, err := me.Exec("rev-parse", "--is-inside-work-tree")
		if out != "true" {
			msg := "current directory does not contain a valid Git repository"
			if err != nil {
				msg = fmt.Sprintf("current directory does not contain a valid .Git repository: %s", err)
				break
			}
			err = errors.New(msg)
			break
		}
	}
	return err
}

func (me *Repository) Clone(u Url) (err error) {
	me.Url = u
	_, err = me.Exec("clone", u)
	return err
}

func (me *Repository) Branch() (branch string, err error) {
	return me.Exec("symbolic-ref", "--short", "HEAD")
}

func (me *Repository) Tags() (tags []string, err error) {
	var out string
	out, err = me.Exec("log", "-1", "--decorate=short", "--pretty=format:%D")
	tags = make([]string, 0)
	for _, t := range strings.Split(strings.TrimSpace(out), ",") {
		if t[:5] != " tag:" {
			continue
		}
		tags = append(tags, t[6:])
	}
	return tags, err
}

func (me *Repository) Commit() (commit *Commit, err error) {
	var out string
	out, err = me.Exec("rev-parse", "--verify", "HEAD")
	if err == nil {
		commit = NewCommit(out)
	}
	return commit, err
}

func (me *Repository) Status() (sts Status, err error) {
	for range Once {
		if me.Handle == nil {
			err = fmt.Errorf("repository not open")
			break
		}
		wt, err := me.Handle.Worktree()
		if err != nil {
			break
		}
		sts, err = wt.Status()
		if err != nil {
			break
		}
	}
	return sts, err
}

func (me *Repository) FilesChanged() (fps Filepaths, err error) {
	for range Once {
		var out string
		out, err = me.Exec("status", "--porcelain")
		if err != nil {
			break
		}
		sts := strings.Split(strings.TrimSpace(out), "\n")
		fps = make(Filepaths, len(sts))
		for i, fp := range sts {
			fps[i] = fp[3:]
		}
	}
	return fps, err
}

func (me *Repository) getHandle() (h *git.Repository, err error) {
	for range Once {
		h = me.Handle
		if h == nil {
			err = fmt.Errorf("repository handle is nil")
			break
		}
	}
	return h, err
}

func (me *Repository) GetTagObject(tag Tagname) (to *Tag, err error) {
	for range Once {
		var h *git.Repository
		h, err = me.getHandle()
		if err != nil {
			break
		}
		var t *Reference
		t, err = h.Tag(tag)
		if err != nil {
			break
		}
		to, err = h.TagObject(t.Hash())
		if err != nil {
			break
		}
	}
	if err != nil {
		err = fmt.Errorf("unable to access tag object '%s''; %s",
			tag,
			err.Error(),
		)
	}
	return to, err
}

func (me *Repository) Lock() (ok bool, err error) {
	for range Once {
		var to *object.Tag
		to, err = me.GetTagObject(LockTag)
		if err != nil {
			break
		}
		_ = to.ID()
	}
	return ok, err
}

func (me *Repository) Pull(opts ...*PullOptions) (err error) {
	if len(opts) == 0 {
		opts = []*PullOptions{}
	}
	for range Once {
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
	return err
}
