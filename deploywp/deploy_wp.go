package deploywp

import (
	"fmt"
	"github.com/wplib/deploywp/cfg"
	"github.com/wplib/deploywp/util"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
)

type DeployWP struct {
	Config         *cfg.Config
	Branch         Reference
	SourceDir      Dir
	DestinationDir Dir
	Meta           *Meta
	Site           *Site
	Source         *Source
	Targets        *Targets
}

type Getter interface {
	GetConfig() *cfg.Config
	GetMeta() *Meta
	GetSite() *Site
	GetSource() *Source
	GetTargets() *Targets
}

func NewDeployWP(config *cfg.Config) *DeployWP {
	return &DeployWP{
		Config:    config,
		SourceDir: util.GetCurrentDir(),
	}
}

func (me *DeployWP) InitializeFromGetter(getter Getter) *DeployWP {
	me.Config = getter.GetConfig()
	me.Meta = getter.GetMeta()
	me.Site = getter.GetSite()
	me.Source = getter.GetSource()
	me.Targets = getter.GetTargets()
	return me
}

func (me *DeployWP) Run() {
	fmt.Printf("%#+v", me)
	return
}

func (me *DeployWP) Clone() {
	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/src-d/go-git",
		Progress: os.Stdout,
	})
	if err != nil {
		panic(err)
	}
}

func (me *DeployWP) Try() {
	r, err := git.PlainOpen(util.GetCurrentDir())
	if err != nil {
		panic(err)
	}

	wt, err := r.Worktree()
	if err != nil {
		panic(err)
	}
	sts, err := wt.Status()
	if err != nil {
		panic(err)
	}
	for fn := range sts {
		fmt.Printf("%#v\n", fn)
	}

	h, err := r.Head()
	if err != nil {
		panic(err)
	}

	bs, err := r.Branches()
	if err != nil {
		panic(err)
	}
	var branch string
	_ = bs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Hash() == h.Hash() {
			branch = util.AfterByte(ref.Name().String(), '/')
		}
		return nil
	})

	tos, err := r.TagObjects()
	if err != nil {
		panic(err)
	}
	var tag string
	_ = tos.ForEach(func(t *object.Tag) error {
		c, err := t.Commit()
		if err != nil {
			panic(err)
		}
		if c.Hash == h.Hash() {
			tag = t.Name
		}
		return nil

	})
	fmt.Printf("Current branch: %s\n", branch)
	fmt.Printf("Current tag: %s\n", tag)

}
