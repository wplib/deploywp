package helperGit

import (
	"gopkg.in/src-d/go-git.v4"
)


type Repository struct {
	Dir    Dir
	Url    Url
	Handle *git.Repository
}


//func HelperNewRepository(d Dir) *Repository {
//	r := Repository{
//		Dir: d,
//	}
//	return &r
//}
//
//
//func (me *TypeGit) Open() *helperTypes.TypeExecCommand {
//	for range only.Once {
//		me.Client = gitcmd.New(nil)
//		me.Exe.Error = me.Client.CanExec()
//		if me.Exe.Error != nil {
//			me.Exe.SetError("`git` does not exist or its command file is not executable: %s", me.Exe.Error)
//			break
//		}
//
//		me.Exe = me.Exec("rev-parse", "--is-inside-work-tree")
//		if me.Exe.Output != "true" {
//			if me.Exe.Error != nil {
//				me.Exe.SetError("current directory does not contain a valid .Git repository: %s", me.Exe.Error)
//				break
//			}
//
//			me.Exe.SetError("current directory does not contain a valid Git repository")
//			break
//		}
//	}
//
//	return me.Exe
//}
//
//
//func (me *Repository) Clone(u Url) (err error) {
//	me.Url = u
//	_, err = me.Exec("clone", u)
//	return err
//}
//
//func (me *Repository) Branch() (branch string, err error) {
//	return me.Exec("symbolic-ref", "--short", "HEAD")
//}
//
//func (me *Repository) Tags() (tags []string, err error) {
//	var out string
//	out, err = me.Exec("log", "-1", "--decorate=short", "--pretty=format:%D")
//	tags = make([]string, 0)
//	for _, t := range strings.Split(strings.TrimSpace(out), ",") {
//		if t[:5] != " tag:" {
//			continue
//		}
//		tags = append(tags, t[6:])
//	}
//	return tags, err
//}
//
//func (me *Repository) Commit() (commit *Commit, err error) {
//	var out string
//	out, err = me.Exec("rev-parse", "--verify", "HEAD")
//	if err == nil {
//		commit = NewCommit(out)
//	}
//	return commit, err
//}
//
//func (me *Repository) Status() (sts Status, err error) {
//	for range only.Once {
//		if me.Handle == nil {
//			err = fmt.Errorf("repository not open")
//			break
//		}
//		wt, err := me.Handle.Worktree()
//		if err != nil {
//			break
//		}
//		sts, err = wt.Status()
//		if err != nil {
//			break
//		}
//	}
//	return sts, err
//}
//
//func (me *Repository) FilesChanged() (fps Filepaths, err error) {
//	for range only.Once {
//		var out string
//		out, err = me.Exec("status", "--porcelain")
//		if err != nil {
//			break
//		}
//		sts := strings.Split(strings.TrimSpace(out), "\n")
//		fps = make(Filepaths, len(sts))
//		for i, fp := range sts {
//			fps[i] = fp[3:]
//		}
//	}
//	return fps, err
//}
//
//func (me *Repository) getHandle() (h *git.Repository, err error) {
//	for range only.Once {
//		h = me.Handle
//		if h == nil {
//			err = fmt.Errorf("repository handle is nil")
//			break
//		}
//	}
//	return h, err
//}
//
//func (me *Repository) GetTagObject(tag Tagname) (to *Tag, err error) {
//	for range only.Once {
//		var h *git.Repository
//		h, err = me.getHandle()
//		if err != nil {
//			break
//		}
//		var t *Reference
//		t, err = h.Tag(tag)
//		if err != nil {
//			break
//		}
//		to, err = h.TagObject(t.Hash())
//		if err != nil {
//			break
//		}
//	}
//	if err != nil {
//		err = fmt.Errorf("unable to access tag object '%s''; %s",
//			tag,
//			err.Error(),
//		)
//	}
//	return to, err
//}
//
//func (me *Repository) Lock() (ok bool, err error) {
//	for range only.Once {
//		var to *object.Tag
//		to, err = me.GetTagObject(LockTag)
//		if err != nil {
//			break
//		}
//		_ = to.ID()
//	}
//	return ok, err
//}
//
//func (me *Repository) Pull(opts ...*PullOptions) (err error) {
//	if len(opts) == 0 {
//		opts = []*PullOptions{}
//	}
//	for range only.Once {
//		var h *git.Repository
//		h, err = me.getHandle()
//		if err != nil {
//			break
//		}
//		wt, err := h.Worktree()
//		if err != nil {
//			break
//		}
//		err = wt.Pull(opts[0])
//	}
//	return err
//}
