package deploywp

import (
	"fmt"
	"github.com/wplib/deploywp/app"
	"github.com/wplib/deploywp/cfg"
	"github.com/wplib/deploywp/git"
	"github.com/wplib/deploywp/util"
	"strings"
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

func (me *DeployWP) Try() {
	d := app.DeployDir
	r := git.NewRepository(d)

	err := r.Open()
	if err != nil {
		panic(err)
	}

	branch, err := r.Branch()
	if err != nil {
		panic(err)
	}

	tags, err := r.Tags()
	if err != nil {
		panic(err)
	}

	commit, err := r.Commit()
	if err != nil {
		panic(err)
	}
	var fc git.Filepaths
	hash := "no commits yet"
	if commit != nil {
		hash = commit.Hash
		fc, err = r.FilesChanged()
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("\nCommit: '%s'\n", hash)
	fmt.Printf("\nBranch: '%s'\n", branch)
	fmt.Printf("\nTag(s): '%s'\n", strings.Join(tags, ","))
	fmt.Printf("\nFiles Changed:\n")
	for _, n := range fc {
		fmt.Printf("- '%s'\n", n)
	}
	fmt.Printf("\n")

}
